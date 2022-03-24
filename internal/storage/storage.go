package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dyleme/pg-redis/internal/model"
	queries "github.com/Dyleme/pg-redis/internal/storage/repository/db"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func fillAddress(address *model.Address) *queries.Address {
	var apartments sql.NullString
	if address.Apartments != "" {
		apartments = sql.NullString{
			String: address.Apartments,
			Valid:  true,
		}
	}

	return &queries.Address{
		Country:    address.Country,
		City:       address.City,
		Street:     address.Street,
		House:      address.House,
		Apartments: apartments,
	}
}

func getID(result sql.Result, err error) (int, error) {
	if err != nil {
		return 0, fmt.Errorf("mysql: %w", err)
	}

	index, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(index), nil
}

var ErrGetPerson = errors.New("no such person")

func (s *Storage) GetPerson(ctx context.Context, id int) (*model.Person, error) {
	var person model.Person
	err := s.execTx(ctx, func(ctx context.Context, q *queries.Queries) error {
		personSQL, err := q.PersonByID(ctx, int32(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrGetPerson
			}
			return err
		}
		address := model.Address{
			Country:    personSQL.Country,
			City:       personSQL.City,
			Street:     personSQL.Street,
			House:      personSQL.House,
			Apartments: personSQL.Apartments.String,
		}
		person = model.Person{
			Name:    personSQL.PersonName,
			Phone:   personSQL.Phone,
			Address: address,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (s *Storage) PersonList(ctx context.Context) ([]*model.Person, error) {
	list := make([]*model.Person, 0)
	err := s.execTx(ctx, func(ctx context.Context, q *queries.Queries) error {
		personsSQL, err := q.PersonList(ctx)
		if err != nil {
			return err
		}
		for _, personSQL := range personsSQL {
			address := model.Address{
				Country:    personSQL.Country,
				City:       personSQL.City,
				Street:     personSQL.Street,
				House:      personSQL.House,
				Apartments: personSQL.Apartments.String,
			}
			person := model.Person{
				Name:    personSQL.PersonName,
				Phone:   personSQL.Phone,
				Address: address,
			}
			list = append(list, &person)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Storage) AddPerson(ctx context.Context, person *model.Person) (personID int, err error) {
	err = s.execTx(ctx, func(ctx context.Context, q *queries.Queries) error {
		address := fillAddress(&person.Address)
		addressID32, err := q.AddressID(ctx, queries.AddressIDParams{
			Country:    address.Country,
			City:       address.City,
			Street:     address.Street,
			House:      address.House,
			Apartments: address.Apartments,
		})
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
		}
		addressID := int(addressID32)

		if addressID == 0 {
			addressID, err = getID(q.AddAddress(ctx, queries.AddAddressParams{
				Country:    address.Country,
				City:       address.City,
				Street:     address.Street,
				House:      address.House,
				Apartments: address.Apartments,
			}))
			if err != nil {
				return err
			}
		}

		personID, err = getID(q.AddPerson(ctx, queries.AddPersonParams{
			PersonName: person.Name,
			Phone:      person.Phone,
			Addressid:  int32(addressID),
		}))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return personID, nil
}

func (s *Storage) execTx(ctx context.Context, f func(ctx context.Context, q *queries.Queries) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	q := queries.New(tx)

	err = f(ctx, q)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return fmt.Errorf("rollback caused by %v failed: %w", err, txErr)
		}

		return fmt.Errorf("successful rollback caused by %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit error: %w", err)
	}

	return nil
}
