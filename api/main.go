package main

import (
	"log"
	"nat-20/data"
	"nat-20/route"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

type controller interface {
	Register()
}

func main() {
	router := mux.NewRouter()
	db := data.NewDB()

	registerControllers(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}

	withMiddleware := cors.
		Default().
		Handler(router)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, withMiddleware))
}

func registerControllers(r *mux.Router, db *gorm.DB) {
	var routers []controller

	// Add subrouters
	routers = append(routers,
		&route.GameController{
			DB:     db,
			Router: r.NewRoute().PathPrefix("/games").Subrouter(),
		},
	)

	routers = append(routers,
		&route.UserController{
			DB:     db,
			Router: r.NewRoute().PathPrefix("/games/{gameId}/users").Subrouter(),
		},
	)

	routers = append(routers,
		&route.LocationController{
			DB:     db,
			Router: r.NewRoute().PathPrefix("/games/{gameId}/locations").Subrouter(),
		},
	)

	routers = append(routers,
		&route.SubLocationController{
			DB:     db,
			Router: r.NewRoute().PathPrefix("/games/{gameId}/locations/{locationId}/sublocations").Subrouter(),
		},
	)

	for _, router := range routers {
		router.Register()
	}
}
