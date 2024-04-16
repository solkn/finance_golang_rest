package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type LocationController struct {
	locationService *services.LocationService
}

func NewLocationController(locationService *services.LocationService) *LocationController {
	return &LocationController{locationService: locationService}
}

func (lc *LocationController) GetLocationByID(ctx *gin.Context) {
	id := ctx.Param("id")

	locationId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	user, err := lc.locationService.GetLocationByID(context.Background(), locationId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (lc *LocationController) GetLocations(ctx *gin.Context) {

	locations, err := lc.locationService.GetLocations(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, locations)

}
func (lc *LocationController) CreateLocation(ctx *gin.Context) {
	var location models.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedLocation, err := lc.locationService.CreateLocation(ctx, &location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, savedLocation)
}

func (lc *LocationController) UpdateLocation(ctx *gin.Context) {
	var location models.Location
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := lc.locationService.UpdateLocation(ctx, &location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (lc *LocationController) DeleteLocation(ctx *gin.Context) {

	id := ctx.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := lc.locationService.DeleteLocation(ctx, userId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("mm,", err2)

	ctx.Status(http.StatusNoContent)

}
