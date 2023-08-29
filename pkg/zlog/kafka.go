package zlog

import (
	"github.com/segmentio/kafka-go"
)

type ErrorKafkaWriter struct {
	Conn *kafka.Conn
}

func (k *ErrorKafkaWriter) Write(p []byte) (n int, err error) {
	_, err = k.Conn.WriteMessages(
		kafka.Message{
			Key:   []byte("error"),
			Value: p,
		},
	)
	n = len(p)
	return
}

func NewErrorKafkaWriter(conn *kafka.Conn) *ErrorKafkaWriter {
	return &ErrorKafkaWriter{Conn: conn}
}

type InfoKafkaWriter struct {
	Conn *kafka.Conn
}

func (k InfoKafkaWriter) Write(p []byte) (n int, err error) {
	_, err = k.Conn.WriteMessages(
		kafka.Message{
			Key:   []byte("info"),
			Value: p,
		},
	)
	n = len(p)
	return
}

func NewInfoKafkaWriter(conn *kafka.Conn) *InfoKafkaWriter {
	return &InfoKafkaWriter{Conn: conn}
}
