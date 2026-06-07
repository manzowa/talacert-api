package handlers

import (
	"errors"

	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"talacert-api/internal/dto"
	"talacert-api/internal/mapper"
	"talacert-api/internal/services"
)

type DocumentHandler struct {
	Service *services.DocumentService
}

func New(service *services.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		Service: service,
	}
}

func (d *DocumentHandler) GetDocumentsHandler(c *gin.Context) {
	documents, err := d.Service.GetDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.ToResponseDocuments(documents))
}

func (d *DocumentHandler) PostDocumentHandler(c *gin.Context) {
	var docRequest dto.DocumentRequest

	if err := c.ShouldBindJSON(&docRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	document := mapper.ToModelDocument(&docRequest)

	if err := d.Service.CreateDocument(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapper.ToResponseDocument(&document))
}

func (d *DocumentHandler) GetDocumentHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}
	document, err := d.Service.GetByDocumentID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.ToResponseDocument(document))
}

func (d *DocumentHandler) PutDocumentHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}
	var docRequest dto.DocumentRequest

	if err := c.ShouldBindJSON(&docRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doc := mapper.ToModelDocument(&docRequest)

	updatedDoc, err := d.Service.UpdateByDocumentID(id, &doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.ToResponseDocument(updatedDoc))
}

func (h *DocumentHandler) DeleteDocumentHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid document ID"})
		return
	}

	err := h.Service.DeleteByDocumentID(id)
	if err != nil {

		// 🔥 cas "not found" (optionnel mais recommandé)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "document deleted successfully",
		"document_id": id,
	})
}

func (d *DocumentHandler) GetDocumentByHashHandler(c *gin.Context) {
	hash := c.Param("hash")
	document, err := d.Service.GetByDocumentHash(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.ToResponseDocument(document))
}
