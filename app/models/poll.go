package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Poll struct {
	Id               bson.ObjectId
	ShortDescription string
	LongDescription  string
	ImageUrl         string
	ValidFrom        time.Time
	ValidTo          time.Time
	CreatedBy        bson.ObjectId
	CreateAt         time.Time
	LastUpdatedBy    bson.ObjectId
	LastUpdatedAt    time.Time
}
