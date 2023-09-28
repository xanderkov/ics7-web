package dto

import "time"

type Treatment struct {
	Id                     int
	Tablets                string
	Survey                 string
	PsychologicalTreatment string
	PatientNumber          int
	UpdateAt               time.Time
}

type Treatments []*Treatment

type CreateTreatment struct {
	Tablets                string
	Survey                 string
	PsychologicalTreatment string
	PatientNumber          int
	UpdateAt               time.Time
}

type UpdateTreatment struct {
	Tablets                string
	Survey                 string
	PsychologicalTreatment string
	PatientNumber          int
	UpdateAt               time.Time
}
