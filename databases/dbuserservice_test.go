package databases

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCheckUserExists(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUserExists(tt.args.username); got != tt.want {
				t.Errorf("CheckUserExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewUser(t *testing.T) {
	type args struct {
		user bson.D
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewUser(tt.args.user)
		})
	}
}

func TestAuthentificationUser(t *testing.T) {
	type args struct {
		bcryptCost int
		username   string
		password   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AuthentificationUser(tt.args.bcryptCost, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("AuthentificationUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllUserInformations(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    bson.M
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllUserInformations(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUserInformations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUserInformations() = %v, want %v", got, tt.want)
			}
		})
	}
}
