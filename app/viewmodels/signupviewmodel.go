package viewmodels

import (
	"time"
)

type SignUpViewModel struct {
	Firsname     string
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
