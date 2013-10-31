package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Question struct {
	_Id           bson.ObjectId `bson: "_id"`
	PollId        bson.ObjectId
	Body          string
	Header        string
	Footer        string
	MinAnswers    uint8
	MaxAnswers    uint8
	ImageUrl      string
	CreatedBy     bson.ObjectId
	CreatedAt     time.Time
	LastUpdatedBy bson.ObjectId
	LastUpdatedAt time.Time
}
