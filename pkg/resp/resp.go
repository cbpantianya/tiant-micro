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

func ErrorResp(hcode int,ecode int, msg string, traceID string) (code int, resp interface{}) {
	if hcode == 0 {
		hcode = http.StatusOK
	}
	return hcode, map[string]string{
		"code": fmt.Sprintf("%d", ecode),
		"msg": msg,
		"data": "",
		"trace_id": traceID,
	}
}
