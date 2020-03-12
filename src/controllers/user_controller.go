package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/ytlvy/gemote/src/datamodels"
	"github.com/ytlvy/gemote/src/services"
	"github.com/ytlvy/gemote/src/utils"
)


type UserController struct {
	Ctx iris.Context
	Service services.UserService
}


var registerStaticView = mvc.View{
	Name:"user/register.html",
	Data:iris.Map{"Title": "User Registration"},
}


func (c *UserController) GetRegister() mvc.Result {


	if utils.GetLoginInstance().IsLoggedIn() {
		utils.GetLoginInstance().Logout()
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

	utils.GetLoginInstance().UpdateUserID(u.ID)

	return mvc.Response {
		Err:err,
		Path:"/",
	}

}

var loginStaticView = mvc.View{
	Name: "user/login.html",
	Data: iris.Map{"Title": "User Login"},
}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *UserController) GetLogin() mvc.Result {
	if utils.GetLoginInstance().IsLoggedIn() {
		// if it's already logged in then destroy the previous session.
		utils.GetLoginInstance().Logout()
	}

	return loginStaticView
}

// PostLogin handles POST: http://localhost:8080/user/register.
func (c *UserController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, found := c.Service.GetByUsernameAndPassword(username, password)

	if !found {
		return mvc.Response{
			Path: "/user/register",
		}
	}

	utils.GetLoginInstance().UpdateUserID(u.ID)

	return mvc.Response{
		Path: "/",
	}
}

// GetMe handles GET: http://localhost:8080/user/me.
func (c *UserController) GetMe() mvc.Result {
	if !utils.GetLoginInstance().IsLoggedIn() {
		// if it's not logged in then redirect user to the login page.
		return mvc.Response{Path: "/user/login"}
	}

	u, found := c.Service.GetByID(utils.GetLoginInstance().GetCurrentUserID())
	if !found {
		// if the  session exists but for some reason the user doesn't exist in the "database"
		// then logout and re-execute the function, it will redirect the client to the
		// /user/login page.
		utils.GetLoginInstance().Logout()
		return c.GetMe()
	}

	return mvc.View{
		Name: "user/me.html",
		Data: iris.Map{
			"Title": "Profile of " + u.Username,
			"User":  u,
		},
	}
}

// AnyLogout handles All/Any HTTP Methods for: http://localhost:8080/user/logout.
func (c *UserController) AnyLogout() {
	if utils.GetLoginInstance().IsLoggedIn() {
		utils.GetLoginInstance().Logout()
	}

	c.Ctx.Redirect("/user/login")
}
