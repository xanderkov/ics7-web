package service

import (
	"context"
	d_dto "hospital/internal/modules/domain/disease/dto"
	"hospital/internal/modules/domain/patient/dto"
	treatment_dto "hospital/internal/modules/domain/treatment/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IPatientRepo

type IPatientRepo interface {
	GetById(ctx context.Context, id int) (*dto.Patient, error)
	List(ctx context.Context) (dto.Patients, error)
	Create(ctx context.Context, dtm *dto.CreatePatient) (*dto.Patient, error)
	Update(ctx context.Context, num int, dtm *dto.UpdatePatient) (*dto.Patient, error)
	Delete(ctx context.Context, num int) error
	GetDiseaseById(ctx context.Context, id int) (*d_dto.Disease, error)
	AddDiseaseById(ctx context.Context, idDoc int, idP int) (*d_dto.Disease, error)
	GetThreatById(ctx context.Context, id int) (treatment_dto.Treatments, error)
}

type PatientService struct {
	repo IPatientRepo
}

func NewPatientService(repo IPatientRepo) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

func (r *PatientService) GetById(ctx context.Context, id int) (*dto.Patient, error) {
	return r.repo.GetById(ctx, id)
}

func (r *PatientService) GetDiseaseById(ctx context.Context, id int) (*d_dto.Disease, error) {
	return r.repo.GetDiseaseById(ctx, id)
}

func (r *PatientService) GetThreatById(ctx context.Context, id int) (treatment_dto.Treatments, error) {
	return r.repo.GetThreatById(ctx, id)
}

func (r *PatientService) AddDiseaseById(ctx context.Context, idP int, idD int) (*d_dto.Disease, error) {
	return r.repo.AddDiseaseById(ctx, idP, idD)
}

func (r *PatientService) List(ctx context.Context) (dto.Patients, error) {
	return r.repo.List(ctx)
}

func (r *PatientService) Create(ctx context.Context, dtm *dto.CreatePatient) (*dto.Patient, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *PatientService) Update(ctx context.Context, id int, dtm *dto.UpdatePatient) (*dto.Patient, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *PatientService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
