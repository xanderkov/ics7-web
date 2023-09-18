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

// @title           Psycho Admin API
// @version         1.0
// @description     This is psycho admin api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8090
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func api(controller *controller.Controller, cfg config.Config, logger *zap.Logger) {
	address := fmt.Sprintf("localhost:%d", cfg.ApiPort)
	logger.Info("Server started on address: " + address)
	r := gin.Default()
	c := controller

	v1 := r.Group("/api/v1")
	{
		doctors := v1.Group("/doctors")
		{
			doctors.GET(":id", c.ShowAccount)
		}
		//...
	}

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(address)
}
