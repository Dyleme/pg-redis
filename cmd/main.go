package main

import (
	"context"
	"log"

	"github.com/Dyleme/pg-redis/internal/handler"
	"github.com/Dyleme/pg-redis/internal/server"
	"github.com/Dyleme/pg-redis/internal/service"
)

func main() {
	config := server.Config{
		Address: "",
		Port:    "8080",
	}
	personService := &service.Service{}
	personHandler := handler.PersonHandler{
		PersonService: personService,
	}
	handlers := handler.Handlers{
		Ph: &personHandler,
	}

	srvr := server.New(handlers.Route(), config)
	if err := srvr.Run(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
