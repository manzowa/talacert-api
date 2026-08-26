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

// Create godoc
// @Summary      Create document
// @Description  Creates a new official document.
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        document  body  dto.CreateDocumentRequest  true  "Document data"
// @Success      201  {object}  interface{}  "Document created successfully"
// @Failure      400  {object}  interface{}  "Invalid request payload"
// @Failure      500  {object}  interface{}  "Failed to create document"
// @Router       /api/v1/documents [post]
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

// GetAll godoc
// @Summary      Get all documents
// @Description  Retrieves all documents.
// @Tags         Documents
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  interface{}  "Documents retrieved successfully"
// @Failure      500  {object}  interface{}  "Failed to retrieve documents"
// @Router       /api/v1/documents [get]
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

// GetDocument godoc
// @Summary      Get document
// @Description  Retrieves a document by its unique identifier.
// @Tags         Documents
// @Produce      json
// @Security     BearerAuth
// @Param        document_id  path  string  true  "Document ID"
// @Success      200  {object}  dto.DocumentResponse  "Document retrieved successfully"
// @Failure      400  {object}  interface{}  "Invalid document ID"
// @Failure      404  {object}  interface{}  "Document not found"
// @Router       /api/v1/documents/{document_id} [get]
func (h *DocumentHandler) GetByDocument(
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
		CreatedAt:  document.CreatedAt,
	}

	utils.Ok(cxt, "Document retrieved successfully", response)

}

// Update godoc
// @Summary      Update document
// @Description  Updates an existing document. Only the provided fields need to be supplied.
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        document_id  path  string  true  "Document ID"
// @Param        document  body  dto.UpdateDocumentRequest  true  "Document data"
// @Success      200  {object}  interface{}  "Document updated successfully"
// @Failure      400  {object}  interface{}  "Invalid document ID or request payload"
// @Failure      500  {object}  interface{}  "Failed to update document"
// @Router       /api/v1/documents/{document_id} [patch]
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

// Delete godoc
// @Summary      Delete document
// @Description  Deletes an existing document by its unique identifier.
// @Tags         Documents
// @Produce      json
// @Security     BearerAuth
// @Param        document_id  path  string  true  "Document ID"
// @Success      200  {object}  interface{}  "Document deleted successfully"
// @Failure      400  {object}  interface{}  "Invalid document ID"
// @Failure      404  {object}  interface{}  "Document not found"
// @Router       /api/v1/documents/{document_id} [delete]
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

// Delete godoc
// @Summary      Delete document
// @Description  Deletes an existing document by its unique identifier.
// @Tags         Documents
// @Produce      json
// @Security     BearerAuth
// @Param        document_id  path  string  true  "Document ID"
// @Success      200  {object}  interface{}  "Document deleted successfully"
// @Failure      400  {object}  interface{}  "Invalid document ID"
// @Failure      404  {object}  interface{}  "Document not found"
// @Router       /api/v1/documents/by-hash/{hash} [GET]
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
		CreatedAt:  document.CreatedAt,
	}

	utils.Ok(
		cxt, "Document retrieved successfully",
		response,
	)
}

// Check godoc
// @Summary      Verify document
// @Description  Verifies the authenticity and validity of a document using its document ID.
// @Tags         Documents
// @Accept       json
// @Produce      json
// @Param        request  body  dto.VerificationRequest  true  "Document verification request"
// @Success      200  {object}  dto.VerificationResponse  "Document verification result"
// @Failure      400  {object}  interface{}  "Invalid request payload or document ID"
// @Failure      404  {object}  interface{}  "Document not found"
// @Router       /api/v1/documents/verify [post]
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
