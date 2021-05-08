package user

import (
	"database/sql"
	"fmt"
	"github.com/KyriakosMilad/go-rest-cms/database"
	"reflect"
	"strconv"
	"time"
)

type User struct {
	ID        int64     `json:"id" db_column_name:"id" db_column_specs:"INT(6) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	FirstName string    `json:"first_name" db_column_name:"first_name" db_column_specs:"VARCHAR(10) NOT NULL"`
	LastName  string    `json:"last_name" db_column_name:"last_name" db_column_specs:"VARCHAR(10) NOT NULL"`
	Email     string    `json:"email"  db_column_name:"email" db_column_specs:"VARCHAR(255) NOT NULL UNIQUE"`
	Password  string    `json:"password"  db_column_name:"password" db_column_specs:"VARCHAR(255) NOT NULL"`
	CreatedAt time.Time `json:"created_at" db_column_name:"created_at" db_column_specs:"TIMESTAMP NOT NULL"`
	UpdatedAt time.Time `json:"updated_at" db_column_name:"updated_at" db_column_specs:"TIMESTAMP NOT NULL"`
	DeletedAt time.Time `json:"deleted_at" db_column_name:"deleted_at" db_column_specs:"TIMESTAMP"`
}

func (u User) GetTable() string {
	return "users"
}

func (u *User) Create() (err error) {
	timeNow := time.Now()

	var result sql.Result
	result, err = database.DB.Exec("INSERT INTO " + u.GetTable() + " (first_name,last_name,email,password,updated_at,created_at) VALUES (\"" + u.FirstName + "\",\"" + u.LastName + "\",\"" + u.Email + "\",\"" + u.Password + "\",\"" + timeNow.Format("2006-01-02 15:04:05") + "\",\"" + timeNow.Format("2006-01-02 15:04:05") + "\")")
	if err != nil {
		return
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return
	}

	u.CreatedAt = timeNow
	u.UpdatedAt = timeNow

	return
}

func (u *User) FindOne(conditions map[string]string) (err error) {
	query := "SELECT id,first_name,last_name,email,password,created_at,updated_at FROM " + u.GetTable() + " WHERE deleted_at IS NULL"

	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)
	for key, val := range conditions {
		// get field
		f := reflect.Indirect(v).FieldByName(key)
		if !f.IsValid() {
			err = fmt.Errorf("user schema does not have the field `%s`", key)
			return
		}

		// get tag db_column_name from field
		field, ok := t.Elem().FieldByName(key)
		if !ok {
			err = fmt.Errorf("user schema does not have the field `%s`", key)
			return
		}
		column := field.Tag.Get("db_column_name")

		// parse field value
		var fVal string
		if reflect.TypeOf(f) == reflect.TypeOf(time.Now()) {
			fVal = f.Interface().(time.Time).Format("2006-01-02 15:04:05")
		} else {
			fVal = fmt.Sprintf("%v", f)
		}

		query += " AND " + column + " " + val + " \"" + fVal + "\""
	}

	err = database.DB.QueryRow(query).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	return
}

func (u *User) FindAll(conditions map[string]string, limit int, offset int) (err error, users []User) {
	query := "SELECT id,first_name,last_name,email,password,created_at,updated_at FROM " + u.GetTable() + " WHERE deleted_at IS NULL"

	// parse conditions and add them to query
	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)
	for key, val := range conditions {
		// get field
		f := reflect.Indirect(v).FieldByName(key)
		if !f.IsValid() {
			err = fmt.Errorf("user schema does not have the field `%s`", key)
			return
		}

		// get tag db_column_name from field
		field, ok := t.Elem().FieldByName(key)
		if !ok {
			err = fmt.Errorf("user schema does not have the field `%s`", key)
			return
		}
		column := field.Tag.Get("db_column_name")

		// parse field value
		var fVal string
		if reflect.TypeOf(f) == reflect.TypeOf(time.Now()) {
			fVal = f.Interface().(time.Time).Format("2006-01-02 15:04:05")
		} else {
			fVal = fmt.Sprintf("%v", f)
		}

		query += " AND " + column + " " + val + " \"" + fVal + "\""
	}

	// pagination
	query += " LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa(offset)

	// get results
	var results *sql.Rows
	results, err = database.DB.Query(query)
	if err != nil {
		return
	}
	defer results.Close()

	// append results to users
	for results.Next() {
		err = results.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return
		}
		users = append(users, *u)
	}

	return
}
