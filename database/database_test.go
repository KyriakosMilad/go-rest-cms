package database

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestSetupDatabase(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(".env variables is needed to run database tests, error : " + err.Error())
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "test load .env variables", wantErr: false}, // required load env variables to run these tests
		{name: "test connection to database", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetupDatabase(); (err != nil) != tt.wantErr {
				t.Errorf("SetupDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
