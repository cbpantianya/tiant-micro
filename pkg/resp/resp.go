package resp

import (
	"fmt"
	"net/http"
)

func SuccessResp(data interface{}, traceID string) (code int, resp interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code": "0",
		"msg": "success",
		"data": data,
		"trace_id": traceID,
	}
}

func ErrorResp(ecode int, msg string, traceID string) (code int, resp interface{}) {
	return 200, map[string]string{
		"code": fmt.Sprintf("%d", ecode),
		"msg": msg,
		"data": "",
		"trace_id": traceID,
	}
}
