package model

import "time"

type Treatment struct {
	Id                     int       `json:"id" example:"1" format:"int64"`
	Tablets                string    `json:"Tablets" example:"Abrobius"`
	PsychologicalTreatment string    `json:"PsychologicalTreatment" example:"Заболел"`
	Survey                 string    `json:"Survey" example:"1"`
	PatientNumber          int       `json:"PatientNumber" example:"1"`
	UpdatedAt              time.Time `json:"UpdatedAt" example:"1"`
}

type AddTreatment struct {
	Tablets                string    `json:"Tablets" example:"Abrobius"`
	PsychologicalTreatment string    `json:"PsychologicalTreatment" example:"Заболел"`
	Survey                 string    `json:"Survey" example:"1"`
	PatientNumber          int       `json:"PatientNumber" example:"1"`
	UpdatedAt              time.Time `json:"UpdatedAt" example:"1"`
}

type UpdateTreatment struct {
	Tablets                string    `json:"Tablets" example:"Abrobius"`
	PsychologicalTreatment string    `json:"PsychologicalTreatment" example:"Заболел"`
	Survey                 string    `json:"Survey" example:"1"`
	PatientNumber          int       `json:"PatientNumber" example:"1"`
	UpdatedAt              time.Time `json:"UpdatedAt" example:"1"`
}
