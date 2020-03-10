package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/ytlvy/gemote/src/datasource"
	"github.com/ytlvy/gemote/src/repositories"
	"github.com/ytlvy/gemote/src/route"
	"github.com/ytlvy/gemote/src/services"
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

	//Load the template files.
	tmpl := iris.HTML("./public/views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)


	//固定资源
	app.HandleDir("/asset", "./public/asset")

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		_ = ctx.View("error.html")
	})

	service := userService(app)

	// session
	sessManage := sessions.New(sessions.Config{
		Cookie:"ytlvy.com",
		Expires: 24 * time.Hour,
	})


	router := route.New(app, sessManage)
	router.Index()
	router.Users()
	router.User(service)

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


func notFoundHandler(ctx iris.Context) {
	_, _ = ctx.HTML("Custom route for 404 not found http code, here you can render a view, html, json <b>any valid response</b>.")
}

