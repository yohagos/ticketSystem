package apperrors

import (
	"errors"
	"net/http"
)

var (
	// ErrorSessionInvalid error
	ErrorSessionInvalid = errors.New("current Session is invalid. Please login first")

	// ErrorPasswordMismatch error
	ErrorPasswordMismatch = errors.New("password mismatches")

	// ErrorRoutesPageNotFound error
	ErrorRoutesPageNotFound = errors.New("page does not exist. Please try again")
	// ErrorRoutesPageDoesntExist error
	ErrorRoutesPageDoesntExist = errors.New("this page does not exist")
	// ErrorRoutesUserNotFound error
	ErrorRoutesUserNotFound = errors.New("username not found")
	// ErrorRoutesInvalidLogin error
	ErrorRoutesInvalidLogin = errors.New("invalid login. Please try again")

	// ErrorUserDoesNotExist error
	ErrorUserDoesNotExist = errors.New("login is invalid. Username / Password does not exists")
	// ErrorUserAlreadyExists error
	ErrorUserAlreadyExists = errors.New("user already exists")

	// ErrorBugTypeAlreadyExists error
	ErrorBugTypeAlreadyExists = errors.New("BugType already exists")

	// ErrorTicketAlreadyExits error
	ErrorTicketAlreadyExits = errors.New("ticket already exits")

	// ErrorVerificationKeyInvalid error
	ErrorVerificationKeyInvalid = errors.New("verification Key does not exist")
)

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
