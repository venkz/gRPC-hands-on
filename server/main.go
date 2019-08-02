package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"gRPC-hands-on/user"
)

func main() {
	serverFlag := flag.String("backend", "foo", "...")

	listener, err := net.Listen("tcp", ":9998")
	if err != nil {
		glog.Infof("Unable to start server")
	}

	grpcServer := grpc.NewServer()
	flag.Parse()
	switch *serverFlag {
	case "grpc_kafka":
		{
			glog.Infoln("Starting gRPC server")
			user.RegisterUserServer(grpcServer, &KafkaServer{})
			grpcServer.Serve(listener)
		}
	case "memory":
		{
			glog.Infoln("Starting in memory server")
			user.RegisterUserServer(grpcServer, &InMemoryStorage{})
			grpcServer.Serve(listener)
		}
	case "http_kafka":
		{
			glog.Infoln("Starting in HTTP server")
			handleRequests()
		}
	case "foo":
		{
			glog.Errorln("Specify backend by passing --backend = [grpc_kafka|memory|http_kafka]")
		}
	}
}
