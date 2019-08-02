package clientstubs

import (
	"gRPC-hands-on/user"
)

var UserHttpObj = &user.User{
	Id:          1,
	Name:        "venky",
	Email:       "vk@mail.com",
	Designation: "staff se",
	Team: user.Team{
		TeamId:   100,
		TeamName: "SCA",
	},
	Addresses: []user.Address{
		user.Address{
			City:    "Milapeitas",
			State:   "OK",
			Country: "USA",
		},
		user.Address{
			City:    "Blazewana",
			State:   "UP",
			Country: "India",
		},
	},
}
