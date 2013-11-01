package controllers

import (
	"github.com/robfig/revel"
)

type App struct {
	BaseController
}

func (c App) Index() revel.Result {
	fname := "Elondo Mb."
	return c.RenderView(fname)
}
