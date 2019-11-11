package user

import "github.com/globalsign/mgo/bson"

// User model
type User struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
