package controllers

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/robfig/revel"
	"runtime"
	"strings"
)

type BaseController struct {
	*revel.Controller
}

func (c BaseController) Render(extraRenderArgs ...interface{}) revel.Result {
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

	return c.RenderTemplate("skins" + "/" + "default" + "/" + c.Lang() + "/" + c.Name + "/" + c.MethodType.Name + ".html")
}

func (c BaseController) Lang() string {
	return strings.Split(c.Request.Locale, "-")[0]
}

func hashText(text, salt string) string {
	sh := sha1.New()
	sh.Write([]byte(salt + text))
	return fmt.Sprintf("%x", sh.Sum(nil))
}

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
