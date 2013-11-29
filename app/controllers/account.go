package controllers

import (
	"encoding/json"
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
		c.Response.Status = 418
		return c.RenderJson(c.Validation.Errors)
	}

	passwdSalt := randString(8)
	passwdHash := hashText(form.Password, passwdSalt)
	dob := time.Date(form.DateOfBirth.Year, time.Month(form.DateOfBirth.Month), form.DateOfBirth.Day, 0, 0, 0, 0, time.UTC)

	repositories.CreateUser(form.Firstname, form.Lastname, form.Alias, form.Gender, form.Email, passwdSalt, passwdHash, dob)

	from, _ := revel.Config.String("email.noreply.address")
	subject := c.Message("email.signup.subject")

	repositories.QueueEmail(from, form.Email, subject, "BODY", time.Now())

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
