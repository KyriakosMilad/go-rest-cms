package user

import (
	"database/sql"
	"github.com/KyriakosMilad/go-rest-cms/database"
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
	DeletedAt time.Time `json:"deleted_at" db_column_name:"deleted_at" db_column_name:"TIMESTAMP"`
}

func (u User) GetTable() string {
	return "users"
}

func (u *User) Create() (err error) {
	var result sql.Result
	result, err = database.DB.Exec("INSERT INTO " + u.GetTable() + " (first_name,last_name,email,password,created_at,updated_at) VALUES (\"" + u.FirstName + "\",\"" + u.LastName + "\",\"" + u.Email + "\",\"" + u.Password + "\",\"" + u.CreatedAt.Format("2006-01-02 15:04:05") + "\",\"" + u.UpdatedAt.Format("2006-01-02 15:04:05") + "\")")
	if err != nil {
		return
	}
	u.ID, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
