package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"

	"gRPC-hands-on/user"
	broker "gRPC-hands-on/utils"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	byteArray, _ := ioutil.ReadAll(r.Body)
	msg := kafka.Message{Value: byteArray}
	err := broker.KafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		glog.Errorln("Failed to create/put a new user")
	}
	userObj := user.User{}
	json.Unmarshal(byteArray, &userObj)
	json.NewEncoder(w).Encode(&userObj)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []user.User{}
	for {
		// Specify a context with expiration to unblock the async call
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Read for all the messages after previously committed offset.
		// NOTE: This is a blocking call until the context expires
		message, err := broker.KafkaReader.ReadMessage(ctx)
		if err != nil {
			break
		}

		// Unmarshal byte data as pb object
		user := user.User{}
		json.Unmarshal(message.Value, &user)

		users = append(users, user)
		// glog.Infof("message at offset %d: %s retrieved", message.Offset, string(message.Key))
	}

	json.NewEncoder(w).Encode(&users)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")

	http.ListenAndServe(":8088", router)
}