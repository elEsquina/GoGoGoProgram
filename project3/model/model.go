package model 

type Address struct {
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	PostalCode string `json:"postal_code"`
}

type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Credit int `json:"credit"`
}

type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Major string `json:"major"`
	Address Address `json:"address"`
	Courses []Course `json:"courses"`
}