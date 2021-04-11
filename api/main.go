package main

import (
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
	gamesController := &route.BaseController{
		DB:     c.DB,
		Router: c.Router.NewRoute().PathPrefix("/games").Subrouter(),
	}
	gamesController.RegisterGames()

	userController := &route.BaseController{
		DB:     c.DB,
		Router: c.Router.NewRoute().PathPrefix("/games/{gameId}/users").Subrouter(),
	}
	userController.RegisterUsers()

	locationsController := &route.BaseController{
		DB:     c.DB,
		Router: c.Router.NewRoute().PathPrefix("/games/{gameId}/locations").Subrouter(),
	}
	locationsController.RegisterLocations()

	subLocationsController := &route.BaseController{
		DB:     c.DB,
		Router: c.Router.NewRoute().PathPrefix("/games/{gameId}/locations/{locationId}/sublocations").Subrouter(),
	}
	subLocationsController.RegisterSubLocations()
}
