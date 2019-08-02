package main

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"

	"gRPC-hands-on/user"
	"gRPC-hands-on/utils"
)

type KafkaServer struct {}


func (server *KafkaServer) CreateUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	byteArray, _ := proto.Marshal(in)
	msg := kafka.Message{Value: byteArray}
	err := broker.KafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{Id: in.Id, Success: true}, nil
}

func (server *KafkaServer) GetUsers(in *user.UserFilter, stream user.User_GetUsersServer) error {
	for {
		// Specify a context with expiration to unblock the async call
		ctx, cancel := context.WithTimeout(context.Background(), broker.KafkaReadTimeout)
		defer cancel()

		// Read for all the messages after previously committed offset.
		// NOTE: This is a blocking call until the context expires
		message, err := broker.KafkaReader.ReadMessage(ctx)
		if err != nil {
			glog.Errorln(err)
			break
		}

		// Unmarshal byte data as pb object
		out := &user.UserRequest{}
		proto.Unmarshal(message.Value, out)

		// Send it to the gRPC stream
		if err := stream.Send(out); err != nil {
			glog.Infoln(err)
			return err
		}
		// glog.Infof("message at offset %d: %s retrieved", message.Offset, string(message.Key))
	}
	broker.KafkaReader.Close()
	return nil
}
