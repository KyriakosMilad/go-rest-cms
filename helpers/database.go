package helpers

import (
	"database/sql"
	"log"
	"os"
)

func SetupDatabase() {
	_, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
}
