package strmessages

import (
	"context"

	"github.com/golang/glog"
	"github.com/segmentio/kafka-go"
)

func Producer(kafkaWriter *kafka.Writer) {
	msg1 := kafka.Message{Value: []byte("This is first message")}
	msg2 := kafka.Message{Value: []byte("This is second message")}
	msg3 := kafka.Message{Value: []byte("This is third message")}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := kafkaWriter.WriteMessages(context.Background(), msg1, msg2, msg3)
	if err != nil {
		glog.Errorln(err)
		return
	}
	glog.Infof("Wrote messages to kafka")
	kafkaWriter.Close()
}
