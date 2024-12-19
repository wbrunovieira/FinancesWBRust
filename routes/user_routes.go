package routes

import (
	"app/handlers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/users").Subrouter()
	subRouter.HandleFunc("", handlers.CreateUser).Methods("POST", "OPTIONS")
	subRouter.HandleFunc("", handlers.GetUsers).Methods("GET", "OPTIONS")
	subRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUser).Methods("GET", "OPTIONS")
	subRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT", "OPTIONS")
	subRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE", "OPTIONS")
}
