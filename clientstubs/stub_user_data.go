package clientstubs

import (
	"gRPC-hands-on/user"
)

var UserOne = &user.UserRequest{
	Id:          1,
	Name:        "Venky",
	Email:       "vk@mail.com",
	Designation: "staff se",

	Team: &user.UserRequest_Team{
		TeamId:   100,
		TeamName: "SCA",
	},

	Addresses: []*user.UserRequest_Address{
		&user.UserRequest_Address{
			City:    "Milapeitas",
			State:   "OK",
			Country: "USA",
		},
		&user.UserRequest_Address{
			City:    "Blazewana",
			State:   "UP",
			Country: "India",
		},
	},
}
