package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Dyleme/pg-redis/internal/model"
)

type PersonService interface {
	// GetPerson(ctx context.Context, personID int) (*model.Person, error)
	// GetAllPersonsFromCountry(ctx context.Context, country string) ([]*model.Person, error)
	// DeletePerson(ctx context.Context, presonID int) error
	AddPerson(ctx context.Context, person *model.Person) (personID int, err error)
}

type PersonHandler struct {
	PersonService PersonService
}

func (ph *PersonHandler) AddPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var person model.Person

	err := decoder.Decode(&person)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := model.ValidatePerson(&person); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := ph.PersonService.AddPerson(ctx, &person)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResponse(w, id)
}
