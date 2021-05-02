package server

import (
	"net/http"
)

// RegisterSubLocationRoutes attaches routes to the given router.
func (s *Server) RegisterSubLocationRoutes(baseUrl string) {
	s.Router.HandleFunc(baseUrl+"/", s.getAllSubLocations).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.getSubLocation).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/", s.postSubLocation).Methods("POST")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.updateSubLocation).Methods("PUT")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.deleteSubLocation).Methods("DELETE")
}

func (s *Server) getAllSubLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) getSubLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) postSubLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) updateSubLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func (s *Server) deleteSubLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}
