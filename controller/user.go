package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/entity"
	"github.com/go-zelus/api-layout/service"
)

type UserController struct {
	Controller
}

// Hello
// @Summary Hello
// @Description Hello
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body entity.HelloRequest true "用户名"
// @Success 200 {object} entity.HelloResponse ""
// @Failure 400 {object} state.Error ""
// @Router /users/user [post]
func (c *UserController) Hello(ctx *gin.Context) {
	user := &entity.HelloRequest{}
	if ok, err := c.BindAndValid(ctx, user); !ok {
		c.JSON(ctx, err)
		return
	}
	us := service.NewUserService()
	res, err := us.Hello(ctx, user)
	if err != nil {
		c.JSON(ctx, err)
		return
	}
	c.JSON(ctx, res)
}
