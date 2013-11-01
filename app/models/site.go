package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Site struct {
	Id              bson.ObjectId `bson:"_id"`
	HostnamePattern string
	SkinFolder      string
	Description     string
	ValidFrom       time.Time
	ValidTo         time.Time
	CreatedAt       time.Time
}
