package models

import (
	"fmt"
	"log"

	"../databases"
	"../utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tickets struct
type Tickets struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	CreatedBy string             `bson:"lastname,omitempty"`
	BugType   string             `bson:"password,omitempty"`
	Status    string             `bson:"status,omitempty"`
	Comments  []Comment          `bson:"comments,omitempty"`
	CreatedAt string             `bson:"createdAt,omitempty"`
	UpdatedAt string             `bson:"updatedAt,omitempty"`
}

// Comments struct
type Comment struct {
	User      string `bson:"user,omitempty"`
	Message   string `bson:"message,omitempty"`
	CreatedAt string `bson:"createdAt,omitempty"`
}

// TestCreateTicket func
/* func TestCreateTicket() {
	timestamp := utils.CreateTimeStamp()
	ticketDocument := bson.D{
		{Key: "name", Value: "test-12345"},
		{Key: "createdby", Value: "test01"},
		{Key: "bugtype", Value: "bug"},
		{Key: "status", Value: "testing"},
		{Key: "createdAt", Value: timestamp},
		{Key: "updatedAt", Value: timestamp},
	}

	databases.CreateNewTicket(ticketDocument)
} */

// GetTicketID func
func (ticket *Tickets) GetTicketID() string {
	return ticket.ID.String()
}

// GetTicketName func
func (ticket *Tickets) GetTicketName() string {
	return ticket.Name
}

// GetTicketCreatedBy func
func (ticket *Tickets) GetTicketCreatedBy() string {
	return ticket.CreatedBy
}

// GetTicketBugType func
func (ticket *Tickets) GetTicketBugType() string {
	return ticket.BugType
}

// GetTicketStatus func
func (ticket *Tickets) GetTicketStatus() string {
	return ticket.Status
}

// GetTicketAllComments func
func (ticket *Tickets) GetTicketAllComments() []Comment {
	return ticket.Comments
}

// GetTicketUpdatedAt func
func (ticket *Tickets) GetTicketUpdatedAt() string {
	return ticket.UpdatedAt
}

// GetTicketCreatedAt func
func (ticket *Tickets) GetTicketCreatedAt() string {
	return ticket.CreatedAt
}

// SetTicketName func
func (ticket *Tickets) SetTicketName(tic string) {
	ticket.Name = tic
}

// SetTicketCreatedBy func
func (ticket *Tickets) SetTicketCreatedBy(tic string) {
	ticket.CreatedBy = tic
}

// SetTicketBugType func
func (ticket *Tickets) SetTicketBugType(tic string) {
	ticket.BugType = tic
}

// SetTicketStatus func
func (ticket *Tickets) SetTicketStatus(tic string) {
	ticket.Status = tic
}

// SetTicketComment func
func (ticket *Tickets) SetTicketComment(user, message, createdAt string) {
	com := ticket.GetTicketAllComments()
	tic := Comment{user, message, createdAt}
	com = append(com, tic)
	ticket.Comments = com
}

// SetTicketUpdatedAt func
func (ticket *Tickets) SetTicketUpdatedAt(tic string) {
	ticket.UpdatedAt = tic
}

// SetTicketCreatedAt func
func (ticket *Tickets) SetTicketCreatedAt(tic string) {
	ticket.CreatedAt = tic
}

// CreateNewTicket func
func (ticket *Tickets) CreateNewTicket() {
	newname := ticket.GetTicketBugType() + "-" + utils.RandomFiveDigitNumber()
	ok := databases.CheckTicketExists(newname)
	if !ok {
		log.Println("Ticket already exists")
		return
	}
	/* num := utils.RandomFiveDigitNumber() */
	timeStamp := utils.CreateTimeStamp()
	ticketDocument := bson.D{
		{Key: "name", Value: newname},
		{Key: "createdby", Value: ticket.GetTicketCreatedBy()},
		{Key: "bugtype", Value: ticket.GetTicketBugType()},
		{Key: "status", Value: ticket.GetTicketStatus()},
		{Key: "comments", Value: ticket.GetTicketAllComments()},
		{Key: "createdAt", Value: timeStamp},
		{Key: "updatedAt", Value: timeStamp},
	}
	databases.CreateNewTicket(ticketDocument)
}

// NewTicketExists func
func NewTicketExists(name string) bool {
	return databases.CheckTicketExists(name)
}

// TicketGetAllInformations func
func TicketGetAllInformations(name string) (*Tickets, error) {
	var ticket Tickets

	result, err := databases.GetAllTicketInformations(name)
	if err != nil {
		log.Println(err)
		return &ticket, err
	}

	for k, v := range result {
		switch k {
		case "name":
			key := fmt.Sprintf("%v", v)
			ticket.SetTicketName(key)
		case "bugtype":
			key := fmt.Sprintf("%v", v)
			ticket.SetTicketBugType(key)
		case "createdby":
			key := fmt.Sprintf("%v", v)
			ticket.SetTicketCreatedBy(key)
		case "status":
			key := fmt.Sprintf("%v", v)
			ticket.SetTicketStatus(key)
		case "createdAt":
			key := fmt.Sprintf("%v", v)
			ticket.SetTicketCreatedAt(key)
		case "updatedAt":
			key := fmt.Sprintf("%v", v)
			ticket.SetTicketUpdatedAt(key)
		case "comments":
			key := fmt.Sprintf("%v", v)
			fmt.Println(key)
		default:

		}
	}

	return &ticket, nil
}

/* func sortComments(sort interface{}) Comments {
	for _,v := range sort {

	}
	return Comments{}
} */

// GetTicketsByUser func
func GetTicketsByUser(username string) (*[]Tickets, error) {
	ticketList, err := databases.AllTicketsByUser(username)
	if err != nil {
		return nil, err
	}

	return bsonToTicketConverter(ticketList), nil
}

func bsonToTicketConverter(list *[]bson.M) *[]Tickets {
	var ticketlist []Tickets

	for _, listValue := range *list {
		var tic Tickets

		for keys, values := range listValue {
			switch keys {
			case "name":
				key := fmt.Sprintf("%v", values)
				tic.SetTicketName(key)
			case "bugtype":
				key := fmt.Sprintf("%v", values)
				tic.SetTicketBugType(key)
			case "createdby":
				key := fmt.Sprintf("%v", values)
				tic.SetTicketCreatedBy(key)
			case "status":
				key := fmt.Sprintf("%v", values)
				tic.SetTicketStatus(key)
			case "createdAt":
				key := fmt.Sprintf("%v", values)
				tic.SetTicketCreatedAt(key)
			case "updatedAt":
				key := fmt.Sprintf("%v", values)
				tic.SetTicketUpdatedAt(key)
			default:
			}
		}
		ticketlist = append(ticketlist, tic)
	}

	return &ticketlist
}
