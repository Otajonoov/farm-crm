package v1

import (
	"farm-crm/internal/domain"
	"farm-crm/internal/port/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnimalHandler struct {
	Usecase domain.AnimalUseCase
	Factory domain.Factory
}

// Create a new animal
// @Summary Create a new animal
// @Description Create a new animal with the provided details
// @Tags animals
// @Accept  json
// @Produce  json
// @Param animal body model.Animal true "Animal"
// @Success 201 {object} model.Animal
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /animals [post]
func (a AnimalHandler) Create(c *gin.Context) {
	var animal model.Animal

	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	animalFactory := a.Factory.ParseToDomainFor(animal.AnimalTypeGuid, animal.Descriptive, animal.Weight, animal.IsIll, animal.IsFed)
	err := a.Usecase.Create(animalFactory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, animal.Guid)
}

// RetrieveById retrieves an animal by its GUID
// @Summary Retrieve an animal by its GUID
// @Description Get details of an animal by its GUID
// @Tags animals
// @Produce  json
// @Param guid path string true "GUID"
// @Success 200 {object} model.Animal
// @Failure 500 {object} map[string]string "error"
// @Router /animals/{guid} [get]
func (a AnimalHandler) RetrieveById(c *gin.Context) {
	guid := c.Param("guid")
	zikr, err := a.Usecase.GetByGuid(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zikr)
}

// RetrieveList retrieves a list of all animals
// @Summary Retrieve a list of all animals
// @Description Get a list of all animals
// @Tags animals
// @Produce  json
// @Success 200 {array} model.Animal
// @Failure 500 {object} map[string]string "error"
// @Router /animals [get]
func (a AnimalHandler) RetrieveList(c *gin.Context) {
	zikrs, err := a.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zikrs)
}

// Update an animal
// @Summary Update an existing animal
// @Description Update an animal with the provided details
// @Tags animals
// @Accept  json
// @Produce  json
// @Param animal body model.Animal true "Animal"
// @Success 200 {string} string "updated"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /animals [put]
func (a AnimalHandler) Update(c *gin.Context) {
	var animal model.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	factory := a.Factory.ParseToDomainFor(animal.AnimalTypeGuid, animal.Descriptive, animal.Weight, animal.IsIll, animal.IsFed)
	err := a.Usecase.Update(factory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "updated")
}

// Delete an animal
// @Summary Delete an animal by its GUID
// @Description Delete an animal by its GUID
// @Tags animals
// @Produce  json
// @Param guid path string true "GUID"
// @Success 200 {string} string "deleted"
// @Failure 500 {object} map[string]string "error"
// @Router /animals/{guid} [delete]
func (a AnimalHandler) Delete(c *gin.Context) {
	guid := c.Param("guid")
	err := a.Usecase.Delete(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}
