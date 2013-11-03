package controllers

import (
	"github.com/robfig/revel"
	"runtime"
	"strings"
)

type BaseController struct {
	*revel.Controller
}

func (c BaseController) RenderView(extraRenderArgs ...interface{}) revel.Result {
	// Get the calling function name.
	_, _, line, ok := runtime.Caller(1)
	if !ok {
		revel.ERROR.Println("Failed to get Caller information")
	}

	// Get the extra RenderArgs passed in.
	if renderArgNames, ok := c.MethodType.RenderArgNames[line]; ok {
		if len(renderArgNames) == len(extraRenderArgs) {
			for i, extraRenderArg := range extraRenderArgs {
				c.RenderArgs[renderArgNames[i]] = extraRenderArg
			}
		} else {
			revel.ERROR.Println(len(renderArgNames), "RenderArg names found for",
				len(extraRenderArgs), "extra RenderArgs")
		}
	} else {
		revel.ERROR.Println("No RenderArg names found for Render call on line", line,
			"(Method", c.MethodType.Name, ")")
	}
	loc := strings.Split(c.Request.Locale, "-")[0]
	return c.RenderTemplate("skins" + "/" + "default" + "/" + loc + "/" + c.Name + "/" + c.MethodType.Name + ".html")
}
