package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ErrorLevel(l zapcore.Level) bool  {
	return l == zap.ErrorLevel
}

func InfoLevel(l zapcore.Level) bool {
	return l == zap.InfoLevel
}

func DebugLevel(l zapcore.Level) bool {
	return l == zap.DebugLevel
}