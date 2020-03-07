package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/neffos"
)

type IndexController struct {
	Ctx iris.Context
	server *neffos.Server
	Session *sessions.Session
}

func NewIndex(server *neffos.Server) *IndexController{

	return &IndexController{server: server}
}


func (c *IndexController) Get() mvc.Result {

	return mvc.View{
		Name: "index.html",
		// Data: map[string]interface{}{
		// "Title": "Hello Page",
		// },
	}
}

func (c *IndexController) Post() mvc.Result{

	//rawBodyAsBytes, err := ioutil.ReadAll(ctx.Request().Body)
	//if err != nil { /* handle the error */ ctx.Writef("%v", err) }
	//
	//rawBodyAsString := string(rawBodyAsBytes)
	//println(rawBodyAsString)
	//
	////c.Ctx.Request().Body
	//fmt.Print(c.Ctx.Request().Body)
	return mvc.Response{
		Err: nil,
	}

}
