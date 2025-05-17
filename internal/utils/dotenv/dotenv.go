package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv(key string) string {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
