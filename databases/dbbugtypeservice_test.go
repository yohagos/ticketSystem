package databases

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewBugType(t *testing.T) {
	type args struct {
		bugtype bson.D
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
			if err := CreateNewBugType(tt.args.bugtype); (err != nil) != tt.wantErr {
				t.Errorf("CreateNewBugType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckBugTypeExists(t *testing.T) {
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
			if got := CheckBugTypeExists(tt.args.acronym); got != tt.want {
				t.Errorf("CheckBugTypeExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllBugTypeInformations(t *testing.T) {
	type args struct {
		acronym string
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
			got, err := GetAllBugTypeInformations(tt.args.acronym)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllBugTypeInformations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllBugTypeInformations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetListOfAllBugTypes(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetListOfAllBugTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListOfAllBugTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListOfAllBugTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterBugTypeList(t *testing.T) {
	type args struct {
		bugtypeList []bson.M
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
			if got := filterBugTypeList(tt.args.bugtypeList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterBugTypeList() = %v, want %v", got, tt.want)
			}
		})
	}
}
