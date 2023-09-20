package model

type Patient struct {
	Id             int     `json:"id" example:"1" format:"int64"`
	Surname        string  `json:"Surname" example:"patient surname"`
	Name           string  `json:"Name" example:"patient Name"`
	Patronymic     string  `json:"Patronymic" example:"patient patronymic"`
	Height         int     `json:"Height" example:"180"`
	Weight         float64 `json:"Weight" example:"42"`
	RoomNumber     int     `json:"RoomNumber" example:"6"`
	DegreeOfDanger int     `json:"DegreeOfDanger" example:"1"`
}

type AddPatient struct {
	Surname        string  `json:"Surname" example:"patient surname"`
	Name           string  `json:"Name" example:"patient Name"`
	Patronymic     string  `json:"Patronymic" example:"patient patronymic"`
	Height         int     `json:"Height" example:"180"`
	Weight         float64 `json:"Weight" example:"42"`
	RoomNumber     int     `json:"RoomNumber" example:"6"`
	DegreeOfDanger int     `json:"DegreeOfDanger" example:"1"`
}

type UpdatePatient struct {
	Surname        string  `json:"Surname" example:"patient surname"`
	Name           string  `json:"Name" example:"patient Name"`
	Patronymic     string  `json:"Patronymic" example:"patient patronymic"`
	Height         int     `json:"Height" example:"180"`
	Weight         float64 `json:"Weight" example:"42"`
	RoomNumber     int     `json:"RoomNumber" example:"6"`
	DegreeOfDanger int     `json:"DegreeOfDanger" example:"1"`
}
