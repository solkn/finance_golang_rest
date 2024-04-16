package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type TxCategoryController struct {
	txCategoryService *services.TxCategoryService
}

func NewTxCategoryController(txCategoryService *services.TxCategoryService) *TxCategoryController {
	return &TxCategoryController{txCategoryService: txCategoryService}
}

func (tc *TxCategoryController) GetTxCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")

	txcId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	transaction, err := tc.txCategoryService.GetTxCategoryByID(context.Background(), txcId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (tc *TxCategoryController) GetTxCategorys(ctx *gin.Context) {

	txcs, err := tc.txCategoryService.GetTxCategorys(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, txcs)

}
func (tc *TxCategoryController) CreateTxCategory(ctx *gin.Context) {
	var txc models.TransactionCategory
	if err := ctx.ShouldBindJSON(&txc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedTxc, err := tc.txCategoryService.CreateTxCategory(ctx, &txc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedTxc)
}

func (tc *TxCategoryController) UpdateTxCategory(ctx *gin.Context) {
	var txc models.TransactionCategory
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&txc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTxc, err := tc.txCategoryService.UpdateTxCategory(ctx, &txc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedTxc)
}

func (tc *TxCategoryController) DeleteTxCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	txcId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := tc.txCategoryService.DeleteTxCategory(ctx, txcId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
