package router

import (
	"KreditBee-assement/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/search", middleware.Search).Methods("GET", "OPTIONS")
	
	return router
}
