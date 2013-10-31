package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type UserActivity struct {
	_id       bson.ObjectId `bson: "_id"`
	UserId    bson.ObjectId
	SessionId string
	IpAddress string
	UserAgent string
	PollId    bson.ObjectId
	CreatedAt time.Time
}
