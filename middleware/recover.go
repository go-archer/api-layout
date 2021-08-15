package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/entity/state"
	"github.com/go-zelus/zelus/core/config"
	"github.com/go-zelus/zelus/core/email"
	"github.com/go-zelus/zelus/core/logx"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mailer := email.New()
		defer func() {
			if err := recover(); err != nil {
				logx.Panicf("panic recover err: %v", err)
				mailer.SendMail(config.GetStringSlice("email.to"),
					fmt.Sprintf("程序异常,发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				ctx.SecureJSON(state.ServerError.StatusCode(), state.ServerError.Response())
				ctx.Abort()
			}
		}()
	}
}
