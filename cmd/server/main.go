package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/api/handler"
	"server/internal/db/database"
	"server/internal/utils/dotenv"

	"github.com/gorilla/mux"
)

// main function to set up the HTTP server
func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Kết nối database thất bại: %v", err)
	}

	// Tạo một router mới bằng mux
	router := mux.NewRouter()

	// Static routes
	uploadDir := dotenv.GetDotEnv("UPLOAD_DIR")
	port := dotenv.GetDotEnv("APP_PORT")
	fs := http.FileServer(http.Dir(uploadDir))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	// API routes
	router.HandleFunc("/api/register", handler.RegisterHandler).Methods("POST")

	// Cấu hình HTTP server
	http.Handle("/", router)

	// Bắt đầu server trên port 8080
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
