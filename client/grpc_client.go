package client

import (
	"io"

	"github.com/golang/glog"
	"golang.org/x/net/context"

	"gRPC-hands-on/user"
)

func CreateUserHelper(client user.UserClient, userObj *user.UserRequest) {
	resp, err := client.CreateUser(context.Background(), userObj)
	if err != nil {
		glog.Infoln("Failed creating a user!!", err)
	}
	if resp.Success {
		glog.Infof("Successfully created a user with id: %d", resp.Id)
	}
}

func GetUsersHelper(client user.UserClient, filter *user.UserFilter) {
	stream, err := client.GetUsers(context.Background(), filter)
	if err != nil {
		glog.Infoln("Failed geting users!!", err)
	}
	counter := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Infoln("Failed retrieving a user!!", err)
		}
		// glog.Infoln("USER: %v", user)
		counter += 1
	}
	glog.Infof("Total records retrieved: %d", counter)
}

// func main() {
// 	operFlag := flag.String("operation", "foo", "...")
//
// 	conn, err := grpc.Dial("localhost:9998", grpc.WithInsecure())
// 	if err != nil {
// 		glog.Infoln("Unable to communicate with server", err)
// 	}
// 	defer conn.Close()
//
// 	client := user.NewUserClient(conn)
//
// 	flag.Parse()
// 	switch *operFlag {
// 	case "get":
// 		{
// 			filter := &user.UserFilter{Keyword: ""}
// 			getUsersHelper(client, filter)
// 		}
// 	case "create":
// 		{
// 			createUserHelper(client, clientstubs.UserOne)
// 		}
// 	case "foo":
// 		{
// 			glog.Errorln("Specify operation using --operation=[get|create]")
// 		}
// 	}
// }
