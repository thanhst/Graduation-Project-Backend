package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/app"
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
	app.Start()

	port := dotenv.GetDotEnv("APP_PORT")

	// Tạo một router mới bằng mux
	httpHandler := router.SetupRouter()

	// Bắt đầu server trên port 8080
	fmt.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, httpHandler))
}
