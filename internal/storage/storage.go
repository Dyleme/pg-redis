package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Dyleme/pg-redis/internal/model"
	postgres "github.com/Dyleme/pg-redis/internal/storage/postgres/db"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) AddPerson(ctx context.Context, person *model.Person) (personID int, err error) {
	q := postgres.New(s.db)

	res, err := q.AddPerson(ctx, postgres.AddPersonParams{})
	if err != nil {
		return 0, fmt.Errorf("postgres: %w", err)
	}

	index, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(index), nil
}
