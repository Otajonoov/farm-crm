package v1

import (
	"farm-crm/internal/domain"
	"farm-crm/internal/port/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnimalTypeHandler struct {
	Usecase domain.AnimalTypeUseCase
	Factory domain.Factory
}

// Create a new type for animal
// @Summary Create a new animal
// @Description Create a new type animal with the provided details
// @Tags animals
// @Accept  json
// @Produce  json
// @Param animal body model.AnimalType true "AnimalType"
// @Success 201 {object} model.AnimalType
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /animalType [post]
func (a AnimalTypeHandler) Create(c *gin.Context) {
	var animalType model.AnimalType

	if err := c.ShouldBindJSON(&animalType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	animalFactory := a.Factory.ParseToDomainForAnimalType(animalType.Guid, animalType.AnimalType, animalType.Count)
	err := a.Usecase.Create(animalFactory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, animalType.Guid)
}
