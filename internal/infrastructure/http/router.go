package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userHandler *UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Define healthcheck endpoint
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	// Define user endpoints
	r.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.FindUserByID).Methods("GET")
	r.HandleFunc("/users", userHandler.AddUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}
