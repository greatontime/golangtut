package models

import "gopkg.in/mgo.v2/bson"

type (
// User represents the structure of our resource
	Users struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Gender string        `json:"gender" bson:"gender"`
		Age    int           `json:"age" bson:"age"`
		Username string	`json:"username" bson:"username"`
		Password string	`json:"password" bson:"password"`
	}
)
