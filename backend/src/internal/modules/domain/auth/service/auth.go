package service

import (
	"context"
	"hospital/internal/modules/config"
	"hospital/internal/modules/domain/auth/dto"
	doctor_dto "hospital/internal/modules/domain/doctor/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IDoctorRepo

type IDoctorRepo interface {
	// GetByTokenId(ctx context.Context, tokenId string) (*doctor_dto.Doctor, error)
	Create(ctx context.Context, dtm *doctor_dto.CreateDoctor) (*doctor_dto.Doctor, error)
}

type AuthService struct {
	repo    IDoctorRepo
	tokenId string
}

func NewAuthService(repo IDoctorRepo, config config.Config) *AuthService {
	return &AuthService{
		repo:    repo,
		tokenId: config.Secret,
	}
}

func (r *AuthService) SignUp(ctx context.Context, newDoctor *dto.NewDoctor) (*doctor_dto.Doctor, error) {

	createDoctor := &doctor_dto.CreateDoctor{
		Surname:    newDoctor.Surname,
		TokenId:    newDoctor.TokenId,
		Speciality: newDoctor.Speciality,
		Role:       newDoctor.Role,
	}

	// Создаем пользователя
	createdDoctor, err := r.repo.Create(ctx, createDoctor)
	if err != nil {
		return nil, err
	}

	return createdDoctor, nil
}

type IAccountRepo interface {
	GetById(ctx context.Context, id int) (*dto.Account, error)
	List(ctx context.Context) (dto.Accounts, error)
	Create(ctx context.Context, dtm *dto.CreateAccount) (*dto.Account, error)
	Update(ctx context.Context, id int, dtm *dto.UpdateAccount) (*dto.Account, error)
	Delete(ctx context.Context, id int) error
	Login(ctx context.Context, dtm *dto.CreateAccount) (*dto.Account, error)
}

type AccountService struct {
	repo IAccountRepo
}

func NewAccountService(repo IAccountRepo) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (r *AccountService) GetById(ctx context.Context, id int) (*dto.Account, error) {
	return r.repo.GetById(ctx, id)
}

func (r *AccountService) List(ctx context.Context) (dto.Accounts, error) {
	return r.repo.List(ctx)
}

func (r *AccountService) Create(ctx context.Context, dtm *dto.CreateAccount) (*dto.Account, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *AccountService) Update(ctx context.Context, id int, dtm *dto.UpdateAccount) (*dto.Account, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *AccountService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func (r *AccountService) Login(ctx context.Context, dtm *dto.CreateAccount) (*dto.Account, error) {
	return r.repo.Login(ctx, dtm)
}
