package middleware

import (
	"tiant-micro/pkg/gerror"
	"tiant-micro/pkg/resp"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// API Limiter
func Limiter(limiter *rate.Limiter) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		traceID := gtx.GetString("trace_id")
		if !limiter.Allow() {
			err := gerror.ToManyRequest(nil)
			gtx.JSON(resp.ErrorResp(429, err.Code, err.HttpMsg, traceID))
			gtx.Abort()
			return
		}
		gtx.Next()
	}
}