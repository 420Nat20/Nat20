package route

import (
	"encoding/json"
	"fmt"
	"nat-20/data"
	"nat-20/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type GameController struct {
	DB     *gorm.DB
	Router *mux.Router
}

// ItemRoutes attaches routes for the item package to a router.
func (c *GameController) Register() {
	c.Router.HandleFunc("/", c.getAllGames).Methods("GET")
	c.Router.HandleFunc("/{id}", c.getGame).Methods("GET")
	c.Router.HandleFunc("/", c.postGame).Methods("POST")
	c.Router.HandleFunc("/{id}", c.updateGame).Methods("PUT")
	c.Router.HandleFunc("/{id}", c.deleteGame).Methods("DELETE")
}

func (c *GameController) getAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	items, err := service.GetAllGameModels(c.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c *GameController) getGame(w http.ResponseWriter, r *http.Request) {
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

func (c *GameController) postGame(w http.ResponseWriter, r *http.Request) {
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

func (c *GameController) updateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingGame, err := service.GetGameModelByServerID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&existingGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateGameModel(c.DB, &existingGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingGame)
}

func (c *GameController) deleteGame(w http.ResponseWriter, r *http.Request) {
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
