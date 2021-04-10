package main

import (
	"fmt"
	"log"
	"nat-20/data"
	"nat-20/route"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	server := &route.BaseController{
		DB:     data.NewDB(),
		Router: mux.NewRouter().StrictSlash(true),
	}

	attachControllers(server)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server.Router))
}

func attachControllers(c *route.BaseController) {
	// Add subrouters
	userController := &route.BaseController{
		DB:     c.DB,
		Router: c.Router.NewRoute().PathPrefix("/users").Subrouter(),
	}
	userController.RegisterUsers()
}
