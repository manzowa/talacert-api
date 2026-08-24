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

func (h *HealthHandler) HealthCheck(cxt *gin.Context) {

	utils.Ok(
		cxt,
		"Talacert API Health check successful",
		dto.HealthResponse{
			Status: "ok",
		},
	)
}
