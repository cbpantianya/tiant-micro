package gerror

import (
	"encoding/json"
	"fmt"
	"regexp"

	"go.uber.org/zap"
)

type GError struct {
	Msg     string
	Code    int
	HCode   int
	Inner   error
	HttpMsg string
}

var Errs = map[int]GError{
	42900: ToManyRequest(nil),
	50000: InnerError(nil),
	50001: NoTraceID(nil),
}

func ToManyRequest(err error) GError {
	return GError{
		HCode:   429,
		Msg:     "too many requests",
		Code:    42900,
		Inner:   err,
		HttpMsg: "too many requests",
	}
}

func NoTraceID(err error) GError {
	return GError{
		HCode:   500,
		Msg:     "trace_id not found",
		Code:    50001,
		Inner:   err,
		HttpMsg: "trace_id not found",
	}
}

func InnerError(err error) GError {
	return GError{
		HCode:   500,
		Msg:     "unexpected error",
		Code:    50000,
		Inner:   err,
		HttpMsg: "inner error",
	}
}

func (e GError) Error() string {
	bt, _ := json.Marshal(e)
	return string(bt)
}

func (e GError) RError() error {
	return e.Inner
}

func Unmarshal(msg string, logger *zap.SugaredLogger) *GError {
	e := GError{}
	// use regex to get the value of desc
	fmt.Println(msg)
	// Example: rpc error: code = Unknown desc = {"Msg":"trace_id not found","Code":50001,"Inner":null,"HttpMsg":"inner error"}
	regexp := regexp.MustCompile(`desc = (.*)`)
	msg = regexp.FindStringSubmatch(msg)[1]

	err := json.Unmarshal([]byte(msg), &e)
	//
	if err != nil {
		logger.Errorw("pkg/gerror",
			zap.String("msg", "failed to unmarshal error"),
		)
		e := InnerError(err)
		return &e
	}
	return &e
}
