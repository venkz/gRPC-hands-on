package main

import (
	"flag"
	"time"

	"github.com/golang/glog"

	"gRPC-hands-on/client"
)

func createUsersHttp(n int)  {
	start := time.Now()
	i := 1
	for i <= n {
		client.CreateUser()
		if i % 25 == 0 {
			glog.Infof(".. %d users created", i)
		}
		i++
	}
	end := time.Since(start)
	glog.Infof("Time taken to create 100 records via HTTP<>Kafka: %s", end)
	glog.Infof("Created %d users", i)
}

func getUsersHttp()  {
	start := time.Now()
	client.GetUsers()
	end := time.Since(start)
	glog.Infof("Time taken to read 100 records via HTTP<>Kafka: %s", end)
}

func main()  {
	flag.Parse()
	flag.Set("logtostderr", "true")

	n := 100
	glog.Infoln("---------- Starting HTTP Kafka performance test ----------")
	glog.Infof(" >>>>> Test to create %d users <<<<<", n)
	createUsersHttp(n)
	glog.Infof(" >>>>> Test to get %d users <<<<<", n)
	getUsersHttp()
}