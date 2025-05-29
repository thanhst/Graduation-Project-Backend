package customcors

import "github.com/rs/cors"

func SetupCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Upgrade", "Connection"},
		AllowCredentials: true,
	})
	return c
}
