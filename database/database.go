package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var DB *sql.DB

func SetupDatabase() (err error) {
	DB, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		return
	}
	err = DB.Ping()
	return
}
