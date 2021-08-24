package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/yohagos/ticketSystem/appsessions"
	"github.com/yohagos/ticketSystem/databases"
	"github.com/yohagos/ticketSystem/mails"
	"github.com/yohagos/ticketSystem/models"
	"github.com/yohagos/ticketSystem/utils"
)

// RegistrationGETHandler func
func RegistrationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.html", nil)
}

// RegistrationPOSTHandler func
func RegistrationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := utils.GenerateVerificationKey()
	mail := r.PostForm.Get("email")
	name := r.PostForm.Get("name")

	if ok := databases.CheckUserExists(name); ok {
		utils.ExecuteTemplate(w, "registration.html", "email already in use")
	}

	var verificationUser models.UserVerification

	verificationUser.SetUserVerificationName(name)
	verificationUser.SetUserVerificationLastname(r.PostForm.Get("lastname"))
	verificationUser.SetUserVerificationPassword(r.PostForm.Get("password"))
	verificationUser.SetUserVerificationEmail(mail)
	verificationUser.SetUserVerificationGeneratedKey(key)

	go verificationUser.CreateVerificationProfile()

	go mails.SendVerificationMail(name, mail, key)

	http.Redirect(w, r, "/verification", 302)
}

// LoginGETHandler func
func LoginGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoginPOSTHandler func
func LoginPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.UserAuthentification(username, password)

	if err != nil {
		utils.ExecuteTemplate(w, "login.html", err)
	}

	SaveCurrentSession(w, r, username)

	redirectString := "/profile/" + username

	http.Redirect(w, r, redirectString, 302)
}

// LogoutGETHandler func
func LogoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := appsessions.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

// ProfileGETHandler func
func ProfileGETHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["user"]

	if strings.EqualFold("favicon.ico", username) {
		return
	}
	ok := models.UserExists(username)
	if !ok {
		log.Println(ok)
		return
	}
	user, err := models.UserGetAllInformations(username)
	if err != nil {
		log.Println(err)
	}

	ticketsList, err := models.GetTicketsByUser(username)
	if err != nil {
		log.Println(err)
	}

	displayTickets := false
	if len(*ticketsList) > 0 {
		displayTickets = true
	}

	utils.ExecuteTemplate(w, "profile.html", struct {
		User           *models.User
		Tickets        *[]models.Tickets
		DisplayTickets bool
	}{
		User:           user,
		Tickets:        ticketsList,
		DisplayTickets: displayTickets,
	})
}
