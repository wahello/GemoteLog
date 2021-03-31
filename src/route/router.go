package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/neffos"

	"github.com/ytlvy/gemote/src/controllers"

	//"github.com/ytlvy/gemote/src/repositories"
	"github.com/ytlvy/gemote/src/services"
	//"time"
)

type mainRouter struct {
	app     *iris.Application
	session *sessions.Sessions
	ws      *neffos.Server
}

func New(app *iris.Application, session *sessions.Sessions, ws *neffos.Server) *mainRouter {
	return &mainRouter{
		app:     app,
		session: session,
		ws:      ws,
	}
}

func (r *mainRouter) Index() {
	//index
	mvc.New(r.app.Party("/", adminMiddleware)).
		Register(
			r.session.Start,
			r.ws,
		).
		Handle(&controllers.IndexController{})
	mvc.New(r.app.Party("/music.yl", adminMiddleware)).
		Register(
			r.session.Start,
			r.ws,
		).
		Handle(&controllers.AndroidController{})
}

func (r *mainRouter) Network() {
	mvc.New(r.app.Party("/network", adminMiddleware)).Register(
		r.session.Start,
		r.ws,
	).
		Handle(&controllers.NetworkController{})
}

func (r *mainRouter) Debug() {
	mvc.New(r.app.Party("/debug", adminMiddleware)).
		Register(
			r.session.Start,
			r.ws,
		).
		Handle(&controllers.DebugController{})
}

func (r *mainRouter) About() {
	mvc.New(r.app.Party("/about", adminMiddleware)).
		Register(
			r.session.Start,
		).
		Handle(&controllers.AboutController{})
}

func (r *mainRouter) Users() {
	users := mvc.New(r.app.Party("/users"))
	users.Handle(&controllers.UsersController{})
}

func (r *mainRouter) User(userService services.UserService) {
	user := mvc.New(r.app.Party("/user"))
	user.Register(
		userService,
		r.session.Start,
	)
	user.Handle(&controllers.UserController{})
}

func adminMiddleware(ctx iris.Context) {
	// if ctx.Request().Method != "POST" && !utils.GetLoginInstance().IsLoggedIn() {
	// 	ctx.Redirect("/user/login")
	// 	return
	// }

	ctx.Next() // to move to the next handler, or don't that if you have any auth logic.
}
