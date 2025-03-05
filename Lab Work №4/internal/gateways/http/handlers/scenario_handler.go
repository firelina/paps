package handlers

import (
	"github.com/gin-gonic/gin"
	"marine/internal/models"
	"marine/internal/usecase"
	"net/http"
	"strconv"
)

type ScenarioHandler struct {
	scenarioUsecase usecase.ScenarioService
}

func NewScenarioHandler(s usecase.ScenarioService) *ScenarioHandler {
	return &ScenarioHandler{s}
}

func (h *ScenarioHandler) CreateScenario(c *gin.Context) {
	var newScenario models.CreateScenarioRequest
	if err := c.ShouldBindJSON(&newScenario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	scenarioID, err := h.scenarioUsecase.CreateScenario(ctx, &models.Scenario{
		Title:       newScenario.Title,
		Description: newScenario.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.CreateScenarioResponse{ID: scenarioID})
}

func (h *ScenarioHandler) GetScenarios(c *gin.Context) {
	ctx := c.Request.Context()
	scenarios, err := h.scenarioUsecase.GetAllScenarios(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, scenarios)
}

func (h *ScenarioHandler) DeleteScenario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	err = h.scenarioUsecase.DeleteScenario(ctx, id)
	if err != nil {
		if err == usecase.ErrScenarioNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Сценарий не найден."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *ScenarioHandler) UpdateScenario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	var updatedScenario models.CreateScenarioRequest
	if err := c.ShouldBindJSON(&updatedScenario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	scenario, err := h.scenarioUsecase.UpdateScenario(ctx, id, &models.Scenario{
		Title:       updatedScenario.Title,
		Description: updatedScenario.Description,
	})
	if err != nil {
		if err == usecase.ErrScenarioNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Сценарий не найден."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, scenario)
}
