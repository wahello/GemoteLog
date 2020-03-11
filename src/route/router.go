package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/websocket"

	"github.com/ytlvy/gemote/src/controllers"
	"github.com/ytlvy/gemote/src/utils"

	//"github.com/ytlvy/gemote/src/repositories"
	"github.com/ytlvy/gemote/src/services"

	//"time"
)

type mainRouter struct {
	app *iris.Application
	session *sessions.Sessions

}

func New(app *iris.Application, session *sessions.Sessions) *mainRouter {
	return &mainRouter{
		app: app,
		session:session,
	}
}

func (r *mainRouter) Index() {
	//websocket
	ws := new(utils.WebsocketUtil).Handler()
	r.app.Get("/ws", websocket.Handler(ws))

	//index
	routeIndex := mvc.New(r.app.Party("/", adminMiddleware))
	routeIndex.
		Register(
			r.session.Start,
			ws,
		).
		Handle(&controllers.IndexController{})
}

func (r *mainRouter) Users(){
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
	sess := utils.Sess.Start(ctx)
	if !utils.IsLoggedIn(sess) {
		ctx.Redirect("/user/login")
		return
	}

	ctx.Next() // to move to the next handler, or don't that if you have any auth logic.
}