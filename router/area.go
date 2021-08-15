package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/controller"
)

func areaRouter(r *gin.RouterGroup) {
	g := r.Group("/areas")
	g.GET("/provinces", new(controller.AreaController).Provinces)
	g.GET("/cities/:code", new(controller.AreaController).Cities)
	g.GET("/districts/:code", new(controller.AreaController).Districts)
	g.GET("/streets/:code", new(controller.AreaController).Streets)
	g.GET("/committees/:code", new(controller.AreaController).Committees)
}
