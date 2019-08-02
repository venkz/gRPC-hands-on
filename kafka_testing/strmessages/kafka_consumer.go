package strmessages

import (
	"context"

	"github.com/golang/glog"
	"github.com/segmentio/kafka-go"
)

func Consumer(kafkaReader *kafka.Reader) {
	glog.Infof("Starting to look for messages")

	for {
		glog.Infof("Starting to read from topic")

		message, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		glog.Infof("message at offset %d: %s = %s", message.Offset, string(message.Key), string(message.Value))
	}

	kafkaReader.Close()
}
