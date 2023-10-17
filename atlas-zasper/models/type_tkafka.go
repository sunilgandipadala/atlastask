package models

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type TKafkaWriter struct {
	Writer *kafka.Writer
	Tracer *Tracing
}

type TConsumerFunc func(kafka.Message, context.Context) error

type TKafkaConsumer struct {
	Reader      *kafka.Reader
	Ctx         context.Context
	ServiceName string
}
