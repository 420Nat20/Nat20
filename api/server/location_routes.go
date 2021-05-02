package server

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type LocationController struct {
	Router *mux.Router
	DB     *sql.DB
}

// RegisterLocationRoutes attaches routes to the given router.
func (s *Server) RegisterLocationRoutes(baseUrl string) {
	s.Router.HandleFunc(baseUrl+"/", s.getAllLocations).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/{id}", s.getLocation).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/", s.postLocation).Methods("POST")
	s.Router.HandleFunc(baseUrl+"/{id}", s.updateLocation).Methods("PUT")
	s.Router.HandleFunc(baseUrl+"/{id}", s.deleteLocation).Methods("DELETE")
}

func (s *Server) getAllLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) getLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) postLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) updateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) deleteLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}
