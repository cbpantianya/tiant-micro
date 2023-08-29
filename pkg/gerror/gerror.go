package gerror

import (
	"encoding/json"
)

type GError struct {
	Msg  string
	Code int
	Inner error
}

func (e GError) Error() string {
	bt, _ := json.Marshal(e)
	return string(bt)
}

func (e GError) RError() error {
	return e.Inner
}

func NoTraceID(err error) GError {
	return GError{
		Msg: "trace_id not found",
		Code: 40412,
		Inner: err,
	}
}
