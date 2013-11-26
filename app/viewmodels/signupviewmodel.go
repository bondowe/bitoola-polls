package viewmodels

import (
	"github.com/robfig/revel"
	"regexp"
	"strings"
)

type SignUpViewModel struct {
	Firstname    string
	Lastname     string
	Alias        string
	Gender       string
	Email        string
	EmailConfirm string
	DateOfBirth  struct {
		Day   int
		Month int
		Year  int
	}
	Password string
}

func (form *SignUpViewModel) Validate(v *revel.Validation) {
	v.Required(form.Firstname)
	v.MinSize(form.Firstname, 2)
	v.MaxSize(form.Firstname, 20)

	v.Required(form.Lastname)
	v.MinSize(form.Lastname, 2)
	v.MaxSize(form.Lastname, 20)

	v.Required(form.Alias)
	v.MinSize(form.Alias, 3)
	v.MaxSize(form.Alias, 10)
	v.Match(form.Alias, regexp.MustCompile(`\w+(:?\d)*`))

	v.Required(form.Gender)
	v.Match(form.Gender, regexp.MustCompile(`^F|M$`))

	v.Required(form.Email)
	v.Email(form.Email)

	v.Required(strings.EqualFold(form.Email, form.EmailConfirm)).Key("form.EmailConfirm")

	v.Required(form.DateOfBirth.Day)
	v.Range(form.DateOfBirth.Day, 1, 31)

	v.Required(form.DateOfBirth.Month)
	v.Range(form.DateOfBirth.Month, 1, 12)

	v.Required(form.DateOfBirth.Year)

	v.Required(form.Password)
	v.MinSize(form.Password, 5)
	v.MaxSize(form.Password, 20)
}
