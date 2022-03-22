package service

import (
	"context"

	"github.com/Dyleme/pg-redis/internal/model"
)

type PersonStorage interface {
	AddPerson(ctx context.Context, person *model.Person) (personID int, err error)
}

type Service struct {
	ps PersonStorage
}

func New(ps PersonStorage) *Service {
	return &Service{
		ps: ps,
	}
}

func (s *Service) AddPerson(ctx context.Context, person *model.Person) (personID int, err error) {
	return 1, nil
}
