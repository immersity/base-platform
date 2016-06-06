package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID        bson.ObjectId `bson:"_id"       json:"id"`
	Role      string        `bson:"role"      json:"role"`
	Email     string        `bson:"email"     json:"email"`
	Password  string        `bson:"password"  json:"password"`
	FirstName string        `bson:"firstName" json:"first-name"`
	LastName  string        `bson:"lastName"  json:"last-name"`
	Active    bool          `bson:"active"    json:"active"`
	CreatedOn time.Time     `bson:"createdOn" json:"created-on"`
}
