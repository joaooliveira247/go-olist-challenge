package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Database string connection
	DBURL = ""
	// API port, default is 3000
	APIPort = 0
)


func LoadEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	DBURL = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		log.Fatal("Error loading 'API_PORT' in .env file")
	}
}