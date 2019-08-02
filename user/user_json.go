package user

type Team struct {
	TeamId   int32  `json:"team_id"`
	TeamName string `json:"team_name"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type User struct {
	Id          int32     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Designation string    `json:"designation"`
	Team        Team      `json:"team"`
	Addresses   []Address `json:"addresses"`
}
