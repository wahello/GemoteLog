package controllers

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/neffos"
)

type AndroidController struct {
	Ctx    iris.Context
	Server *neffos.Server
}

func (c *AndroidController) Get() {

	c.Ctx.ServeFile("./public/views/index.html", false)
}

func (c *AndroidController) Post() mvc.Result {
	clientip := getIP(c.Ctx.Request())
	log.Printf(" remote IP [%s]", c.Ctx.RemoteAddr())

	rawBodyAsBytes, err := ioutil.ReadAll(c.Ctx.Request().Body)
	if err != nil { /* handle the error */
		c.Ctx.Writef("%v", err)
	}

	decoded, _ := base64.StdEncoding.DecodeString(string(rawBodyAsBytes))
	//println(rawBodyAsString)
	msgs := strings.Join([]string{clientip, string(decoded)}, " =-= ")

	c.Server.Broadcast(nil, neffos.Message{Body: []byte(msgs)})
	return mvc.Response{
		Err: nil,
	}

}
