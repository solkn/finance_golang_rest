package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type TxTagController struct {
	txTagService *services.TxTagService
}

func NewTxTagController(txTagService *services.TxTagService) *TxTagController {
	return &TxTagController{txTagService: txTagService}
}

func (tc *TxTagController) GetTxTagByID(ctx *gin.Context) {
	id := ctx.Param("id")

	tagId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	transaction, err := tc.txTagService.GetTxTagByID(context.Background(), tagId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (tc *TxTagController) GetTxTags(ctx *gin.Context) {

	tags, err := tc.txTagService.GetTxTags(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tags)

}
func (tc *TxTagController) CreateTxTag(ctx *gin.Context) {
	var tag models.TxTag
	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedOrg, err := tc.txTagService.CreateTxTag(ctx, &tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedOrg)
}

func (tc *TxTagController) UpdateTxTag(ctx *gin.Context) {
	var tag models.TxTag
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTag, err := tc.txTagService.UpdateTxTag(ctx, &tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedTag)
}

func (tc *TxTagController) DeleteTxTag(ctx *gin.Context) {

	id := ctx.Param("id")

	tagId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := tc.txTagService.DeleteTxTag(ctx, tagId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
