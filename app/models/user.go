package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Firstname    string
	Lastname     string
	Alias        string
	Gender       string
	DateOfBirth  time.Time
	Email        string
	Country      string
	PasswordSalt string
	PasswordHash string
	CreatedAt    time.Time
}
