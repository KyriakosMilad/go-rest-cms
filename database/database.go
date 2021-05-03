package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
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

func CreateTable(v interface{ GetTable() string }) (err error) {
	query := "CREATE TABLE IF NOT EXISTS " + v.GetTable() + "("
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		if (t.NumField() - i) > 1 {
			query += t.Field(i).Tag.Get("column") + ","
		} else {
			query += t.Field(i).Tag.Get("column")
		}
	}
	query += ")"

	_, err = DB.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println(v.GetTable() + " migrated successfully")
	return
}
