package utils

import (
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	instance               *LoginUtil
	once                   sync.Once
	cookieNameForSessionID = "ytlvy.com"
	Sess                   = sessions.New(sessions.Config{
		Cookie:  cookieNameForSessionID,
		Expires: 24 * time.Hour,
	})
)

func GetLoginInstance() *LoginUtil {
	once.Do(func() {
		instance = &LoginUtil{}
	})
	return instance
}

type LoginUtil struct {
	session *sessions.Session
}

func (c *LoginUtil) UpdateSession(ctx iris.Context) {
	if ctx == nil {
		panic("ctx should not be nil")
		return
	}
	// c.session = Sess.Start(ctx)
}

const userIDKey = "UserID"

func (c *LoginUtil) GetCurrentUserID() int64 {
	if c.session == nil {
		panic("session should not be nil")
		return -1
	}

	userID := c.session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *LoginUtil) IsLoggedIn() bool {
	return c.GetCurrentUserID() > 0
}

func (c *LoginUtil) Logout() {
	if c.session == nil {
		panic("session should not be nil")
		return
	}
	c.session.Destroy()
}

func (c *LoginUtil) UpdateUserID(id int64) {
	if c.session == nil {
		panic("session should not be nil")
		return
	}
	c.session.Set(userIDKey, id)
}
