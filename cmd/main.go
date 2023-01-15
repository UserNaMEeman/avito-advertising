package main

import (
	"advertising/internal/handler"
	"advertising/internal/repository"
	"advertising/internal/service"
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	usernameDB *string
	passwordDB *string
	addrDB     *string
	port       *string
)

func init() {
	usernameDB = flag.String("u", "postgres", "username db")
	passwordDB = flag.String("pd", "password", "password db")
	addrDB = flag.String("ad", "localhost:5432", "address db")
	port = flag.String("p", "8080", "listen port")
}

func main() {
	flag.Parse()
	portList := ":" + *port
	dbURI := "postgres://" + *usernameDB + ":" + *passwordDB + "@" + *addrDB + "/postgres?sslmode=disable"
	fmt.Println(dbURI)
	// fmt.Println(*port)
	db, err := repository.NewPostgresDB(dbURI)
	if err != nil {
		logrus.Fatal("failed to connect to db: %s\n", err.Error())
	}
	if err = repository.InitDB(db); err != nil {
		panic(err)
	}
	newPostgres := repository.NewProductPostgres(db)
	newService := service.NewServices(newPostgres)
	handlers := handler.NewHandler(newService)
	if err := http.ListenAndServe(portList, handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
