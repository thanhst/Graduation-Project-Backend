package router

import (
	"net/http"
	"server/internal/utils/dotenv"

	"github.com/gorilla/mux"
)

func SetupStaticRoutes(r *mux.Router) {
	uploadDir := dotenv.GetDotEnv("UPLOAD_DIR")
	fs := http.FileServer(http.Dir(uploadDir))
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))
}
