package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Server type holds information about the main server instance.
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func main() {
	server := &Server{
		DB:     newDB(),
		Router: mux.NewRouter().StrictSlash(true),
	}

	server.attachControllers()

	log.Fatal(http.ListenAndServe(":8000", server.Router))
}

func (s *Server) attachControllers() {
	// Add subrouters
	// userController := &taskitem.ItemController{
	// 	DB:     s.DB,
	// 	Router: s.Router.NewRoute().PathPrefix("/tasks").Subrouter(),
	// }
	// userController.Register()
}
