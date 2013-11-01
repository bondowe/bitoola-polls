package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Answer struct {
	Id            bson.ObjectId `bson: "_id"`
	QuestionId    bson.ObjectId
	Body          string
	Note          string
	ImageUrl      string
	Order         uint8
	CreatedBy     bson.ObjectId
	CreatedAt     time.Time
	LastUpdatedBy bson.ObjectId
	LastUpdatedAt time.Time
}
