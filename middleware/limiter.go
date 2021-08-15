package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zelus/api-layout/entity/state"
	"github.com/go-zelus/zelus/core/limiter"
)

func RateLimiter(l limiter.ILimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				ctx.SecureJSON(state.TooManyRequests.StatusCode(), state.TooManyRequests)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
