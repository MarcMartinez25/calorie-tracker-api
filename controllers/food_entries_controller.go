package controllers

import (
	"net/http"

	"calorie-tracker-api/services"

	"github.com/gin-gonic/gin"
)

type FoodEntriesController struct {
	service *services.FoodEntryService
}

func NewFoodEntriesController(service *services.FoodEntryService) *FoodEntriesController {
	return &FoodEntriesController{
		service: service,
	}
}

// GetFoodEntry handles GET requests for a single food entry by ID
func (c *FoodEntriesController) GetFoodEntry(ctx *gin.Context) {
	id := ctx.Param("id")

	entry, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch food entry",
		})
		return
	}

	if entry == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Food entry not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"entry": entry,
	})
}

// GetFoodEntries handles GET requests for food entries by user ID
func (c *FoodEntriesController) GetFoodEntries(ctx *gin.Context) {
	userId := ctx.Param("userId")

	entries, err := c.service.GetByUserID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch food entries",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"entries": entries,
	})
}
