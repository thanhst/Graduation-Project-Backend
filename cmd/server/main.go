package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/db/database"
	router "server/internal/handlers/routes"
	"server/internal/utils/dotenv"
)

// main function to set up the HTTP server
func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Kết nối database thất bại: %v", err)
	}

	port := dotenv.GetDotEnv("APP_PORT")

	// Tạo một router mới bằng mux
	r := router.SetupRouter()

	// Cấu hình HTTP server
	http.Handle("/", r)

	// Bắt đầu server trên port 8080
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
