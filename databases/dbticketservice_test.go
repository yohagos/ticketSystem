package databases

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewTicket(t *testing.T) {
	type args struct {
		ticket bson.D
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewTicket(tt.args.ticket)
		})
	}
}

func TestCheckTicketExists(t *testing.T) {
	type args struct {
		name string
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
			if got := CheckTicketExists(tt.args.name); got != tt.want {
				t.Errorf("CheckTicketExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllTicketInformations(t *testing.T) {
	type args struct {
		name string
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
			got, err := GetAllTicketInformations(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTicketInformations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTicketInformations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllTicketsByUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]bson.M
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AllTicketsByUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllTicketsByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllTicketsByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
