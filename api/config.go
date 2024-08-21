package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func envAccountSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file")
	}

	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAuthToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return os.Getenv("TWILIO_AUTHTOKEN")
}

func envServicesID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return os.Getenv("TWILIO_SERVICES_ID")
}
