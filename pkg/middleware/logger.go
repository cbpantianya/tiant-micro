package middleware

import (
	"tiant-micro/pkg/random"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Trace struct {
	TraceID string `json:"trace_id"`
}

func Logger(sugar *zap.SugaredLogger) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		// generate trace_id
		traceID := random.NewTraceID()
		gtx.Set("trace_id", traceID)
		start := time.Now()
		gtx.Next()
		end := time.Now()
		sugar.Infow("req",
			zap.String("trace_id", traceID),
			zap.String("consuming", end.Sub(start).String()),
		)
	}
}
