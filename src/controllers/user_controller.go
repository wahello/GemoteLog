package controllers

import (
	"github.com/ytlvy/gemote/src/datamodels"
	"github.com/ytlvy/gemote/src/services"
	"log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)


type UserController struct {
	Ctx iris.Context
	Service services.UserService
	Session *sessions.Session
}

const userIDKey = "UserID"

func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool  {
	return c.getCurrentUserID() > 0
}

func (c *UserController) logout()  {
	c.Session.Destroy()

}

var registerStaticView = mvc.View{
	Name:"user/register.html",
	Data:iris.Map{"Title": "User Registration"},
}

func (c *UserController) GetRegister() mvc.Result {
	log.Print(c.Ctx)

	if c.isLoggedIn() {
		c.logout()
	}

	return registerStaticView
}

func (c *UserController) PostRegister() mvc.Result {
	var (
		firstname = c.Ctx.FormValue("firstname")
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, err := c.Service.Create(password, datamodels.User{
		Username: username,
		Firstname:firstname,
	})

	c.Session.Set(userIDKey, u.ID)

	return mvc.Response {
		Err:err,
		Path:"/user/me",
	}

}