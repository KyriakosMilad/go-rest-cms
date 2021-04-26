package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	println(os.Getenv("NAME"))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
