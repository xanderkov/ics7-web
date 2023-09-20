package rest_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"hospital/docs"
	"hospital/internal/modules/config"
	"hospital/internal/modules/view/rest-api/controller"
)

// @securityDefinitions.basic  BasicAuth

func api(controller *controller.Controller, cfg config.Config, logger *zap.Logger) {
	address := fmt.Sprintf("localhost:%d", cfg.ApiPort)
	logger.Info("Server started on address: " + address)
	r := gin.Default()
	c := controller

	v1 := r.Group("/api/v1")
	{
		doctors := v1.Group("/doctors")
		{
			doctors.GET(":id", c.ShowDoctor)
			doctors.POST("", c.SingUp)
		}
		//...
	}

	docs.SwaggerInfo.Title = "Psycho Admin API"
	docs.SwaggerInfo.Description = "This is psycho admin api."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8090"
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(address)
}
