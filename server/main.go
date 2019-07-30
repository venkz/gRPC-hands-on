package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gRPC-hands-on/user"
)

type server struct {
	savedUsers []*user.UserRequest
}

func (s *server) CreateUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	s.savedUsers = append(s.savedUsers, in)
	log.Printf("New user created. Total users in system: %d", len(s.savedUsers))
	return &user.UserResponse{Id: in.Id, Success: true}, nil
}

func (s *server) GetUsers(in *user.UserFilter, stream user.User_GetUsersServer) error {
	for _, user := range s.savedUsers {
		if in.Keyword != "" {
			if !strings.Contains(user.Name, in.Keyword) {
				continue
			}
		}
		if err := stream.Send(user); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":9998")
	if err != nil {
		log.Printf("Unable to start server")
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}
