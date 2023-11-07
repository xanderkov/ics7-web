package model

type DoctorsPatients struct {
	DoctorId  int `json:"doctorId" example:"1" format:"int64"`
	PatientId int `json:"patientId" example:"1" format:"int64"`
}

type AddDoctorsPatients struct {
	DoctorId  int `json:"doctorId" example:"1" format:"int64"`
	PatientId int `json:"patientId" example:"1" format:"int64"`
}
