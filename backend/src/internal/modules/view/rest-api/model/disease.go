package model

type Disease struct {
	Id             int    `json:"id" example:"1" format:"int64"`
	Name           string `json:"Name" example:"Abrobius"`
	Threat         string `json:"Threat" example:"Заболел"`
	DegreeOfDanger int    `json:"DegreeOfDanger" example:"1"`
}

type AddDisease struct {
	Name           string `json:"Name" example:"Abrobius"`
	Threat         string `json:"Threat" example:"Заболел"`
	DegreeOfDanger int    `json:"DegreeOfDanger" example:"1"`
}

type UpdateDisease struct {
	Name           string `json:"Name" example:"Abrobius"`
	Threat         string `json:"Threat" example:"Заболел"`
	DegreeOfDanger int    `json:"DegreeOfDanger" example:"1"`
}
