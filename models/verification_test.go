package models

import "testing"

func TestUserVerification_GetVerifiedUserName(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.verif.GetVerifiedUserName(); got != tt.want {
				t.Errorf("UserVerification.GetVerifiedUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVerification_GetVerifiedUserLastname(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.verif.GetVerifiedUserLastname(); got != tt.want {
				t.Errorf("UserVerification.GetVerifiedUserLastname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVerification_GetVerifiedUserEmail(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.verif.GetVerifiedUserEmail(); got != tt.want {
				t.Errorf("UserVerification.GetVerifiedUserEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVerification_GetVerifiedUserPassword(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.verif.GetVerifiedUserPassword(); got != tt.want {
				t.Errorf("UserVerification.GetVerifiedUserPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVerification_GetVerifiedUserGeneratedKey(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.verif.GetVerifiedUserGeneratedKey(); got != tt.want {
				t.Errorf("UserVerification.GetVerifiedUserGeneratedKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVerification_GetVerifiedUserVerified(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.verif.GetVerifiedUserVerified(); got != tt.want {
				t.Errorf("UserVerification.GetVerifiedUserVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVerification_SetUserVerificationEmail(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		verif *UserVerification
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.SetUserVerificationEmail(tt.args.str)
		})
	}
}

func TestUserVerification_SetUserVerificationName(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		verif *UserVerification
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.SetUserVerificationName(tt.args.str)
		})
	}
}

func TestUserVerification_SetUserVerificationLastname(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		verif *UserVerification
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.SetUserVerificationLastname(tt.args.str)
		})
	}
}

func TestUserVerification_SetUserVerificationGeneratedKey(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		verif *UserVerification
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.SetUserVerificationGeneratedKey(tt.args.str)
		})
	}
}

func TestUserVerification_SetUserVerificationPassword(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		verif *UserVerification
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.SetUserVerificationPassword(tt.args.str)
		})
	}
}

func TestUserVerification_SetUserVerificationVerified(t *testing.T) {
	type args struct {
		bo bool
	}
	tests := []struct {
		name  string
		verif *UserVerification
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.SetUserVerificationVerified(tt.args.bo)
		})
	}
}

func TestUserVerification_CreateVerificationProfile(t *testing.T) {
	tests := []struct {
		name  string
		verif *UserVerification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.verif.CreateVerificationProfile()
		})
	}
}

func TestCheckVerification(t *testing.T) {
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
			got, got1 := CheckVerification(tt.args.key)
			if got != tt.want {
				t.Errorf("CheckVerification() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckVerification() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreateNewUser(t *testing.T) {
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
			CreateNewUser(tt.args.mail)
		})
	}
}
