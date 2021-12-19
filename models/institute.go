package models

import "gopkg.in/mgo.v2/bson"

type Institude struct {
	ID   bson.ObjectId `json:"ID",bson:"_id"`
	Name string        `json:"Name" bson:"Name"`
}
