package main

import (
	"log"
	"net/http"

	"github.com/yohagos/ticketSystem/appsessions"
	"github.com/yohagos/ticketSystem/databases"
	"github.com/yohagos/ticketSystem/mails"
	"github.com/yohagos/ticketSystem/routes"
	"github.com/yohagos/ticketSystem/utils"
)

func main() {
	log.Println("connecting to database...")
	databases.Init()

	mails.EmailInit()
	appsessions.SessionInit()

	log.Println("loading templates...")
	utils.LoadTemplate("./templates/*.html")

	r := routes.NewRouter()
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("server error! message : ", err)
	}
}
