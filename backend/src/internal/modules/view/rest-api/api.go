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
	address := fmt.Sprintf("0.0.0.0:%d", cfg.ApiPort)
	logger.Info("Server started on address: " + address)
	r := gin.Default()
	c := controller

	v1 := r.Group("/api/v1")
	{
		doctors := v1.Group("/doctors")
		{
			doctors.GET(":id", c.ShowDoctor)
			doctors.POST("", c.AddDoctor)
			doctors.GET("", c.ListDoctors)
			doctors.PATCH(":id", c.UpdateDoctor)
			doctors.DELETE(":id", c.DeleteDoctor)
		}
		patients := v1.Group("/patients")
		{
			patients.GET(":id", c.ShowPatient)
			patients.GET("", c.ListPatients)
			patients.POST("", c.AddPatient)
			patients.PATCH(":id", c.UpdatePatient)
			patients.DELETE(":id", c.DeletePatient)
		}

		room := v1.Group("/rooms")
		{
			room.GET(":id", c.ShowRoom)
			room.GET("", c.ListRooms)
			room.POST("", c.AddRoom)
			room.PATCH(":id", c.UpdateRoom)
			room.DELETE(":id", c.DeleteRoom)
		}

		disease := v1.Group("/diseases")
		{
			disease.GET(":id", c.ShowDisease)
			disease.GET("", c.ListDisease)
			disease.POST("", c.AddDisease)
			disease.PATCH(":id", c.UpdateDisease)
			disease.DELETE(":id", c.DeleteDisease)
		}

		treatment := v1.Group("/treatments")
		{
			treatment.GET(":id", c.ShowTreatment)
			treatment.GET("", c.ListTreatments)
			treatment.POST("", c.AddTreatment)
			treatment.PATCH(":id", c.UpdateTreatment)
			treatment.DELETE(":id", c.DeleteTreatment)
		}

		account := v1.Group("/accounts")
		{
			account.POST("/old", c.AddDoctor)

			account.GET(":id", c.ShowAccount)
			account.GET("", c.ListAccount)
			account.POST("", c.AddAccount)
			account.PATCH(":id", c.UpdateAccount)
			account.DELETE(":id", c.DeleteAccount)
		}
	}

	docs.SwaggerInfo.Title = "Psycho Admin API"
	docs.SwaggerInfo.Description = "This is psycho admin api."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8090"
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(address)
}
