package utils

import (
	"github.com/kataras/iris/v12/sessions"
	"sync"
	"time"
)

var (
	instance *LoginUtil
	once sync.Once
	cookieNameForSessionID = "ytlvy.com"
	Sess = sessions.New(sessions.Config{
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

//func UpdateSession(session *sessions.Session)  {
//	c.session = session
//}

const userIDKey = "UserID"
func GetCurrentUserID(session *sessions.Session) int64 {
	userID := session.GetInt64Default(userIDKey, 0)
	return userID
}

func IsLoggedIn(session *sessions.Session) bool  {
	return GetCurrentUserID(session) > 0
}


func Logout(session *sessions.Session)  {
	session.Destroy()
}

func UpdateUserID(id int64, session *sessions.Session)  {
	session.Set(userIDKey, id)
}