package repositories

import (
	"github.com/bondowe/bitoola.polls/app/models"
)

func CreateUser(user *models.User) {
	session := getSession()
	defer session.Close()

	err := session.DB("bitoola-polls").C("users").Insert(user)

	if err != nil {
		panic(err)
	}
}

func UpdateUser(user *models.User) {

}
