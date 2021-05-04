package main

import (
	"github.com/KyriakosMilad/go-rest-cms/database"
	"github.com/KyriakosMilad/go-rest-cms/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load .env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// connect to to the database
	err = database.SetupDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database: " + err.Error())
	}

	server.Serve()
}
