package model

type Room struct {
	Id             int    `json:"id" example:"1" format:"int64"`
	Num            int    `json:"Num" example:"1"`
	Floor          int    `json:"Floor" example:"1"`
	NumberBeds     int    `json:"NumberBeds" example:"1"`
	TypeRoom       string `json:"TypeRoom" example:"1"`
	NumberPatients int    `json:"NumberPatients" example:"1"`
}

type AddRoom struct {
	Id             int    `json:"id" example:"1" format:"int64"`
	Num            int    `json:"Num" example:"1"`
	Floor          int    `json:"Floor" example:"1"`
	NumberBeds     int    `json:"NumberBeds" example:"1"`
	TypeRoom       string `json:"TypeRoom" example:"1"`
	NumberPatients int    `json:"NumberPatients" example:"1"`
}

type UpdateRoom struct {
	Id             int    `json:"id" example:"1" format:"int64"`
	Num            int    `json:"Num" example:"1"`
	Floor          int    `json:"Floor" example:"1"`
	NumberBeds     int    `json:"NumberBeds" example:"1"`
	TypeRoom       string `json:"TypeRoom" example:"1"`
	NumberPatients int    `json:"NumberPatients" example:"1"`
}
