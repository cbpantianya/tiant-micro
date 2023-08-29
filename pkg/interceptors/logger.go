package interceptors

import (
	"context"
	"encoding/json"
	"tiant-micro/pkg/gerror"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Trace struct {
	TraceID string `json:"trace_id"`
}

func Logger(sugar *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		// trace_id
		rstr, err := json.Marshal(req)
		if err != nil {
			err = gerror.NoTraceID(err)
			sugar.Errorw("req",
				zap.String("req", string(rstr)),
				zap.Error(err),
			)
			return
		}
		trace := &Trace{}
		err = json.Unmarshal(rstr, trace)
		if err != nil || trace.TraceID == "" {
			err = gerror.NoTraceID(err)
			sugar.Errorw("req",
				zap.String("req", string(rstr)),
				zap.Error(err),
			)
			return
		}

		start := time.Now()
		resp, err = handler(ctx, req)
		end := time.Now()
		if err != nil {
			sugar.Errorw("req",
				zap.String("trace_id", trace.TraceID),
				zap.String("req", string(rstr)),
				zap.Error(err),
			)
			return
		}
		sugar.Infow("req",
			zap.String("trace_id", trace.TraceID),
			zap.String("consuming", end.Sub(start).String()),
		)
		return
	}
}
