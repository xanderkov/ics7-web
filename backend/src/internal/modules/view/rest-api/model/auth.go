package model

type Account struct {
	Id           int    `json:"id" example:"1" format:"int64"`
	Login        string `json:"Login" example:"Login"`
	PasswordHash string `json:"PasswordHash" example:"PasswordHash"`
}

type AddAccount struct {
	Login        string `json:"Login" example:"Login"`
	PasswordHash string `json:"PasswordHash" example:"PasswordHash"`
}

type UpdateAccount struct {
	Login        string `json:"Login" example:"Login"`
	PasswordHash string `json:"PasswordHash" example:"PasswordHash"`
}
