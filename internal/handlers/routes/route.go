package router

import (
	"net/http"
	customcors "server/config"
	websocket "server/internal/handlers/websocket"

	"github.com/gorilla/mux"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	SetupUserRoutes(api)
	SetupAuthApp(api)
	SetupClassroomApp(api)
	SetupStaticRoutes(r)
	websocket.SetupWebsocket(r)

	c := customcors.SetupCors()

	return c.Handler(r)
}
