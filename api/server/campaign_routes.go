package server

import (
	"net/http"
)

// RegisterCampaignRoutes attaches routes to the given router.
func (s *Server) RegisterCampaignRoutes(baseUrl string) {
	s.Router.HandleFunc(baseUrl+"/", s.getAllGames).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/{id}", s.getGame).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/", s.postGame).Methods("POST")
	s.Router.HandleFunc(baseUrl+"/{id}", s.updateGame).Methods("PUT")
	s.Router.HandleFunc(baseUrl+"/{id}", s.deleteGame).Methods("DELETE")
}

func (s *Server) getAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(items)
}

func (s *Server) getGame(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(item)
}

func (s *Server) postGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(newItem)
}

func (s *Server) updateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//params := mux.Vars(r)
	//id, err := strconv.Atoi(params["id"])
	//
	//json.NewEncoder(w).Encode(existingGame)
}

func (s *Server) deleteGame(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")
	//fmt.Fprintf(w, "User deleted")
}
