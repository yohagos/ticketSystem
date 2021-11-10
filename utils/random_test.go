package utils

import "testing"

func TestRandomKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomKey(); got != tt.want {
				t.Errorf("RandomKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomFiveDigitNumber(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomFiveDigitNumber(); got != tt.want {
				t.Errorf("RandomFiveDigitNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateVerificationKey(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateVerificationKey(); got != tt.want {
				t.Errorf("GenerateVerificationKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomStringBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomStringBytes(tt.args.n); got != tt.want {
				t.Errorf("randomStringBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
