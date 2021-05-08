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

func TestUser_FindOne(t *testing.T) {
	type args struct {
		conditions map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test find user by id",
			args:    args{map[string]string{"ID": "="}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.FindOne(tt.args.conditions); (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v, args %v", err, tt.wantErr, tt.args.conditions)
			}
		})
	}
}

func TestUser_FindAll(t *testing.T) {
	type args struct {
		conditions map[string]string
		limit      int
		offset     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test find all users",
			args: args{
				conditions: map[string]string{},
				limit:      10,
				offset:     0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, users := u.FindAll(tt.args.conditions, tt.args.limit, tt.args.offset)
			if err != nil {
				t.Errorf("FindAll() err = %v, want %v", err, tt.wantErr)
			}
			if len(users) == 0 {
				t.Errorf("FindAll() return zero results, 1 expected at least")
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test update user",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u.FirstName = "Kero"
			if err := u.Update(); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_SoftDelete(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test soft delete user",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.SoftDelete(); (err != nil) != tt.wantErr {
				t.Errorf("SoftDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
