package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Person interface {
	AddPerson(w http.ResponseWriter, r *http.Request)
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

	return &Handler{
		Router: router,
	}
}
