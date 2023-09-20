package model

import "errors"

type Doctor struct {
	Id         int    `json:"id" example:"1" format:"int64"`
	TokenId    string `json:"TokenId" example:"TokenId"`
	Surname    string `json:"Surname" example:"doctor surname"`
	Speciality string `json:"Speciality" example:"doctor speciality"`
	Role       string `json:"Role" example:"doctor role"`
}

var (
	ErrSurnameInvalid = errors.New("name is empty")
)

type AddDoctor struct {
	TokenId    string `json:"TokenId" example:"TokenId"`
	Surname    string `json:"surname" example:"doctor surname"`
	Speciality string `json:"Speciality" example:"doctor speciality"`
	Role       string `json:"Role" example:"doctor role"`
}
