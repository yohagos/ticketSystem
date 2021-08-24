package routes

import (
	"net/http"
	"time"

	"github.com/yohagos/ticketSystem/apperrors"
	"github.com/yohagos/ticketSystem/models"
	"github.com/yohagos/ticketSystem/utils"
)

// BugtypeGETHandler func
func BugtypeGETHandler(w http.ResponseWriter, r *http.Request) {
	user := CheckCurrentSession(r)
	utils.ExecuteTemplate(w, "bugtypes.html", struct {
		User string
	}{
		User: user,
	})
}

// BugtypePOSTHandler func
func BugtypePOSTHandler(w http.ResponseWriter, r *http.Request) {
	sessionKey := CheckCurrentSession(r)
	if sessionKey == "" {
		utils.ExecuteTemplate(w, "bugtypes.html", apperrors.ErrorSessionInvalid)
		time.Sleep(15 * time.Second)
		http.Redirect(w, r, "/", 302)
	}
	r.ParseForm()

	bugtypeAcronym := r.PostForm.Get("acronym")
	bugtypeName := r.PostForm.Get("name")
	bugtypeDescription := r.PostForm.Get("description")

	ok := models.BugTypeExists(bugtypeAcronym)

	if ok {
		utils.ExecuteTemplate(w, "bugtypes.html", ok)
		return
	}

	var newBugType models.BugTypes
	newBugType.Acronym = bugtypeAcronym
	newBugType.Name = bugtypeName
	newBugType.Description = bugtypeDescription

	err := newBugType.CreateNewBugType()
	if err != nil {
		utils.ExecuteTemplate(w, "bugtypes.html", err)
	}

	SaveCurrentSession(w, r, sessionKey)

	redirect := "/profile/" + sessionKey
	http.Redirect(w, r, redirect, 302)
}
