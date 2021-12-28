package models

import "gopkg.in/mgo.v2/bson"

type Chat struct {
	id      bson.ObjectId `json:"ID",bson:"_id"`
	message string        `json:"message" bson:"message"`
}
