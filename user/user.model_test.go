package user

import (
	"github.com/KyriakosMilad/go-rest-cms/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// load .env variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(".env variables is needed to run database tests, error : " + err.Error())
	}

	// connect to to the database
	err = database.SetupDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database: " + err.Error())
	}

	os.Exit(m.Run())
}

func TestUser_Create(t *testing.T) {
	type fields struct {
		ID        int64
		FirstName string
		LastName  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test create user",
			fields: fields{
				FirstName: "Kyriakos",
				LastName:  "Milad",
				Email:     "contact.kyriakos@gmail.com",
				Password:  "0635156",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := u.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
