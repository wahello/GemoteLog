package controllers

import (
	"github.com/kataras/iris/v12"
)

type AboutController struct {
	Ctx iris.Context
}

func (c *AboutController) Get() {

	c.Ctx.ServeFile("./public/views/about.html", false)
}
