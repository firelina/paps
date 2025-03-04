package http

import (
	"github.com/gin-gonic/gin"
	"marine/internal/gateways/http/handlers"
)

func setupRouter(r *gin.Engine, useCases UseCases) {
	userHandler := handlers.NewUserHandler(useCases.User)

	scenarioHandler := handlers.NewScenarioHandler(useCases.Scenario)

	materialHandler := handlers.NewMaterialHandler(useCases.Material)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/api/v1/users", userHandler.RegisterUser)
	r.GET("/api/v1/users/:id", userHandler.GetUser)

	r.POST("/api/v1/scenarios", scenarioHandler.CreateScenario)
	r.GET("/api/v1/scenarios", scenarioHandler.GetScenarios)
	r.PUT("/api/v1/scenarios/:id", scenarioHandler.UpdateScenario)
	r.DELETE("/api/v1/scenarios/:id", scenarioHandler.DeleteScenario)

	r.POST("/api/v1/materials", materialHandler.UploadMaterial)
	r.GET("/api/v1/materials", materialHandler.GetMaterials)
}
