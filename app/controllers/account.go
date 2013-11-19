package controllers

import (
	"encoding/json"
	"github.com/bondowe/bitoola.polls/app/models"
	"github.com/bondowe/bitoola.polls/app/repositories"
	"github.com/bondowe/bitoola.polls/app/viewmodels"
	"github.com/robfig/revel"
)

type Account struct {
	BaseController
}

func (c Account) SignUpForm() revel.Result {
	return c.Render()
}

func (c Account) SignUp() revel.Result {

	var model = viewmodels.SignUpViewModel{}
	_ = json.NewDecoder(c.Request.Body).Decode(&model)

	model.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		// c.Flash.Error(msg, ...)
		return c.Redirect(Account.SignUpForm)
	}

	user := models.User{
		Firstname: model.Firstname,
		Lastname:  model.Lastname,
	}

	repositories.CreateUser(&user)

	return c.RenderJson(user)
}

func (c Account) SignInForm() revel.Result {
	return c.Render()
}

func (c Account) SignIn(model viewmodels.SignInViewModel) revel.Result {
	return c.Todo()
}

func (c Account) PasswordResetForm() revel.Result {
	return c.Render()
}

func (c Account) PasswordReset(model viewmodels.PasswordResetViewModel) revel.Result {
	return c.Todo()
}

func (c Account) PasswordUpdateForm() revel.Result {
	return c.Render()
}

func (c Account) PasswordUpdate(model viewmodels.PasswordResetViewModel) revel.Result {
	return c.Todo()
}

func (c Account) SignOut() revel.Result {
	return c.Todo()
}

func (c Account) Captcha() revel.Result {
	c.RenderArgs["lang"] = c.Lang()
	return c.Render()
}
