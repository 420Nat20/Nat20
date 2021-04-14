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

type LocationController struct {
	DB     *gorm.DB
	Router *mux.Router
}

// ItemRoutes attaches routes for the item package to a router.
func (c *LocationController) Register() {
	c.Router.HandleFunc("/", c.getAllLocations).Methods("GET")
	c.Router.HandleFunc("/{id}", c.getLocation).Methods("GET")
	c.Router.HandleFunc("/", c.postLocation).Methods("POST")
	c.Router.HandleFunc("/{id}", c.updateLocation).Methods("PUT")
	c.Router.HandleFunc("/{id}", c.deleteLocation).Methods("DELETE")
}

func (c *LocationController) getAllLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	items, err := service.GetAllLocationModels(c.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (c *LocationController) getLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	item, err := service.GetLocationModelByID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (c *LocationController) postLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	gameId, err := strconv.Atoi(params["gameId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newItem data.LocationModel
	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateLocationModel(c.DB, gameId, &newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newItem)
}

func (c *LocationController) updateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	existingLocation, err := service.GetLocationModelByID(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&existingLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateLocationModel(c.DB, id, &existingLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingLocation)
}

func (c *LocationController) deleteLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	w.Header().Add("Content-Type", "application/json")

	err = service.DeleteLocationModel(c.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User deleted")
}
