package dto

type Treatment struct {
	Id                     int
	Tablets                string
	Survey                 string
	PsychologicalTreatment string
}

type Treatments []*Treatment

type CreateTreatment struct {
	Tablets                string
	Survey                 string
	PsychologicalTreatment string
}

type UpdateTreatment struct {
	Tablets                string
	Survey                 string
	PsychologicalTreatment string
}
