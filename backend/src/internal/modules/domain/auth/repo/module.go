package repo

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewAccountRepo)
	Invokables = fx.Invoke()
)
