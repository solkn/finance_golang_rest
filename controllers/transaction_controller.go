package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (tc *TransactionController) GetTransactionByID(ctx *gin.Context) {
	id := ctx.Param("id")

	transactionId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	transaction, err := tc.transactionService.GetTransactionByID(context.Background(), transactionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) GetTransactions(ctx *gin.Context) {

	tags, err := tc.transactionService.GetTransactions(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tags)

}

// func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
// 	var transaction models.Transaction
// 	if err := ctx.ShouldBindJSON(&transaction); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	savedTransaction, err := tc.transactionService.CreateTransaction(ctx, &transaction)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, savedTransaction)
// }

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {

	// Access file path from form parameter (replace "filepath" with actual parameter name)
	filePath, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing file path"})
		return
	}

	fmt.Println("file:", filePath.Filename)

	src, err := filePath.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error opening file"})
		return
	}
	defer src.Close()

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "uploaded-file-*.xlsx") // Adjust the file extension and name prefix as needed
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating temporary file"})
		return
	}
	defer tempFile.Close()

	// Copy the uploaded file content to the temporary file
	_, err = io.Copy(tempFile, src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error copying file content"})
		return
	}

	transaction := tc.transactionService.CreateTransaction(context.Background(), tempFile.Name())

	fmt.Println("transaction:", transaction)
	if transaction != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("error registering transactions: %w", transaction)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "transaction registered successfully"})
}

func (tc *TransactionController) UpdateTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	id := ctx.Param("id")
	transactionId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTransaction, err := tc.transactionService.UpdateTransaction(ctx, transactionId, &transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedTransaction)
}

func (tc *TransactionController) DeleteTransaction(ctx *gin.Context) {

	id := ctx.Param("id")

	transactionId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := tc.transactionService.DeleteTransaction(ctx, transactionId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
