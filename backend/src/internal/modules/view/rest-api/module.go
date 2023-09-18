package rest_api

import (
	"go.uber.org/fx"
	"hospital/internal/modules/view/rest-api/controller"
)

var (
	Module     = fx.Provide(controller.NewController)
	Invokables = fx.Invoke(api)
)
