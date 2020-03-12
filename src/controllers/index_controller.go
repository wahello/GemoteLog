package controllers

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/neffos"
)

type IndexController struct {
	Ctx     iris.Context
	Server  *neffos.Server
}

func (c *IndexController) Get() {

	c.Ctx.ServeFile("./public/views/index.html", false)
}

func getIP(req *http.Request) string {

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err == nil {
		return  ip
	}

	forward := req.Header.Get("X-Forwarded-For")
	return forward

}

func (c *IndexController) Post() mvc.Result {
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
