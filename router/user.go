package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/controller"
)

func userRouter(r *gin.RouterGroup) {
	g := r.Group("/users")
	g.POST("/user", new(controller.UserController).Hello)
}
