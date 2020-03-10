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

	return make(map[int64]datamodels.User), nil
}