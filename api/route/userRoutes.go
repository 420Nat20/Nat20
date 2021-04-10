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
func (c *BaseController) RegisterUsers() {
	c.Router.HandleFunc("/", c.getAll).Methods("GET")
	c.Router.HandleFunc("/{id}", c.get).Methods("GET")
	c.Router.HandleFunc("/", c.post).Methods("POST")
	c.Router.HandleFunc("/{id}", c.update).Methods("UPDATE")
	c.Router.HandleFunc("/{id}", c.delete).Methods("DELETE")
}

func (c *BaseController) getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	items, err := service.GetAllUserModels(c.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c *BaseController) get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	item, err := service.GetUserModelByID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *BaseController) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newItem data.UserModel
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateUserModel(c.DB, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *BaseController) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newItem data.UserModel
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateUserModel(c.DB, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *BaseController) delete(w http.ResponseWriter, r *http.Request) {
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
