package repositories

import (
	"github.com/bondowe/bitoola.polls/app/models"
	"time"
)

func CreateUser(fname, lname, alias, gender, email, passwdSalt, passwdHash string, dob time.Time) {
	session := getSession()
	defer session.Close()

	user := models.User{
		Firstname:    fname,
		Lastname:     lname,
		Alias:        alias,
		Gender:       gender,
		DateOfBirth:  dob,
		Email:        email,
		PasswordSalt: passwdSalt,
		PasswordHash: passwdHash,
		Status:       models.USUnconfirmed,
		CreatedAt:    time.Now()}

	err := session.DB(DBBitoolaPolls).C(ColUsers).Insert(user)

	if err != nil {
		panic(err)
	}
}

func QueueEmail(from, to, subject, body string, sendAt time.Time) {
	session := getSession()
	defer session.Close()

	email := models.Email{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
		SendAt:  time.Now()}

	err := session.DB(DBBitoolaPolls).C(ColEmails).Insert(email)

	if err != nil {
		panic(err)
	}
}

func UpdateUser(user *models.User) {

}
