package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID        bson.ObjectId `bson:"_id"       json:"id"`
	Email     string        `bson:"email"     json:"email"`
	Password  string        `bson:"password"  json:"password"`
	CreatedOn time.Time     `bson:"createdOn" json:"createdOn"`
}
