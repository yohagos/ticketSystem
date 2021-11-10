package mails

import "testing"

func TestEmailInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmailInit()
		})
	}
}

func TestSendVerificationMail(t *testing.T) {
	type args struct {
		username string
		email    string
		key      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendVerificationMail(tt.args.username, tt.args.email, tt.args.key)
		})
	}
}
