package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type ArticleComment struct {
	Id        bson.ObjectId `bson: "_id"`
	ArticleId bson.ObjectId
	Body      string
	CreatedBy bson.ObjectId
	CreatedAt time.Time
}
