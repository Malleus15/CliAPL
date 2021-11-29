package models

//In this file there are support structures for requests
//and the respective json mapping of the keys

type UserReq struct {
	FirstName string `json:"first_name"`
	Name string  `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// User inherit UserReq
type User struct {
	UserReq
	ID int `json:"id"`
	IsActive bool `json:"is_active"`
}

// LoginReq  support struct for login
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ParametersReq  body to save a set of parameters
type ParametersReq struct {
	InvestorsNumber int `json:"investors_number"`
	NumbRtPlayers int `json:"number_rt_players"`
	PriceCpu float32 `json:"price_cpu"`
	//it can be very large
	HostingCapacity int64 `json:"hosting_capacity"`
	DurationCpu int `json:"duration_cpu"`
	UserId int `json:"user_id"`
}

// Parameters body to save a set of parameters
type Parameters struct {
	ParametersReq
	Id int `json:"id"`
}

// InvestmentReq body to request an Investment
type InvestmentReq struct {
	Fairness     bool `json:"fairness"`
	ParametersId int  `json:"parameters_id"`
}
