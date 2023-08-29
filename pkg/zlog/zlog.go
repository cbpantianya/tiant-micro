package zlog

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewEncoder() (encoder zapcore.Encoder) {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder = zapcore.NewJSONEncoder(cfg)
	return
}

func NewDevZLogger() (sugar *zap.SugaredLogger, err error) {
	encoder := NewEncoder()

	std := NewStdWriteSyncer()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.Lock(std), zap.LevelEnablerFunc(DebugLevel)),
		zapcore.NewCore(encoder, zapcore.Lock(std), zap.LevelEnablerFunc(InfoLevel)),
		zapcore.NewCore(encoder, zapcore.Lock(std), zap.LevelEnablerFunc(ErrorLevel)),
	)

	sugar = zap.New(core, zap.AddCaller()).Sugar()
	defer sugar.Sync()
	sugar.Infow("ZLog initialization successful",
		zap.String("text", "zlog"),
	)

	return
}

func NewProdZLogger(addr string, port int, topic string, partition int) (sugar *zap.SugaredLogger, err error) {
	encoder := NewEncoder()

	conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%d", addr, port), topic, partition)
	if err != nil {
		return
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(NewInfoKafkaWriter(conn)), zap.LevelEnablerFunc(InfoLevel)),
		zapcore.NewCore(encoder, zapcore.AddSync(NewErrorKafkaWriter(conn)), zap.ErrorLevel),
	)
	
	sugar = zap.New(core, zap.AddCaller()).Sugar()
	defer sugar.Sync()
	sugar.Infow("ZLog initialization successful",
		zap.String("text", "zlog"),
	)

	return
}
