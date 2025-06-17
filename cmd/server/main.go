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

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Kết nối database thất bại: %v", err)
	}
	fmt.Println("***GORM with first,last,... error record not found! GORM with find doesn't have this error (Return array object len == 0)***")
	app.Start()

	port := dotenv.GetDotEnv("APP_PORT")

	httpHandler := router.SetupRouter()

	fmt.Printf("Starting server on %s\n", port)
	log.Printf("Server start!")
	log.Fatal(http.ListenAndServe(":"+port, httpHandler))
	// err = http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", httpHandler)
	if err != nil {
		log.Fatalf("HTTPS server failed to start: %v", err)
	}
}
