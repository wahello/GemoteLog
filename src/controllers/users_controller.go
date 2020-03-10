package controllers

import (
	"github.com/ytlvy/gemote/src/datamodels"
	"github.com/ytlvy/gemote/src/services"
	"github.com/kataras/iris/v12"
)

type UsersController struct {
	Ctx iris.Context
	Service services.UserService
}

func (c *UsersController) Get() (results []datamodels.User){
	return c.Service.GetAll()
}

func (c *UsersController)GetBy(id int64) (user datamodels.User, found bool)  {
	u, found := c.Service.GetByID(id)
	if found {
		c.Ctx.Values().Set("message", "User couldn't be found!")
	}

	return u, found
}

func (c *UsersController) PutBy(id int64) (datamodels.User, error)  {
	u := datamodels.User{}
	if err := c.Ctx.ReadForm(&u); err != nil {
		return u, err
	}

	return c.Service.Update(id, u)
}

func (c *UsersController) DeleteBy(id int64) interface{}  {
	isDel := c.Service.DeleteByID(id)
	if isDel {
		// 创建一个 key 为 string value 为任意的 字典结构 map[string] interface{}
		return map[string] interface{}{"deleted": id}
	}

	return iris.StatusBadRequest

}

