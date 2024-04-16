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

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	user, err := uc.userService.GetUserByID(context.Background(), userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {

	users, err := uc.userService.GetUsers(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)

}
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := uc.userService.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("User:", savedUser)

	ctx.JSON(http.StatusCreated, savedUser)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	// user_id := ctx.Param("id")

	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	

	updatedUser, err := uc.userService.UpdateUser(ctx, userId,&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := uc.userService.DeleteUser(ctx, userId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("mm,", err2)

	ctx.Status(http.StatusNoContent)

}
