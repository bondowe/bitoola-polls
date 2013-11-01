package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type UserAnswer struct {
	Id         bson.ObjectId `bson: "_id"`
	UserId     bson.ObjectId
	PollId     bson.ObjectId
	QuestionId bson.ObjectId
	AnswerId   bson.ObjectId
	CreatedAt  time.Time
}
