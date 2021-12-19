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

type BatchController struct {
	session *mgo.Session
}

// Create New User Controller with mongodb session
func NewBatchController(s *mgo.Session) *BatchController {
	return &BatchController{
		session: s,
	}
}

// Get all user list
func (uc BatchController) GetAllBatchs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// Get teachers
	var teachers []models.Batch

	// Get all teachers
	if err := uc.session.DB("mongo-golang").C("teachers").Find(nil).All(&teachers); err != nil {
		fmt.Println("Error finding all teachers")
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(teachers)

	// Write content-type, statuscode, payload
	fmt.Fprintf(w, "%s", uj)
}

// Get User
func (uc BatchController) GetBatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.Batch{}

	// Fetch user
	if err := uc.session.DB("mongo-golang").C("teachers").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// Create User
func (uc BatchController) CreateBatch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.Batch{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.ID = bson.NewObjectId()

	// Write the user to mongo
	uc.session.DB("mongo-golang").C("teachers").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// Delete User
func (uc BatchController) DeleteBatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB("mongo-golang").C("teachers").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
