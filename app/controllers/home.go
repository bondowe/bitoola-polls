package controllers

import (
	"github.com/robfig/revel"
)

type Home struct {
	BaseController
}

func (c Home) Default() revel.Result {
	return c.RenderView()
}

func (c Home) Index() revel.Result {
	return c.RenderView()
}
