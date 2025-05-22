package router

import (
	"net/http"
	customcors "server/config"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	SetupUserRoutes(api)
	SetupAuthApp(api)
	SetupStaticRoutes(r)

	c := customcors.SetupCors()

	return c.Handler(r)
}
