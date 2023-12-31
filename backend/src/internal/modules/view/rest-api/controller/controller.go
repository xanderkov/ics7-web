package controller

import (
	auth_serv "hospital/internal/modules/domain/auth/service"
	disease_servis "hospital/internal/modules/domain/disease/service"
	doctor_server "hospital/internal/modules/domain/doctor/service"
	patient_servis "hospital/internal/modules/domain/patient/service"
	room_servis "hospital/internal/modules/domain/room/service"
	treatment_servis "hospital/internal/modules/domain/treatment/service"
)

type Controller struct {
	doctorService    *doctor_server.DoctorService
	authService      *auth_serv.AuthService
	patientService   *patient_servis.PatientService
	roomService      *room_servis.RoomService
	diseaseService   *disease_servis.DiseaseService
	treatmentService *treatment_servis.TreatmentService
	accountService   *auth_serv.AccountService
}

func NewController(
	doctorService *doctor_server.DoctorService,
	authService *auth_serv.AuthService,
	accountService *auth_serv.AccountService,
	patientService *patient_servis.PatientService,
	roomService *room_servis.RoomService,
	diseaseService *disease_servis.DiseaseService,
	treatmentService *treatment_servis.TreatmentService,

) *Controller {

	r := &Controller{
		authService:      authService,
		doctorService:    doctorService,
		patientService:   patientService,
		roomService:      roomService,
		diseaseService:   diseaseService,
		treatmentService: treatmentService,
		accountService:   accountService,
	}

	return r
}
