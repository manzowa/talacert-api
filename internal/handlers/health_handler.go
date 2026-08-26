package handlers

import (
	"talacert-api/internal/dto"
	"talacert-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Checks whether the Talacert API is healthy and available.
// @Tags         Health
// @Produce      json
// @Success      200  {object}  dto.HealthResponse  "Health check successful"
// @Failure      500  {object}  interface{}  "Health check failed"
// @Router       /health [get]
func (h *HealthHandler) HealthCheck(cxt *gin.Context) {

	utils.Ok(
		cxt,
		"Talacert API Health check successful",
		dto.HealthResponse{
			Status: "ok",
		},
	)
}
