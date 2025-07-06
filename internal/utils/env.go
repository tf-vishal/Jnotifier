package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Skipping .env load: running in Railway or no local .env")
	}

}
