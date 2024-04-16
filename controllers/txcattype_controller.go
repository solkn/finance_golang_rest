package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type TxCatTypeController struct {
	txCatTypeService *services.TxCatTypeService
}

func NewTxCatTypeController(txCatTypeService *services.TxCatTypeService) *TxCatTypeController {
	return &TxCatTypeController{txCatTypeService: txCatTypeService}
}

func (tc *TxCatTypeController) GetTxCatTypeByID(ctx *gin.Context) {
	id := ctx.Param("id")

	txcId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	transaction, err := tc.txCatTypeService.GetTxCatTypeByID(context.Background(), txcId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (tc *TxCatTypeController) GetTxCatTypes(ctx *gin.Context) {

	txcs, err := tc.txCatTypeService.GetTxCatTypes(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, txcs)

}
func (tc *TxCatTypeController) CreateTxCatType(ctx *gin.Context) {
	var txc models.TxCategoryType
	if err := ctx.ShouldBindJSON(&txc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedTxc, err := tc.txCatTypeService.CreateTxCatType(ctx, &txc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedTxc)
}

func (tc *TxCatTypeController) UpdateTxCatType(ctx *gin.Context) {
	var txc models.TxCategoryType
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&txc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTxc, err := tc.txCatTypeService.UpdateTxCatType(ctx, &txc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedTxc)
}

func (tc *TxCatTypeController) DeleteTxCatType(ctx *gin.Context) {

	id := ctx.Param("id")

	txcId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := tc.txCatTypeService.DeleteTxCatType(ctx, txcId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
