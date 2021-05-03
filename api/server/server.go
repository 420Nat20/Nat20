package server

import (
	"context"
	"database/sql"
	"github.com/420Nat20/Nat20/nat-20/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	Ctx    context.Context
	DB     *sql.DB
	Router *mux.Router

	CampaignService service.CampaignService
	LocationService service.LocationService
	UserService     service.UserService
}

func (s *Server) InitServer() {
	s.Router = mux.NewRouter().StrictSlash(true)
	s.registerControllers()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}

	withMiddleware := cors.
		Default().
		Handler(s.Router)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, withMiddleware))
}

func (s *Server) registerControllers() {
	s.RegisterCampaignRoutes("campaigns")
	s.RegisterUserRoutes("/users")
	s.RegisterLocationRoutes("locations")
	s.RegisterSubLocationRoutes("locations/{locationId:[0-9]+}/sublocations")
}
