package service

import (
	"context"
	"hospital/internal/modules/domain/treatment/dto"
)

type ITreatmentRepo interface {
	GetById(ctx context.Context, id int) (*dto.Treatment, error)
	List(ctx context.Context) (dto.Treatments, error)
	Create(ctx context.Context, dtm *dto.CreateTreatment) (*dto.Treatment, error)
	Update(ctx context.Context, num int, dtm *dto.UpdateTreatment) (*dto.Treatment, error)
	Delete(ctx context.Context, num int) error
}

type TreatmentService struct {
	repo ITreatmentRepo
}

func NewTreatmentService(repo ITreatmentRepo) *TreatmentService {
	return &TreatmentService{
		repo: repo,
	}
}

func (r *TreatmentService) GetById(ctx context.Context, id int) (*dto.Treatment, error) {
	return r.repo.GetById(ctx, id)
}

func (r *TreatmentService) List(ctx context.Context) (dto.Treatments, error) {
	return r.repo.List(ctx)
}

func (r *TreatmentService) Create(ctx context.Context, dtm *dto.CreateTreatment) (*dto.Treatment, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *TreatmentService) Update(ctx context.Context, id int, dtm *dto.UpdateTreatment) (*dto.Treatment, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *TreatmentService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
