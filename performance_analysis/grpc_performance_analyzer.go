package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"gRPC-hands-on/client"
	"gRPC-hands-on/clientstubs"
	"gRPC-hands-on/user"
	broker "gRPC-hands-on/utils"
)

func createUsersGRPC(userClient user.UserClient,  n int)  {
	start := time.Now()
	i := 1
	for i <= n {
		client.CreateUserHelper(userClient, clientstubs.UserOne)
		if i % 25 == 0 {
			glog.Infof(".. %d users created", i)
		}
		i++
	}
	end := time.Since(start)
	glog.Infof("Time taken to create 100 records via gRPC<>Kafka: %s", end)
	glog.Infof("Created %d users", i)
}

func getUsersGRPC(userClient user.UserClient)  {
	start := time.Now()
	filter := &user.UserFilter{Keyword: ""}
	client.GetUsersHelper(userClient, filter)
	// Remove the KafkaReadTimeout since the ReadMessage function would have waited for that long before breaking conn.
	end := time.Since(start) - broker.KafkaReadTimeout
	glog.Infof("Time taken to read 100 records via gRPC<>Kafka: %s", end)
}

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	conn, err := grpc.Dial("localhost:9998", grpc.WithInsecure())
	if err != nil {
		glog.Infoln("Unable to communicate with server", err)
	}
	defer conn.Close()
	client := user.NewUserClient(conn)

	n := 100
	glog.Infoln("---------- Starting gRPC Kafka performance test ----------")
	glog.Infof(" >>>>> Test to create %d users <<<<<", n)
	createUsersGRPC(client, n)
	glog.Infof(" >>>>> Test to get %d users <<<<<", n)
	getUsersGRPC(client)
}