package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
	"time"
)

var DB *sql.DB

func SetupDatabase() (err error) {
	DB, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True")
	if err != nil {
		return
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	err = DB.Ping()
	return
}

func CreateTable(v interface{ GetTable() string }) (err error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.TypeOf(struct{}{}).Kind() {
		return errors.New("can only proceed `struct` kind")
	}

	query := "CREATE TABLE IF NOT EXISTS " + v.GetTable() + "("
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("db_column_name") == "" || t.Field(i).Tag.Get("db_column_specs") == "" {
			continue
		}
		query += t.Field(i).Tag.Get("db_column_name") + " " + t.Field(i).Tag.Get("db_column_specs") + ","
	}
	// remove last comma if exists and close query
	if query[len(query)-1] == ',' {
		query = query[:len(query)-1]
	}
	query += ")"

	_, err = DB.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println(v.GetTable() + " migrated successfully")
	return
}

func DropTable(v interface{ GetTable() string }) (err error) {
	query := "DROP TABLE IF EXISTS " + v.GetTable()

	_, err = DB.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println(v.GetTable() + " dropped successfully")
	return
}
