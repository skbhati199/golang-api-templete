package models

import "gopkg.in/mgo.v2/bson"

type Batch struct {
	ID   bson.ObjectId `json:"ID",bson:"_id"`
	Name string        `json:"Name" bson:"Name"`
}
