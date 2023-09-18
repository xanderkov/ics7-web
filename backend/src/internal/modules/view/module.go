package view

import (
	"go.uber.org/fx"
	"hospital/internal/modules/view/controllers"
	rest_api "hospital/internal/modules/view/rest-api"
	"hospital/internal/modules/view/telegram"
)

var (
	Module     = fx.Provide(controllers.NewController)
	Invokables = fx.Invoke(rest_api.Api, telegram.StartBot)
)
