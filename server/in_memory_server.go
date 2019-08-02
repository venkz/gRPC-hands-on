package main

import (
	"strings"

	"github.com/golang/glog"
	"golang.org/x/net/context"

	"gRPC-hands-on/user"
)

type InMemoryStorage struct {
	savedUsers []*user.UserRequest
}

func (s *InMemoryStorage) CreateUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	s.savedUsers = append(s.savedUsers, in)
	glog.Infof("New user created. Total users in system: %d", len(s.savedUsers))
	return &user.UserResponse{Id: in.Id, Success: true}, nil
}

func (s *InMemoryStorage) GetUsers(in *user.UserFilter, stream user.User_GetUsersServer) error {
	for _, user := range s.savedUsers {
		if in.Keyword != "" {
			if !strings.Contains(user.Name, in.Keyword) {
				continue
			}
		}
		if err := stream.Send(user); err != nil {
			glog.Infoln(err)
			return err
		}
	}
	return nil
}


