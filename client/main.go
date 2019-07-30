package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gRPC-hands-on/clientstubs"
	"gRPC-hands-on/user"
)

func createUserHelper(client user.UserClient, userObj *user.UserRequest) {
	resp, err := client.CreateUser(context.Background(), userObj)
	if err != nil {
		log.Println("Failed creating a user!!", err)
	}
	if resp.Success {
		log.Println("Successfully created a user with id: ", resp.Id)
	}
}

func getUsersHelper(client user.UserClient, filter *user.UserFilter) {
	stream, err := client.GetUsers(context.Background(), filter)
	if err != nil {
		log.Println("Failed geting users!!", err)
	}
	counter := 0
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Failed retrieving a user!!", err)
		}
		log.Println("USER: %v", user)
		counter += 1
	}
	log.Printf("Total records retrieved: %d", counter)
}

func main() {
	log.Println("First gRPC implementation")

	conn, err := grpc.Dial("localhost:9998", grpc.WithInsecure())
	if err != nil {
		log.Println("Unable to communicate with server", err)
	}
	defer conn.Close()

	client := user.NewUserClient(conn)

	for i := 0; i < 100; i++ {
		createUserHelper(client, clientstubs.UserOne)
		createUserHelper(client, clientstubs.UserTwo)
	}

	filter := &user.UserFilter{Keyword: ""}
	getUsersHelper(client, filter)
}
