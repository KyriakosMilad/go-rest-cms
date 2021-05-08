package user

import (
	"github.com/KyriakosMilad/go-rest-cms/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var u = User{
	FirstName: "Kyriakos",
	LastName:  "Milad",
	Email:     "contact.kyriakos@gmail.com",
	Password:  "0123456",
}

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
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test create user",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
