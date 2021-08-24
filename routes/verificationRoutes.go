package routes

import (
	"net/http"

	"github.com/yohagos/ticketSystem/apperrors"
	"github.com/yohagos/ticketSystem/models"
	"github.com/yohagos/ticketSystem/utils"
)

// VerificationGETHandler func
func VerificationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "verification.html", nil)
}

// VerificationPOSTHandler func
func VerificationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.PostForm.Get("verification")

	mail, ok := models.CheckVerification(key)
	if !ok {
		utils.ExecuteTemplate(w, "verification.html", struct {
			Error error
		}{
			Error: apperrors.ErrorVerificationKeyInvalid,
		})
	}

	models.CreateNewUser(mail)
	http.Redirect(w, r, "/login", 302)
}
