package middleware

import (
	"net/http"

	"github.com/yohagos/ticketSystem/appsessions"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

// AuthRequired func - checking User Authorization
func AuthRequired(handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := appsessions.Store.Get(r, "session")
		_, ok := session.Values["username"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler(w, r)
	}
}
