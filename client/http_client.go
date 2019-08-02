package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"gRPC-hands-on/clientstubs"
	"gRPC-hands-on/user"
)

func GetUsers() {
	response, err := http.Get("http://localhost:8088/users")
	if err != nil {
		glog.Errorln("Error getting users!")
		return
	}
	users := []user.User{}
	json.NewDecoder(response.Body).Decode(&users)
	glog.Infof("Retrieved %d records from server", len(users))
}

func CreateUser()  {
	data, _ := json.Marshal(clientstubs.UserHttpObj)
	response, err := http.Post("http://localhost:8088/user", "application/json", bytes.NewBuffer(data))
	if err != nil {
		glog.Errorln("Error creating user!")
		return
	}
	userObj := &user.User{}
	json.NewDecoder(response.Body).Decode(userObj)
	glog.Infof("Successfully created a user with id: %d ", userObj.Id)
}

// func main() {
// 	createUser()
// 	getUsers()
// }