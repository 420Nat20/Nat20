package route

import (
	"encoding/json"
	"fmt"
	"nat-20/data"
	"nat-20/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ItemRoutes attaches routes for the item package to a router.
func (c *BaseController) RegisterGames() {
	c.Router.HandleFunc("/", c.getAllGames).Methods("GET")
	c.Router.HandleFunc("/{id}", c.getGame).Methods("GET")
	c.Router.HandleFunc("/", c.postGame).Methods("POST")
	c.Router.HandleFunc("/{id}", c.updateGame).Methods("UPDATE")
	c.Router.HandleFunc("/{id}", c.deleteGame).Methods("DELETE")

	c.Router.HandleFunc("/locations/", c.getAllGames).Methods("GET")
	c.Router.HandleFunc("/locations/{id}", c.getGame).Methods("GET")
	c.Router.HandleFunc("/locations/", c.postGame).Methods("POST")
	c.Router.HandleFunc("/locations/{id}", c.updateGame).Methods("UPDATE")
	c.Router.HandleFunc("/locations/{id}", c.deleteGame).Methods("DELETE")
}

func (c *BaseController) getAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	items, err := service.GetAllGameModels(c.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c *BaseController) getGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	item, err := service.GetGameModelByServerID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *BaseController) postGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newItem data.GameModel
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateGameModel(c.DB, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *BaseController) updateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newItem data.GameModel
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateGameModel(c.DB, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *BaseController) deleteGame(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	err = service.DeleteGameModel(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User deleted")
}
