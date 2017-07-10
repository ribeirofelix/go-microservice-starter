package router

import (
	"github.com/gorilla/mux"
)

// NewRouter -
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
