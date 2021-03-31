package main

// cross build for windows
// GOOS=windows GOARCH=amd64 go build -o gemote-win.exe
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gemote-linux

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/websocket"
	"github.com/ytlvy/gemote/src/utils"

	"flag"

	"github.com/ytlvy/gemote/src/datasource"
	"github.com/ytlvy/gemote/src/repositories"
	"github.com/ytlvy/gemote/src/route"
	"github.com/ytlvy/gemote/src/services"
)

func newApp() *iris.Application {

	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	//Load the template files.
	tmpl := iris.HTML("./public/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	//固定资源
	app.HandleDir("/asset", "./public/asset")

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		_ = ctx.View("error.html")
	})

	service := userService(app)

	sessionManager := utils.Sess
	app.Use(func(ctx iris.Context) {
		utils.GetLoginInstance().UpdateSession(ctx)
		//if ctx.Path() != "/user/login" {
		//}
		ctx.Next()
	})

	ws := new(utils.WebsocketUtil).Handler()
	app.Get("/ws", websocket.Handler(ws))

	router := route.New(app, sessionManager, ws)
	router.Index()
	router.Users()
	router.User(service)
	router.Network()
	router.Debug()
	router.About()

	return app
}

func userService(app *iris.Application) services.UserService {
	//service
	db, err := datasource.LoadUsers(datasource.Memory)
	if err != nil {
		app.Logger().Fatalf("error while loading the user: %v", err)
	}
	repo := repositories.NewUserRepository(db)
	return services.NewUserService(repo)
}

var port = flag.Int("p", 10088, "server port")

func main() {
	app := newApp()

	flag.Parse()
	//run server
	address := fmt.Sprintf(":%d", *port)
	_ = app.Run(
		iris.Addr(address),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}

func notFoundHandler(ctx iris.Context) {
	_, _ = ctx.HTML("Custom route for 404 not found http code, here you can render a view, html, json <b>any valid response</b>.")
}
