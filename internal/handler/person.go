package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dyleme/pg-redis/internal/model"
	"github.com/gorilla/mux"
)

type PersonService interface {
	// GetPerson(ctx context.Context, personID int) (*model.Person, error)
	// GetAllPersonsFromCountry(ctx context.Context, country string) ([]*model.Person, error)
	// DeletePerson(ctx context.Context, presonID int) error
	AddPerson(ctx context.Context, person *model.Person) (personID int, err error)
	PersonList(ctx context.Context) ([]*model.Person, error)
	GetPerson(ctx context.Context, personID int) (*model.Person, error)
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

func (ph *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	strID, ok := vars["id"]
	if !ok {
		errorResponse(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(strID)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	person, err := ph.PersonService.GetPerson(ctx, id)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResponse(w, person)
}

func (ph *PersonHandler) PersonList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	list, err := ph.PersonService.PersonList(ctx)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResponse(w, list)
}
