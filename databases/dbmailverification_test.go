package databases

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCheckVerificationExists(t *testing.T) {
	type args struct {
		mail string
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
			if got := CheckVerificationExists(tt.args.mail); got != tt.want {
				t.Errorf("CheckVerificationExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewVerificationProfile(t *testing.T) {
	type args struct {
		profile bson.D
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewVerificationProfile(tt.args.profile)
		})
	}
}

func TestGetVerificationKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetVerificationKey(tt.args.key)
			if got != tt.want {
				t.Errorf("GetVerificationKey() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetVerificationKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeleteVerificationDocument(t *testing.T) {
	type args struct {
		mail string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteVerificationDocument(tt.args.mail)
		})
	}
}

func TestGetAllVerificationInformation(t *testing.T) {
	type args struct {
		mail string
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllVerificationInformation(tt.args.mail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllVerificationInformation() = %v, want %v", got, tt.want)
			}
		})
	}
}
