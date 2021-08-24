package appsessions

import (
	"github.com/gorilla/sessions"

	"github.com/yohagos/ticketSystem/utils"
)

var Store *sessions.CookieStore

func SessionInit() {
	authKeyOne := utils.GenerateVerificationKey()

	Store = sessions.NewCookieStore([]byte(authKeyOne))
	Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 120,
	}
}
