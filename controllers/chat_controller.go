package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skbhati199/learningspacepro-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ChatController struct {
	session *mgo.Session
}

func NewChatController(s *mgo.Session) *ChatController {
	return &ChatController{
		session: s,
	}
}

// Get all Chats
func (uc ChatController) GetAllChats(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// Get chats
	var chats []models.Chat

	// Get all chats
	if err := uc.session.DB("mongo-golang").C("chats").Find(nil).All(&chats); err != nil {
		fmt.Println("Error finding all chats")
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(chats)

	// Write content-type, statuscode, payload
	fmt.Fprintf(w, "%s", uj)
}

// Get Chat by _id
func (uc ChatController) GetChat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Get chat
	var chat models.Chat

	// Get chat
	if err := uc.session.DB("mongo-golang").C("chats").FindId(oid).One(&chat); err != nil {
		fmt.Println("Error finding chat")
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(chat)

	// Write content-type, statuscode, payload
	fmt.Fprintf(w, "%s", uj)
}

// Delete Chat by _id
func (uc ChatController) DeleteChat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove chat
	if err := uc.session.DB("mongo-golang").C("chats").RemoveId(oid); err != nil {
		fmt.Println("Error removing chat")
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "Deleted chat")
}

// Update chat
func (uc ChatController) UpdateChat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Decode
	decoder := json.NewDecoder(r.Body)
	var chat models.Chat
	err := decoder.Decode(&chat)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Update chat
	if err := uc.session.DB("mongo-golang").C("chats").UpdateId(oid, &chat); err != nil {
		fmt.Println("Error updating chat")
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "Updated chat")
}

// Get All chat message
func (uc ChatController) GetAllChatMessages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Get chat
	var chat models.Chat

	// Get chat
	if err := uc.session.DB("mongo-golang").C("chats").FindId(oid).One(&chat); err != nil {
		fmt.Println("Error finding chat")
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(chat)

	// Write content-type, statuscode, payload
	fmt.Fprintf(w, "%s", uj)
}

// Create chat
func (uc ChatController) CreateChat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Decode
	decoder := json.NewDecoder(r.Body)
	var chat models.Chat
	err := decoder.Decode(&chat)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create chat
	if err := uc.session.DB("mongo-golang").C("chats").Insert(&chat); err != nil {
		fmt.Println("Error creating chat")
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(chat)

	// Write content-type, statuscode, payload
	fmt.Fprintf(w, "%s", uj)
}
