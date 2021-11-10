package routes

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexGETHandler(t *testing.T) {
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
			indexGETHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_pageNotFoundHandler(t *testing.T) {
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
			pageNotFoundHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestCheckCurrentSession(t *testing.T) {
	type args struct {
		r *http.Request
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
			if got := CheckCurrentSession(tt.args.r); got != tt.want {
				t.Errorf("CheckCurrentSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveCurrentSession(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		key string
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
			if err := SaveCurrentSession(tt.args.w, tt.args.r, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("SaveCurrentSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
