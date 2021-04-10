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
	server.Router.HandleFunc("/", HelloServer).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, server.Router))
}

func attachControllers(c *route.BaseController) {
	// Add subrouters
	userController := &route.BaseController{
		DB:     c.DB,
		Router: c.Router.NewRoute().PathPrefix("/users").Subrouter(),
	}
	userController.RegisterUsers()
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there.")
}
