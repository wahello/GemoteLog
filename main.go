package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/websocket"
	"github.com/ytlvy/gemote/src/controllers"
	"github.com/ytlvy/gemote/src/utils"
	"time"
)


func newApp() *iris.Application {

	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Load the template files.
	// tmpl := iris.HTML("./public/views", ".html").
	// 	// 	Layout("shared/layout.html").
	// 	Reload(true)
	// app.RegisterView(tmpl)

	app.RegisterView(iris.HTML("./public/views", ".html"))

	//固定资源
	app.HandleDir("/asset", "./public/asset")

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		_ = ctx.View("error.html")
	})

	ws := new(utils.WebsocketManage).Handler()
	app.Get("/ws", websocket.Handler(ws))


	// "/user" based mvc application.
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})

	//index
	routeIndex := mvc.New(app.Party("/", adminMiddleware))
	routeIndex.
		Register(
		sessManager.Start,
		ws,
		).
		Handle(&controllers.IndexController{})

	return app
}


func main() {
	app := newApp()

	//run server
	_ = app.Run(
		iris.Addr(":8080"),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}

func adminMiddleware(ctx iris.Context) {
	// [...]
	ctx.Next() // to move to the next handler, or don't that if you have any auth logic.
}

func notFoundHandler(ctx iris.Context) {
	_, _ = ctx.HTML("Custom route for 404 not found http code, here you can render a view, html, json <b>any valid response</b>.")
}

