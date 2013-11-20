package controllers

import (
	"encoding/json"
	"github.com/bondowe/bitoola.polls/app/models"
	"github.com/bondowe/bitoola.polls/app/repositories"
	"github.com/bondowe/bitoola.polls/app/viewmodels"
	"github.com/robfig/revel"
	"time"
)

type Account struct {
	BaseController
}

func (c Account) SignUpForm() revel.Result {
	return c.Render()
}

func (c Account) SignUp() revel.Result {

	var form = viewmodels.SignUpViewModel{}
	_ = json.NewDecoder(c.Request.Body).Decode(&form)

	form.Validate(c.Validation)

	if c.Validation.HasErrors() {
		return c.RenderJson(c.Validation.Errors)
	}

	user := models.User{
		Firstname:   form.Firstname,
		Lastname:    form.Lastname,
		Alias:       form.Alias,
		Gender:      form.Gender,
		DateOfBirth: time.Date(form.DateOfBirth.Year, time.Month(form.DateOfBirth.Month), form.DateOfBirth.Day, 0, 0, 0, 0, time.UTC),
		Email:       form.Email,
		CreatedAt:   time.Now()}

	repositories.CreateUser(&user)

	return c.RenderJson(c.Validation.Errors)
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
