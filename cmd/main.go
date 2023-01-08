package main

import (
	"advertising/internal/handler"
	"advertising/internal/repository"
	"advertising/internal/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	dbURI := "postgres://postgres:password@localhost:5432/advertising?sslmode=disable"
	db, err := repository.NewPostgresDB(dbURI)
	if err != nil {
		logrus.Fatal("failed to connect to db: %s\n", err.Error())
	}
	newPostgres := repository.NewProductPostgres(db)
	newService := service.NewServices(newPostgres)
	handlers := handler.NewHandler(newService)
	if err := http.ListenAndServe(":8080", handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
