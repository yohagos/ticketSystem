package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"../apperrors"
	"../appsessions"
	"../models"
	"../utils"

	"github.com/gorilla/mux"
)

// TicketsGETHandler func
func TicketsGETHandler(w http.ResponseWriter, r *http.Request) {
	user := CheckCurrentSession(r)
	list, _ := models.BugTypeListOfAcronyms()

	utils.ExecuteTemplate(w, "tickets.html", struct {
		User string
		List []string
	}{
		User: user,
		List: list,
	})
}

// TicketsPOSTHandler func
func TicketsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	sessionKey := CheckCurrentSession(r)
	if sessionKey == "" {
		utils.ExecuteTemplate(w, "tickets.html", apperrors.ErrorSessionInvalid)
		time.Sleep(15 * time.Second)
		http.Redirect(w, r, "/", 302)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		bugtype := r.PostForm.Get("bugtype")
		status := r.PostForm.Get("status")

		var ticket models.Tickets
		ticket.SetTicketBugType(bugtype)
		ticket.SetTicketStatus(status)
		ticket.SetTicketCreatedBy(sessionKey)

		ticket.CreateNewTicket()

		SaveCurrentSession(w, r, sessionKey)

		redirect := "/profile/" + sessionKey

		http.Redirect(w, r, redirect, 302)
	}
}

// TicketDetailGETHandler func
func TicketDetailGETHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ticket, err := models.TicketGetAllInformations(id)
	if err != nil {
		log.Println(err)
	}

	user, err := models.UserGetAllInformations(ticket.GetTicketCreatedBy())
	if err != nil {
		log.Println(err)
	}

	/* comments := ticket.Comment
	var display bool
	if len(comments) > 0 {
		display = true
	} */

	utils.ExecuteTemplate(w, "ticketdetails.html", struct {
		Ticket          *models.Tickets
		User            *models.User
		DisplayComments bool
		/* Comments        []models.Comments */
	}{
		Ticket: ticket,
		User:   user,
		/* DisplayComments: display,
		Comments:        comments, */
	})
}

// TicketDetailPOSTHandler func
func TicketDetailPOSTHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	r.ParseForm()
	update := r.PostForm.Get("update")

	createdAt := utils.CreateTimeStamp()

	ticket, err := models.TicketGetAllInformations(id)
	if err != nil {
		log.Println(err)
	}

	session, _ := appsessions.Store.Get(r, "session")
	user := fmt.Sprintf("%v", session.Values["username"])

	ticket.SetTicketComment(user, update, createdAt)

	http.Redirect(w, r, "/ticket/detail/{id}", 302)
}
