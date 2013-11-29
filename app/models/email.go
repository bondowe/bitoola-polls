package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Email struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	From    string
	To      string
	Subject string
	Body    string
	SendAt  time.Time
	SentAt  time.Time
}
