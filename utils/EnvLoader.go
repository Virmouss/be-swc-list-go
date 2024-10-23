package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Println("cant load env files")
	}
}
