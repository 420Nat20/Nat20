package route

import (
	"encoding/json"
	"fmt"
	"github.com/420Nat20/Nat20/nat-20/data/model"
	"net/http"
	"strconv"

	"github.com/420Nat20/Nat20/nat-20/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type SubLocationController struct {
	DB     *gorm.DB
	Router *mux.Router
}

// ItemRoutes attaches routes for the item package to a router.
func (c *SubLocationController) Register() {
	c.Router.HandleFunc("/", c.getAllSubLocations).Methods("GET")
	c.Router.HandleFunc("/{id}", c.getSubLocation).Methods("GET")
	c.Router.HandleFunc("/", c.postSubLocation).Methods("POST")
	c.Router.HandleFunc("/{id}", c.updateSubLocation).Methods("PUT")
	c.Router.HandleFunc("/{id}", c.deleteSubLocation).Methods("DELETE")
}

func (c *SubLocationController) getAllSubLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	locationId, err := strconv.Atoi(params["locationId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items, err := service.GetAllSubLocationModels(c.DB, locationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c *SubLocationController) getSubLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	item, err := service.GetSubLocationModelByID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *SubLocationController) postSubLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	gameId, err := strconv.Atoi(params["locationId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newItem model.SubLocationModel
	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateSubLocationModel(c.DB, gameId, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *SubLocationController) updateSubLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	existingSubLocation, err := service.GetSubLocationModelByID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&existingSubLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateSubLocationModel(c.DB, id, &existingSubLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingSubLocation)
}

func (c *SubLocationController) deleteSubLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	err = service.DeleteSubLocationModel(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User deleted")
}
