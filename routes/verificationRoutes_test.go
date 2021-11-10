package routes

import (
	"net/http"
	"testing"
)

func TestVerificationGETHandler(t *testing.T) {
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
			VerificationGETHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestVerificationPOSTHandler(t *testing.T) {
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
			VerificationPOSTHandler(tt.args.w, tt.args.r)
		})
	}
}
