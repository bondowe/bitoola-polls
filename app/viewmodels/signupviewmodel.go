package viewmodels

import (
	"github.com/robfig/revel"
	"regexp"
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
}

func (model *SignUpViewModel) Validate(v *revel.Validation) {
	v.Required(model.Firstname)
	v.MinSize(model.Firstname, 2)
	v.MaxSize(model.Firstname, 20)

	v.Required(model.Lastname)
	v.MinSize(model.Lastname, 2)
	v.MaxSize(model.Lastname, 20)

	v.Required(model.Alias)
	v.MinSize(model.Alias, 3)
	v.MaxSize(model.Alias, 10)
	v.Match(model.Alias, regexp.MustCompile(`\w+(:?\d)*`))

	v.Required(model.Gender)
	v.Match(model.Gender, regexp.MustCompile(`^F|M$`))

	v.Required(model.Email)
	v.Email(model.Email)

	v.Required(model.EmailConfirm)
	v.Email(model.EmailConfirm)

	v.Required(model.DateOfBirth.Day)
	v.Range(model.DateOfBirth.Day, 1, 31)

	v.Required(model.DateOfBirth.Month)
	v.Range(model.DateOfBirth.Month, 1, 12)

	v.Required(model.DateOfBirth.Year)
}
