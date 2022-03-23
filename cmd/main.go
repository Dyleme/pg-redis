package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Dyleme/pg-redis/internal/handler"
	"github.com/Dyleme/pg-redis/internal/server"
	"github.com/Dyleme/pg-redis/internal/service"
	"github.com/Dyleme/pg-redis/internal/storage"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	personStorage := storage.New(db)
	config := server.Config{
		Address: "",
		Port:    "8080",
	}
	personService := service.New(personStorage)
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
