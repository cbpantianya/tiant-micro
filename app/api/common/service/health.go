package service

import (
	"tiant-micro/app/api/common/model"
	"tiant-micro/app/api/common/server"
	"tiant-micro/pkg/resp"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthService(s *server.Server) gin.HandlerFunc {
	time.Sleep(1 * time.Second)
	return func(gtx *gin.Context) {
		trace := gtx.GetString("trace_id")
		gtx.JSON(resp.SuccessResp(model.HealthResp{
			Msg:  "ok",
			Time: time.Now().Unix(),
		}, trace))
	}

}
