package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type RegisterController struct {
	registerService *services.RegisterService
}

func NewRegisterController(registerService *services.RegisterService) *RegisterController {
	return &RegisterController{registerService: registerService}
}

func (rc *RegisterController) Registers(ctx *gin.Context) {

	registers, err := rc.registerService.GetRegisters(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, registers)

}

func (rc *RegisterController) GetRegisterByID(ctx *gin.Context) {
	id := ctx.Param("id")

	registerId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	user, err := rc.registerService.GetRegisterByID(context.Background(), registerId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (rc *RegisterController) CreateRegister(ctx *gin.Context) {
	var register models.Register
	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedRegister, err := rc.registerService.CreateRegister(ctx, &register)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedRegister)
}

func (rc *RegisterController) UpdateRegister(ctx *gin.Context) {
	var register models.Register

	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPayee, err := rc.registerService.UpdateRegister(ctx, &register)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedPayee)
}

func (rc *RegisterController) DeleteRegister(ctx *gin.Context) {

	id := ctx.Param("id")

	registerId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := rc.registerService.DeleteRegister(ctx, registerId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
