package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	//err := godotenv.Load(".env")
	//In Railway, .env is not used, so we skip loading it
	//If running locally, .env is used, so we load it
	//This way, we can use the same code for both Railway and local development
	//If .env is not found, it will not throw an error, just a warning
	//This is useful for local development, so we can use .env to set environment variables
	//If running in Railway, .env is not used, so we skip loading it
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️ Skipping .env load: running in Railway or no local .env")
	}

}
