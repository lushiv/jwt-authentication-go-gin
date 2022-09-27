package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading .env files")
	}
}
