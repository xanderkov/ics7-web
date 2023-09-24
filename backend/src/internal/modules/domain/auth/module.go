package auth

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/auth/repo"
	"hospital/internal/modules/domain/auth/service"
	doctor_repo "hospital/internal/modules/domain/doctor/repo"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,
		fx.Provide(
			fx.Annotate(
				func(r *doctor_repo.DoctorRepo) *doctor_repo.DoctorRepo { return r },
				fx.As(new(service.IDoctorRepo)),
			),
		),
		fx.Provide(
			fx.Annotate(
				func(r *repo.AccountRepo) *repo.AccountRepo { return r },
				fx.As(new(service.IAccountRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
	)
)
