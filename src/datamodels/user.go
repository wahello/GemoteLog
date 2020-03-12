package datamodels

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int64    `json:"id" form:"id"`
	Firstname     string   `json:"firstname form:"firstname`
	Username      string   `json:"username" form:"username"`
	HashedPassword string  `json:"-" form:"-"`
	CreatedAt     time.Time  `json:"created_at" form:"created_at"`
}

func (u User) IsValid() bool  {
	return u.ID > 0
}

func GeneratePassword(userPassword string) (string, error)  {
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(hash), err
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error)  {
	if err:= bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err !=nil {
		return false, err
	}
	return  true, nil
}

