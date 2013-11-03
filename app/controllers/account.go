package controllers

import (
	"github.com/bondowe/bitoola.polls/app/viewmodels"
	"github.com/robfig/revel"
)

type Account struct {
	BaseController
}

func (c Account) SignUpForm() revel.Result {
	return c.RenderView()
}

func (c Account) SignUp(model viewmodels.SignUpViewModel) revel.Result {
	return c.Todo()
}

func (c Account) SignInForm() revel.Result {
	return c.Todo()
}

func (c Account) SignIn(model viewmodels.SignInViewModel) revel.Result {
	return c.Todo()
}

func (c Account) SignOut() revel.Result {
	return c.Todo()
}
