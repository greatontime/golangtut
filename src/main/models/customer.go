package models
import "gopkg.in/mgo.v2/bson"
type (
	Customer struct{
		Id bson.ObjectId `json:"id" bson:"_id"`
		Name string "json:'name' bson:'name'"
		Address string "json:'address' bson:'address'"
	}
)