package datasource

import (
	"errors"
	"github.com/ytlvy/gemote/src/datamodels"
)

type Engine uint32

const (
	Memory Engine = iota
	Bolt
	MySQL
)

func LoadUsers(engine Engine) (map[int64]datamodels.User, error) {
	if engine != Memory {
		return nil, errors.New("not implemented")
	}

	users := make(map[int64]datamodels.User)
	admin := datamodels.User{Firstname:"Y.t", Username:"admin", ID:1000}

	hashed, err := datamodels.GeneratePassword("admin")
	if err == nil {

		admin.HashedPassword = hashed
		users[0] = admin
	}

	return users, nil
}