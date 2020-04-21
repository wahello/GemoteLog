package controllers

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/neffos"
)

type DebugController struct {
	Ctx    iris.Context
	Server *neffos.Server
}

func (c *DebugController) Get() {

	c.Ctx.ServeFile("./public/views/debug.html", false)
}

func (c *DebugController) Post() mvc.Result {
	clientip := getIP(c.Ctx.Request())
	log.Printf(" remote IP [%s]", c.Ctx.RemoteAddr())

	rawBodyAsBytes, err := ioutil.ReadAll(c.Ctx.Request().Body)
	if err != nil { /* handle the error */
		c.Ctx.Writef("%v", err)
	}

	rawBodyAsString := string(rawBodyAsBytes)
	//println(rawBodyAsString)
	msgs := strings.Join([]string{clientip, rawBodyAsString}, " =-= ")

	c.Server.Broadcast(nil, neffos.Message{Body: []byte(msgs)})
	return mvc.Response{
		Err: nil,
	}

}
