package handlers

import (
	"github.com/gin-gonic/gin"
	"marine/internal/models"
	"marine/internal/usecase"
	"net/http"
)

type MaterialHandler struct {
	materialUsecase usecase.MaterialService
}

func NewMaterialHandler(m usecase.MaterialService) *MaterialHandler {
	return &MaterialHandler{m}
}

func (h *MaterialHandler) UploadMaterial(c *gin.Context) {
	var uploadMaterial models.CreateMaterialRequest
	if err := c.ShouldBindJSON(&uploadMaterial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	materialID, err := h.materialUsecase.UploadMaterial(ctx, &models.Material{
		Name: uploadMaterial.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.CreateMaterialResponse{ID: materialID})
}

func (h *MaterialHandler) GetMaterials(c *gin.Context) {
	ctx := c.Request.Context()
	materials, err := h.materialUsecase.GetAllMaterials(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, materials)
}
