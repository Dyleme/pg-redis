package service

import (
	"context"

	"github.com/Dyleme/pg-redis/internal/model"
)

type PersonStorage interface {
	AddPerson(ctx context.Context, person *model.Person) (personID int, err error)
	GetPerson(ctx context.Context, id int) (person *model.Person, err error)
	PersonList(ctx context.Context) ([]*model.Person, error)
}

type Service struct {
	ps PersonStorage
}

func New(ps PersonStorage) *Service {
	return &Service{
		ps: ps,
	}
}

func (s *Service) GetPerson(ctx context.Context, id int) (*model.Person, error) {
	return s.ps.GetPerson(ctx, id)
}

func (s *Service) AddPerson(ctx context.Context, person *model.Person) (personID int, err error) {
	return s.ps.AddPerson(ctx, person)
}

func (s *Service) PersonList(ctx context.Context) ([]*model.Person, error) {
	return s.ps.PersonList(ctx)
}
