package routes

import (
	"net/http"
	"testing"
)

func TestTicketsGETHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TicketsGETHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestTicketsPOSTHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TicketsPOSTHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestTicketDetailGETHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TicketDetailGETHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestTicketDetailPOSTHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TicketDetailPOSTHandler(tt.args.w, tt.args.r)
		})
	}
}
