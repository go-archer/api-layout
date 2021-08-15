package middleware

import (
	"bytes"
	"time"

	"github.com/go-zelus/zelus/core/logx"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBufferString(""),
		}
		ctx.Writer = bodyWriter
		begin := time.Now()
		ctx.Next()
		end := time.Now()
		duration := time.Since(begin) / 1e6
		logx.Infof("path: %s,method: %s,status_code: %d,begin_time: %d,end_time:%d,duration: %d ms",
			ctx.FullPath(),
			ctx.Request.Method,
			bodyWriter.Status(),
			begin.Unix(),
			end.Unix(),
			duration,
		)
	}
}
