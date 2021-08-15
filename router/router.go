package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/middleware"
)

func NoRoute(ctx *gin.Context) {
	ctx.SecureJSON(http.StatusNotFound, gin.H{"code": 404, "msg": "没有找到路由"})
}

func NoMethod(ctx *gin.Context) {
	ctx.SecureJSON(http.StatusNotFound, gin.H{"code": 404, "msg": "没有找到方法"})
}

func New() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(middleware.Recover())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Unique())
	r.Use(middleware.Translations())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.NoRoute(NoRoute)
	r.NoMethod(NoMethod)

	apiV1(r)

	return r
}

func apiV1(e *gin.Engine) {
	pre := e.Group("/api/v1")
	userRouter(pre)
}
