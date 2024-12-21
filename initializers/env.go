package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvInitializer() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("Env variables loaded successfully")
}
