package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type PollComment struct {
	Id         bson.ObjectId `bson: "_id"`
	PollId     bson.ObjectId
	QuestionId bson.ObjectId
	Body       string
	CreatedBy  bson.ObjectId
	CreatedAt  time.Time
}
