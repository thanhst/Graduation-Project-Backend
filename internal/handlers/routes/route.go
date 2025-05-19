package router

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	SetupUserRoutes(api)
	SetupStaticRoutes(r)

	return r
}
