package user

import "time"

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
