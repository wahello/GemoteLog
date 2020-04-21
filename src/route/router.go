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
	//websocket
	//index
	mvc.New(r.app.Party("/", adminMiddleware)).
		Register(
			r.session.Start,
			r.ws,
		).
		Handle(&controllers.IndexController{})
}

func (r *mainRouter) Network() {
	//websocket
	// ws := new(utils.WebsocketUtil).Handler()
	// r.app.Get("/ws", websocket.Handler(ws))

	//index
	mvc.New(r.app.Party("/network", adminMiddleware)).Register(
		r.session.Start,
		r.ws,
	).
		Handle(&controllers.NetworkController{})
}

func (r *mainRouter) Debug() {
	//websocket
	// ws := new(utils.WebsocketUtil).Handler()
	// r.app.Get("/ws", websocket.Handler(ws))

	//index
	mvc.New(r.app.Party("/debug", adminMiddleware)).
		Register(
			r.session.Start,
			r.ws,
		).
		Handle(&controllers.DebugController{})
}

func (r *mainRouter) About() {
	//index
	mvc.New(r.app.Party("/about", adminMiddleware)).
		Register(
			r.session.Start,
		).
		Handle(&controllers.AboutController{})
}

func (r *mainRouter) Users() {
	users := mvc.New(r.app.Party("/users"))
	//users.Router.Use(middleware.BasicAuth)
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
	// if !utils.GetLoginInstance().IsLoggedIn() {
	// 	ctx.Redirect("/user/login")
	// 	return
	// }

	ctx.Next() // to move to the next handler, or don't that if you have any auth logic.
}
