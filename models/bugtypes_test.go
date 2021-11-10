package models

import (
	"reflect"
	"testing"
)

func TestBugTypes_GetBugTypeID(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bugtype.GetBugTypeID(); got != tt.want {
				t.Errorf("BugTypes.GetBugTypeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypes_GetBugTypeDescription(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bugtype.GetBugTypeDescription(); got != tt.want {
				t.Errorf("BugTypes.GetBugTypeDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypes_GetBugTypeAcronym(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bugtype.GetBugTypeAcronym(); got != tt.want {
				t.Errorf("BugTypes.GetBugTypeAcronym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypes_GetBugTypeName(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bugtype.GetBugTypeName(); got != tt.want {
				t.Errorf("BugTypes.GetBugTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypes_GetBugTypeCreatedAt(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bugtype.GetBugTypeCreatedAt(); got != tt.want {
				t.Errorf("BugTypes.GetBugTypeCreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypes_GetBugTypeUpdatedAt(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bugtype.GetBugTypeUpdatedAt(); got != tt.want {
				t.Errorf("BugTypes.GetBugTypeUpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypes_setBugTypeDescription(t *testing.T) {
	type args struct {
		bt string
	}
	tests := []struct {
		name    string
		bugtype *BugTypes
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bugtype.setBugTypeDescription(tt.args.bt)
		})
	}
}

func TestBugTypes_setBugTypeAcronym(t *testing.T) {
	type args struct {
		bt string
	}
	tests := []struct {
		name    string
		bugtype *BugTypes
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bugtype.setBugTypeAcronym(tt.args.bt)
		})
	}
}

func TestBugTypes_setBugTypeName(t *testing.T) {
	type args struct {
		bt string
	}
	tests := []struct {
		name    string
		bugtype *BugTypes
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bugtype.setBugTypeName(tt.args.bt)
		})
	}
}

func TestBugTypes_setBugTypeCreatedAt(t *testing.T) {
	type args struct {
		bt string
	}
	tests := []struct {
		name    string
		bugtype *BugTypes
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bugtype.setBugTypeCreatedAt(tt.args.bt)
		})
	}
}

func TestBugTypes_setBugTypeUpdatedAt(t *testing.T) {
	type args struct {
		bt string
	}
	tests := []struct {
		name    string
		bugtype *BugTypes
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bugtype.setBugTypeUpdatedAt(tt.args.bt)
		})
	}
}

func TestBugTypes_CreateNewBugType(t *testing.T) {
	tests := []struct {
		name    string
		bugtype *BugTypes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bugtype.CreateNewBugType(); (err != nil) != tt.wantErr {
				t.Errorf("BugTypes.CreateNewBugType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBugTypeExists(t *testing.T) {
	type args struct {
		acronym string
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
			if got := BugTypeExists(tt.args.acronym); got != tt.want {
				t.Errorf("BugTypeExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBugTypeExists(t *testing.T) {
	type args struct {
		acronym string
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
			if err := NewBugTypeExists(tt.args.acronym); (err != nil) != tt.wantErr {
				t.Errorf("NewBugTypeExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBugTypeGetAllInformations(t *testing.T) {
	type args struct {
		acronym string
	}
	tests := []struct {
		name    string
		args    args
		want    BugTypes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BugTypeGetAllInformations(tt.args.acronym)
			if (err != nil) != tt.wantErr {
				t.Errorf("BugTypeGetAllInformations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BugTypeGetAllInformations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBugTypeListOfAcronyms(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BugTypeListOfAcronyms()
			if (err != nil) != tt.wantErr {
				t.Errorf("BugTypeListOfAcronyms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BugTypeListOfAcronyms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
