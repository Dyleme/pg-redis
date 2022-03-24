package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Person interface {
	AddPerson(w http.ResponseWriter, r *http.Request)
	PersonList(w http.ResponseWriter, r *http.Request)
	GetPerson(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Ph Person
}

type Handler struct {
	ph Person

	*mux.Router
}

func (hs *Handlers) Route() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", hs.Ph.AddPerson)
	router.HandleFunc("/{id}", hs.Ph.GetPerson)
	router.HandleFunc("/list", hs.Ph.PersonList)

	return &Handler{
		Router: router,
	}
}
