package treatment

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/treatment/repo"
	"hospital/internal/modules/domain/treatment/service"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.TreatmentRepo) *repo.TreatmentRepo { return r },
				fx.As(new(service.ITreatmentRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
