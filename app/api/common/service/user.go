package service

import (
	"context"
	"tiant-micro/app/api/common/server"
	"tiant-micro/app/rpc/user/pb"
	"tiant-micro/pkg/gerror"
	"tiant-micro/pkg/resp"

	"github.com/gin-gonic/gin"
)

func UserService(s *server.Server) gin.HandlerFunc {
	return func(gtx *gin.Context) {
		trace := gtx.GetString("trace_id")
		res, err := s.Svc.User.UserProfile(context.Background(), &pb.UserProfileRequest{
			UserId:  "1",
			TraceId: trace,
		})

		if err != nil {
			// err to string
			ge := gerror.Unmarshal(err.Error(), s.Svc.Logger)
			gtx.JSON(resp.ErrorResp(ge.Code,ge.HttpMsg, trace))
			return
		}

		gtx.JSON(resp.SuccessResp(res,trace))

	}

}
