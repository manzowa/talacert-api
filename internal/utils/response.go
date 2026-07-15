package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse représente le format standard des réponses de l'API.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success renvoie une réponse de succès.
func Success(cxt *gin.Context, statusCode int, message string, data interface{}) {
	cxt.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created renvoie une réponse HTTP 201.
func Created(cxt *gin.Context, message string, data interface{}) {
	Success(cxt, http.StatusCreated, message, data)
}

// Ok renvoie une réponse HTTP 200.
func Ok(cxt *gin.Context, message string, data interface{}) {
	Success(cxt, http.StatusOK, message, data)
}

// Error renvoie une réponse d'erreur.
func Error(cxt *gin.Context, statusCode int, message string, err interface{}) {
	cxt.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// BadRequest renvoie une réponse HTTP 400.
func BadRequest(cxt *gin.Context, message string, err interface{}) {
	Error(cxt, http.StatusBadRequest, message, err)
}

// Unauthorized renvoie une réponse HTTP 401.
func Unauthorized(cxt *gin.Context, message string) {
	Error(cxt, http.StatusUnauthorized, message, nil)
}

// Forbidden renvoie une réponse HTTP 403.
func Forbidden(cxt *gin.Context, message string) {
	Error(cxt, http.StatusForbidden, message, nil)
}

// NotFound renvoie une réponse HTTP 404.
func NotFound(cxt *gin.Context, message string) {
	Error(cxt, http.StatusNotFound, message, nil)
}

// Conflict renvoie une réponse HTTP 409.
func Conflict(cxt *gin.Context, message string, err interface{}) {
	Error(cxt, http.StatusConflict, message, err)
}

// InternalServerError renvoie une réponse HTTP 500.
func InternalServerError(cxt *gin.Context, message string, err interface{}) {
	Error(cxt, http.StatusInternalServerError, message, err)
}
