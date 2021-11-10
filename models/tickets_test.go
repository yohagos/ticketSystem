package models

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestTickets_GetTicketID(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketID(); got != tt.want {
				t.Errorf("Tickets.GetTicketID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketName(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketName(); got != tt.want {
				t.Errorf("Tickets.GetTicketName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketCreatedBy(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketCreatedBy(); got != tt.want {
				t.Errorf("Tickets.GetTicketCreatedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketBugType(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketBugType(); got != tt.want {
				t.Errorf("Tickets.GetTicketBugType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketStatus(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketStatus(); got != tt.want {
				t.Errorf("Tickets.GetTicketStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketAllComments(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   []Comment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketAllComments(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tickets.GetTicketAllComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketUpdatedAt(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketUpdatedAt(); got != tt.want {
				t.Errorf("Tickets.GetTicketUpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_GetTicketCreatedAt(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ticket.GetTicketCreatedAt(); got != tt.want {
				t.Errorf("Tickets.GetTicketCreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTickets_SetTicketName(t *testing.T) {
	type args struct {
		tic string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketName(tt.args.tic)
		})
	}
}

func TestTickets_SetTicketCreatedBy(t *testing.T) {
	type args struct {
		tic string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketCreatedBy(tt.args.tic)
		})
	}
}

func TestTickets_SetTicketBugType(t *testing.T) {
	type args struct {
		tic string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketBugType(tt.args.tic)
		})
	}
}

func TestTickets_SetTicketStatus(t *testing.T) {
	type args struct {
		tic string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketStatus(tt.args.tic)
		})
	}
}

func TestTickets_SetTicketComment(t *testing.T) {
	type args struct {
		user      string
		message   string
		createdAt string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketComment(tt.args.user, tt.args.message, tt.args.createdAt)
		})
	}
}

func TestTickets_SetTicketUpdatedAt(t *testing.T) {
	type args struct {
		tic string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketUpdatedAt(tt.args.tic)
		})
	}
}

func TestTickets_SetTicketCreatedAt(t *testing.T) {
	type args struct {
		tic string
	}
	tests := []struct {
		name   string
		ticket *Tickets
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.SetTicketCreatedAt(tt.args.tic)
		})
	}
}

func TestTickets_CreateNewTicket(t *testing.T) {
	tests := []struct {
		name   string
		ticket *Tickets
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ticket.CreateNewTicket()
		})
	}
}

func TestNewTicketExists(t *testing.T) {
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
			if got := NewTicketExists(tt.args.name); got != tt.want {
				t.Errorf("NewTicketExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTicketGetAllInformations(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Tickets
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TicketGetAllInformations(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketGetAllInformations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TicketGetAllInformations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTicketsByUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]Tickets
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTicketsByUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicketsByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTicketsByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bsonToTicketConverter(t *testing.T) {
	type args struct {
		list *[]bson.M
	}
	tests := []struct {
		name string
		args args
		want *[]Tickets
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bsonToTicketConverter(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bsonToTicketConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
