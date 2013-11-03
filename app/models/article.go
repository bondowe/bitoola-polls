package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Article struct {
	Id            bson.ObjectId `bson: "_id"`
	Title         string
	Subtitle      string
	Header        string
	Body          string
	Footer        string
	CreatedBy     bson.ObjectId
	CreateAt      time.Time
	LastUpdatedBy bson.ObjectId
	LastUpdatedAt time.Time
}
