package routes

import (
	"net/http"
	"testing"
)

func TestBugtypeGETHandler(t *testing.T) {
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
			BugtypeGETHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestBugtypePOSTHandler(t *testing.T) {
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
			BugtypePOSTHandler(tt.args.w, tt.args.r)
		})
	}
}
