package handlers

import (
	"errors"

	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"talacert-api/internal/dto"
	"talacert-api/internal/services"
	"talacert-api/internal/utils"
)

type DocumentHandler struct {
	Service *services.DocumentService
}

func NewDocument(service *services.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		Service: service,
	}
}

func (h *DocumentHandler) Create(
	cxt *gin.Context,
) {
	var req dto.CreateDocumentRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(
			cxt, "Invalid request payload",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)
		return
	}

	err := h.Service.Create(cxt, &req)

	if err != nil {
		utils.InternalServerError(
			cxt, "Failed to create document",
			errors.New(http.StatusText(http.StatusInternalServerError)),
		)
		return
	}
	utils.Created(
		cxt, "Document created successfully",
		nil,
	)
}

func (h *DocumentHandler) GetAll(
	cxt *gin.Context,
) {
	documents, err := h.Service.GetAll(cxt)
	if err != nil {
		utils.InternalServerError(
			cxt, "Failed to retrieve documents",
			errors.New(http.StatusText(http.StatusInternalServerError)),
		)

		return
	}
	utils.Ok(
		cxt, "Documents retrieved successfully",
		documents,
	)
}

func (h *DocumentHandler) GetByDocumentID(
	cxt *gin.Context,
) {
	documentID := cxt.Param("document_id")

	if documentID == "" {
		utils.BadRequest(
			cxt, "Invalid document ID",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)

		return
	}
	document, err := h.Service.GetByDocumentID(cxt, documentID)
	if err != nil {
		utils.NotFound(cxt, "Document not found")

		return
	}

	var response = dto.DocumentResponse{
		DocumentID: document.DocumentID,
		OwnerName:  document.OwnerName,
		Type:       document.Type,
		Issuer:     document.Issuer,
		Hash:       document.Hash,
		Status:     document.Status,
		CreatedAt:  document.CreatedAt,
	}

	utils.Ok(cxt, "Document retrieved successfully", response)

}

func (h *DocumentHandler) Update(
	cxt *gin.Context,
) {
	documentID := cxt.Param("document_id")

	if documentID == "" {
		utils.BadRequest(
			cxt, "Invalid document ID",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)

		return
	}
	var req dto.UpdateDocumentRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(
			cxt, "Invalid request payload",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)

		return
	}

	err := h.Service.Update(cxt, documentID, req)

	if err != nil {
		utils.InternalServerError(
			cxt, "Failed to update document",
			errors.New(http.StatusText(http.StatusInternalServerError)),
		)

		return
	}

	utils.Ok(cxt, "Document updated successfully", nil)

}

func (h *DocumentHandler) Delete(
	cxt *gin.Context,
) {
	documentID := cxt.Param("document_id")

	if documentID == "" {
		utils.BadRequest(
			cxt, "Invalid document ID",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)

		return
	}

	err := h.Service.Delete(cxt, documentID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(cxt, "Document not found")
			return
		}
	}

	utils.Ok(cxt, "Document deleted successfully", nil)
}

func (h *DocumentHandler) GetByHash(
	cxt *gin.Context,
) {
	hash := cxt.Param("hash")

	if hash == "" {
		utils.BadRequest(
			cxt, "Invalid hash",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)

		return
	}

	document, err := h.Service.GetByHash(cxt, hash)

	if err != nil {
		utils.InternalServerError(
			cxt, "Failed to retrieve document",
			errors.New(http.StatusText(http.StatusInternalServerError)),
		)

		return
	}

	var response = dto.DocumentResponse{
		DocumentID: document.DocumentID,
		OwnerName:  document.OwnerName,
		Type:       document.Type,
		Issuer:     document.Issuer,
		Hash:       document.Hash,
		Status:     document.Status,
		CreatedAt:  document.CreatedAt,
	}

	utils.Ok(
		cxt, "Document retrieved successfully",
		response,
	)
}

func (h *DocumentHandler) Check(
	cxt *gin.Context,
) {
	var req dto.VerificationRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(
			cxt, "Invalid request payload",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)
		return
	}

	if req.DocumentID == "" {
		utils.BadRequest(
			cxt, "Invalid document ID",
			errors.New(http.StatusText(http.StatusBadRequest)),
		)

		return
	}
	document, err := h.Service.GetByDocumentID(cxt, req.DocumentID)

	if err != nil {
		utils.NotFound(cxt, "Document not found")

		return
	}

	response := dto.VerificationResponse{
		Status: document.Status,
		Data: dto.VerificationData{
			Owner:  document.OwnerName,
			Type:   document.Type,
			Issuer: document.Issuer,
		},
	}
	message := "Document Valid"
	if response.Status != "valid" {
		message = "Document Invalid"
	}

	utils.Ok(cxt, message, response)

}
