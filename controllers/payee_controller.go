package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type PayeeController struct {
	payeeService *services.PayeeService
}

func NewPayeeController(payeeService *services.PayeeService) *PayeeController {
	return &PayeeController{payeeService: payeeService}
}

func (pc *PayeeController) GetPayeeByID(ctx *gin.Context) {
	id := ctx.Param("id")

	payeeId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	user, err := pc.payeeService.GetPayeeByID(context.Background(), payeeId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (pc *PayeeController) GetPayees(ctx *gin.Context) {

	payees, err := pc.payeeService.GetPayees(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payees)

}
func (pc *PayeeController) CreatePayee(ctx *gin.Context) {
	var payee models.Payee
	if err := ctx.ShouldBindJSON(&payee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedOrg, err := pc.payeeService.CreatePayee(ctx, &payee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedOrg)
}

func (pc *PayeeController) UpdatePayee(ctx *gin.Context) {
	var payee models.Payee
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&payee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPayee, err := pc.payeeService.UpdatePayee(ctx, &payee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedPayee)
}

func (pc *PayeeController) DeletePayee(ctx *gin.Context) {

	id := ctx.Param("id")

	payeeId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := pc.payeeService.DeletePayee(ctx, payeeId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
