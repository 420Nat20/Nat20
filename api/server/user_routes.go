package server

import (
	"encoding/json"
	"fmt"
	"github.com/420Nat20/Nat20/nat-20/common"
	"github.com/420Nat20/Nat20/nat-20/data/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// RegisterUserRoutes attaches routes to the given router.
func (s *Server) RegisterUserRoutes(baseUrl string) {
	s.Router.HandleFunc(baseUrl+"/", s.getAllUsers).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.getUser).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/", s.postUser).Methods("POST")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.updateUser).Methods("PUT")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.deleteUser).Methods("DELETE")
}

func (s *Server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	users, err := s.UserService.GetAllUsers()
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		httpError := common.GetHttpError(common.BadRequest.New("Invalid id"))
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	user, clientErr := s.UserService.GetUser(userId)
	if clientErr != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) postUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	clientErr := s.UserService.CreateUser(&user)
	if clientErr != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		httpError := common.GetHttpError(common.BadRequest.New("Invalid id"))
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	var body map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	user, clientErr := s.UserService.UpdateUser(userId, body)
	if clientErr != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		httpError := common.GetHttpError(common.BadRequest.New("Invalid id"))
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	clientErr := s.UserService.DeleteUser(userId)
	if clientErr != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	_, err = fmt.Fprint(w, "Delete Successful")
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}
