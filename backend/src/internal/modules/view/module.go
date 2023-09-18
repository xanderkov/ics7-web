package telegram

import (
	"go.uber.org/fx"
	"hospital/internal/modules/view/rest-api"
	"hospital/internal/modules/view/telegram"
)

var (
	Module     = fx.Provide(rest_api.Module, telegram.Module)
	Invokables = fx.Invoke(rest_api.Invokables, telegram.Invokables)
)
