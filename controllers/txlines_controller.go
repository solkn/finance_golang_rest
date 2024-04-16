package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type TxLinesController struct {
	txLinesService *services.TxLinesService
}

func NewTxLinesController(txLinesService *services.TxLinesService) *TxLinesController {
	return &TxLinesController{txLinesService: txLinesService}
}

func (tc *TxLinesController) GetTxLinesByID(ctx *gin.Context) {
	id := ctx.Param("id")

	lineId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	transaction, err := tc.txLinesService.GetTxLinesByID(context.Background(), lineId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (tc *TxLinesController) GetTxLines(ctx *gin.Context) {

	lines, err := tc.txLinesService.GetTxTags(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, lines)

}
func (tc *TxLinesController) CreateTxLines(ctx *gin.Context) {
	var line models.TxLines
	if err := ctx.ShouldBindJSON(&line); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedOrg, err := tc.txLinesService.CreateTxLines(ctx, &line)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedOrg)
}

func (tc *TxLinesController) UpdateTxLines(ctx *gin.Context) {
	var line models.TxLines
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&line); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTag, err := tc.txLinesService.UpdateTxLines(ctx, &line)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedTag)
}

func (tc *TxLinesController) DeleteTxLines(ctx *gin.Context) {

	id := ctx.Param("id")

	lineId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := tc.txLinesService.DeleteTxLines(ctx, lineId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
