package handlers

import (
	"errors"

	"talacert-api/internal/constants"
	"talacert-api/internal/dto"
	"talacert-api/internal/services"
	"talacert-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUser(service *services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Create(cxt *gin.Context) {

	var req dto.CreateUserRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(
			cxt,
			"Invalid request payload",
			err.Error(),
		)
		return
	}

	err := h.Service.Create(
		cxt.Request.Context(),
		req,
	)

	if err != nil {

		if errors.Is(err, services.ErrEmailExists) ||
			errors.Is(err, services.ErrUsernameExists) {

			utils.Conflict(
				cxt,
				"User already exists",
				err.Error(),
			)
			return
		}

		utils.InternalServerError(
			cxt,
			"Failed to create user",
			err.Error(),
		)

		return
	}

	utils.Created(
		cxt,
		"User created successfully",
		nil,
	)
}

func (h *UserHandler) GetAll(cxt *gin.Context) {

	users, err := h.Service.GetAll(
		cxt.Request.Context(),
	)

	if err != nil {

		utils.InternalServerError(
			cxt,
			"Failed to retrieve users",
			err.Error(),
		)

		return
	}

	utils.Ok(
		cxt,
		"Users retrieved successfully",
		users,
	)
}

func (h *UserHandler) GetByID(cxt *gin.Context) {

	id, err := parseID(cxt)

	if err != nil {
		utils.BadRequest(
			cxt,
			"Invalid user ID",
			err.Error(),
		)
		return
	}

	user, err := h.Service.GetUserByID(
		cxt.Request.Context(),
		id,
	)

	if err != nil {

		if errors.Is(err, services.ErrUserNotFound) {

			utils.NotFound(
				cxt,
				"User not found",
			)

			return
		}

		utils.InternalServerError(
			cxt,
			"Failed to retrieve user",
			err.Error(),
		)

		return
	}

	utils.Ok(
		cxt,
		"User retrieved successfully",
		user,
	)
}

func (h *UserHandler) Update(cxt *gin.Context) {

	id, err := parseID(cxt)

	if err != nil {
		utils.BadRequest(cxt, "Invalid user ID", err.Error())
		return
	}

	var req dto.UpdateUserRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {

		utils.BadRequest(
			cxt,
			"Invalid request payload",
			err.Error(),
		)

		return
	}

	err = h.Service.Update(
		cxt.Request.Context(),
		id,
		req,
	)

	if err != nil {

		switch {

		case errors.Is(err, services.ErrUserNotFound):

			utils.NotFound(
				cxt,
				"User not found",
			)

		case errors.Is(err, services.ErrEmailExists),
			errors.Is(err, services.ErrUsernameExists):

			utils.Conflict(
				cxt,
				"User already exists",
				err.Error(),
			)

		default:

			utils.InternalServerError(
				cxt,
				"Failed to update user",
				err.Error(),
			)
		}

		return
	}

	utils.Ok(
		cxt,
		"User updated successfully",
		nil,
	)
}

func (h *UserHandler) Delete(cxt *gin.Context) {

	id, err := parseID(cxt)

	if err != nil {
		utils.BadRequest(cxt, "Invalid user ID", err.Error())
		return
	}

	err = h.Service.Delete(
		cxt.Request.Context(),
		id,
	)

	if err != nil {

		if errors.Is(err, services.ErrUserNotFound) {

			utils.NotFound(
				cxt,
				"User not found",
			)

			return
		}

		utils.InternalServerError(
			cxt,
			"Failed to delete user",
			err.Error(),
		)

		return
	}

	utils.Ok(
		cxt,
		"User deleted successfully",
		nil,
	)
}

func (h *UserHandler) ChangePassword(cxt *gin.Context) {

	id, err := parseID(cxt)

	if err != nil {
		utils.BadRequest(cxt, "Invalid user ID", err.Error())
		return
	}

	var req dto.UserChangePasswordRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {

		utils.BadRequest(
			cxt,
			"Invalid request payload",
			err.Error(),
		)

		return
	}

	err = h.Service.ChangePassword(
		cxt.Request.Context(),
		id,
		req.Password,
	)

	if err != nil {

		utils.InternalServerError(
			cxt,
			"Failed to change password",
			err.Error(),
		)

		return
	}

	utils.Ok(
		cxt,
		"Password changed successfully",
		nil,
	)
}

func (h *UserHandler) ChangeRole(cxt *gin.Context) {

	id, err := parseID(cxt)

	if err != nil {
		utils.BadRequest(cxt, "Invalid user ID", err.Error())
		return
	}

	var req dto.UserChangeRoleRequest

	if err := cxt.ShouldBindJSON(&req); err != nil {

		utils.BadRequest(
			cxt,
			"Invalid request payload",
			err.Error(),
		)

		return
	}

	err = h.Service.ChangeRole(
		cxt.Request.Context(),
		id,
		constants.Role(req.Role),
	)

	if err != nil {

		utils.InternalServerError(
			cxt,
			"Failed to change role",
			err.Error(),
		)

		return
	}

	utils.Ok(
		cxt,
		"Role changed successfully",
		nil,
	)
}
