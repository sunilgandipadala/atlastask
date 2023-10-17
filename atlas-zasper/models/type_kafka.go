package models

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	KafkaWriter *kafka.Writer `json:"-"`
}

type KafkaReader struct {
	KafkaReader *kafka.Reader `json:"-"`
}
