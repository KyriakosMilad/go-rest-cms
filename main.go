package main

import (
	"github.com/KyriakosMilad/go-rest-cms/database"
	"github.com/KyriakosMilad/go-rest-cms/schema"
	"github.com/KyriakosMilad/go-rest-cms/server"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-migrate":
			schema.Migrate()
		case "-drop":
			schema.Drop()
		case "-fresh":
			schema.Drop()
			schema.Migrate()
		default:
			server.Serve()
		}
	} else {
		server.Serve()
	}
}
