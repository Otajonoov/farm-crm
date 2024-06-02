package v1

import (
	"farm-crm/internal/domain"
	"farm-crm/internal/port/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnimalHandler struct {
	usecase domain.AnimalUseCase
	factory domain.Factory
}

func (a AnimalHandler) Create(c *gin.Context) {
	var animal model.Animal

	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	animalFactory := a.factory.ParseToDomainFor(animal.AnimalTypeGuid, animal.Descriptive, animal.Weight, animal.IsIll, animal.IsFed)
	err := a.usecase.Create(animalFactory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, animal)
}

func (a AnimalHandler) RetrieveById(c *gin.Context) {
	guid := c.Query("guid")
	zikr, err := a.usecase.GetByGuid(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zikr)
}

func (a AnimalHandler) RetrieveList(c *gin.Context) {
	zikrs, err := a.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, zikrs)
}

func (a AnimalHandler) Update(c *gin.Context) {
	var animal model.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	factory := a.factory.ParseToDomainFor(animal.AnimalTypeGuid, animal.Descriptive, animal.Weight, animal.IsIll, animal.IsFed)
	err := a.usecase.Update(factory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "updated")
}

func (a AnimalHandler) Delete(c *gin.Context) {
	guid := c.Query("guid")
	err := a.usecase.Delete(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}
