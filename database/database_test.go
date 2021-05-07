package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

type test struct {
	id   int64  `db_column_name:"id" db_column_specs:"INT(6) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	name string `db_column_name:"name" db_column_specs:"VARCHAR(10) NOT NULL"`
}

func (t test) GetTable() string {
	return "tests"
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(".env variables is needed to run database tests, error : " + err.Error())
	}

	os.Exit(m.Run())
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

func TestDropTable(t *testing.T) {
	type args struct {
		v interface{ GetTable() string }
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "drop table", args: args{test{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DropTable(tt.args.v)
			if err != nil {
				t.Error(err.Error(), tt.args)
			}
		})
	}
}
