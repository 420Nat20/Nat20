package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/420Nat20/Nat20/nat-20/data"
	"github.com/420Nat20/Nat20/nat-20/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserController struct {
	DB     *gorm.DB
	Router *mux.Router
}

// ItemRoutes attaches routes for the item package to a router.
func (c *UserController) Register() {
	c.Router.HandleFunc("/", c.getAllUsers).Methods("GET")
	c.Router.HandleFunc("/{id}", c.getUser).Methods("GET")
	c.Router.HandleFunc("/", c.postUser).Methods("POST")
	c.Router.HandleFunc("/{id}", c.updateUser).Methods("PUT")
	c.Router.HandleFunc("/{id}", c.deleteUser).Methods("DELETE")
}

func (c *UserController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	items, err := service.GetAllUserModels(c.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c *UserController) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	item, err := service.GetUserModelByDiscordID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *UserController) postUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	gameId, err := strconv.Atoi(params["gameId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newItem data.UserModel
	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateUserModel(c.DB, gameId, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	gameId, err := strconv.Atoi(params["gameId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDiscordId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	existingUser, err := service.GetUserModelByDiscordID(c.DB, userDiscordId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&existingUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateUserModel(c.DB, gameId, &existingUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingUser)
}

func (c *UserController) deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	err = service.DeleteUserModel(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User deleted")
}
