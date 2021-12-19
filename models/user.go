package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID     bson.ObjectId `json:"ID,bson:"_id"`
	Name   string        `json:"Name" bson:"Name"`
	Gender string        `json:"Gender" bson:"Gender"`
	Age    int           `json:"Age" bson:"Age"`
}
