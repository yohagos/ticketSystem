package models

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestUser_GetUserID(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserID(); got != tt.want {
				t.Errorf("User.GetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUserName(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserName(); got != tt.want {
				t.Errorf("User.GetUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUserLastname(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserLastname(); got != tt.want {
				t.Errorf("User.GetUserLastname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUserEmail(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserEmail(); got != tt.want {
				t.Errorf("User.GetUserEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUserPassword(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserPassword(); got != tt.want {
				t.Errorf("User.GetUserPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUserCreatedAt(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserCreatedAt(); got != tt.want {
				t.Errorf("User.GetUserCreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUserUpdatedAt(t *testing.T) {
	tests := []struct {
		name string
		user *User
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserUpdatedAt(); got != tt.want {
				t.Errorf("User.GetUserUpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_setUserEmail(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		user *User
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.setUserEmail(tt.args.str)
		})
	}
}

func TestUser_setUserName(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		user *User
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.setUserName(tt.args.str)
		})
	}
}

func TestUser_setUserLastname(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		user *User
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.setUserLastname(tt.args.str)
		})
	}
}

func TestUser_setUserPassword(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		user *User
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.setUserPassword(tt.args.str)
		})
	}
}

func TestUser_setUserCreatedAt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		user *User
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.setUserCreatedAt(tt.args.str)
		})
	}
}

func TestUser_setUserUpdatedAt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		user *User
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.user.setUserUpdatedAt(tt.args.str)
		})
	}
}

func TestUser_CreateNewUser(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.CreateNewUser(); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserExists(t *testing.T) {
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
			if got := UserExists(tt.args.username); got != tt.want {
				t.Errorf("UserExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAuthentification(t *testing.T) {
	type args struct {
		username string
		password string
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
			if err := UserAuthentification(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UserAuthentification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserGetAllInformations(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserGetAllInformations(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserGetAllInformations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserGetAllInformations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBsonToUser(t *testing.T) {
	type args struct {
		list bson.M
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BsonToUser(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BsonToUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
