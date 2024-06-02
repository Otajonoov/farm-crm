package main

import (
	"farm-crm/internal/domain"
	_ "farm-crm/internal/port/http/docs"
	v1 "farm-crm/internal/port/http/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Farm CRM API
// @version 1.0
// @description This is a sample server for a farm CRM.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:50055
// @BasePath /
func main() {
	r := gin.Default()

	var animalUsecase domain.AnimalUseCase
	var factory domain.Factory

	animalHandler := v1.AnimalHandler{
		Usecase: animalUsecase,
		Factory: factory,
	}

	r.POST("/animals", animalHandler.Create)
	r.GET("/animals/:guid", animalHandler.RetrieveById)
	r.GET("/animals", animalHandler.RetrieveList)
	r.PUT("/animals", animalHandler.Update)
	r.DELETE("/animals/:guid", animalHandler.Delete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":50055")
}
