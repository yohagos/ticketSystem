package utils

import "testing"

func TestIsError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IsError(tt.args.err)
		})
	}
}

func TestCreateTimeStamp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTimeStamp(); got != tt.want {
				t.Errorf("CreateTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
