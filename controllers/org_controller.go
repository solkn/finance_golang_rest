package controllers

import (
	"context"
	"finance/models"
	"finance/services"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type OrgController struct {
	orgService *services.OrgService
}

func NewOrgController(orgService *services.OrgService) *OrgController {
	return &OrgController{orgService: orgService}
}

func (oc *OrgController) GetOrgByID(ctx *gin.Context) {
	id := ctx.Param("id")

	orgId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	user, err := oc.orgService.GetOrgByID(context.Background(), orgId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (oc *OrgController) GetOrgs(ctx *gin.Context) {

	orgs, err := oc.orgService.GetOrgs(context.Background())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orgs)

}
func (oc *OrgController) CreateOrg(ctx *gin.Context) {
	var org models.Org
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedOrg, err := oc.orgService.CreateOrg(ctx, &org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusCreated, savedOrg)
}

func (oc *OrgController) UpdateOrg(ctx *gin.Context) {
	var org models.Org
	// user_id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedOrg, err := oc.orgService.UpdateOrg(ctx, &org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedOrg)
}

func (oc *OrgController) DeleteOrg(ctx *gin.Context) {

	id := ctx.Param("id")

	orgId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	err2 := oc.orgService.DeleteOrg(ctx, orgId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
