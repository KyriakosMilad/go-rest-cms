package database

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

type test struct {
	id   int64  `column:"id INT(6) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	name string `column:"name VARCHAR(10) NOT NULL"`
}

func (t test) GetTable() string {
	return "tests"
}

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

func TestCreateTable(t *testing.T) {
	type args struct {
		v interface{ GetTable() string }
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "create table", args: args{test{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateTable(tt.args.v)
			if err != nil {
				t.Error(err.Error(), tt.args)
			}
		})
	}
}
