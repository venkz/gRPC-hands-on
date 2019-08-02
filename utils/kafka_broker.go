package broker

import (
	"time"

	"github.com/segmentio/kafka-go"
)

var topic = "TopicGRPC"
var address = "localhost:9092"
var KafkaReadTimeout = 5*time.Second

var KafkaWriter = kafka.NewWriter(kafka.WriterConfig{
	Brokers:  []string{address},
	Topic:    topic,
	Balancer: &kafka.LeastBytes{},
})

var KafkaReader = kafka.NewReader(kafka.ReaderConfig{
	Brokers:  []string{address},
	Topic:    topic,
	Partition: 0,
	MinBytes:  10e3,
	MaxBytes:  10e6,
})