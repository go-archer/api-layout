package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zelus/zelus/core/util"
)

func Unique() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get("X-Request-ID")
		if requestId == "" {
			requestId = util.NewUUID()
		}
		ctx.Set("X-Request-ID", requestId)
		ctx.Writer.Header().Set("X-Request-ID", requestId)
		ctx.Next()
	}
}
