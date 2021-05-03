package main

import (
	"fmt"
	"github.com/KyriakosMilad/go-rest-cms/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

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

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
