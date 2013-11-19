package controllers

import (
	"github.com/robfig/revel"
)

type Home struct {
	BaseController
}

func (c Home) Main() revel.Result {
	return c.Render()
}

func (c Home) Index() revel.Result {
	return c.Render()
}
