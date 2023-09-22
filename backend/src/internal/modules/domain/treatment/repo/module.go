package repo

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewTreatmentRepo)
	Invokables = fx.Invoke()
)
