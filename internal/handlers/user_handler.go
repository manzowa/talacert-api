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

// Create godoc
// @Summary      Create user
// @Description  Creates a new user.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body  dto.CreateUserRequest  true  "User data"
// @Success      201  {object}  interface{}  "User created successfully"
// @Failure      400  {object}  interface{}  "Invalid request payload"
// @Failure      409  {object}  interface{}  "User already exists"
// @Failure      500  {object}  interface{}  "Failed to create user"
// @Router       /api/v1/users [post]
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

// GetAll godoc
// @Summary      Get all users
// @Description  Retrieves all users.
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  interface{}  "Users retrieved successfully"
// @Failure      500  {object}  interface{}  "Failed to retrieve users"
// @Router       /api/v1/users [get]
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

// GetByID godoc
// @Summary      Get user by ID
// @Description  Retrieves a user by their unique identifier.
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  interface{}  "User retrieved successfully"
// @Failure      400  {object}  interface{}  "Invalid user ID"
// @Failure      404  {object}  interface{}  "User not found"
// @Failure      500  {object}  interface{}  "Failed to retrieve user"
// @Router       /api/v1/users/{id} [get]
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

// Update godoc
// @Summary      Update user
// @Description  Updates an existing user.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path  int                    true  "User ID"
// @Param        request  body  dto.UpdateUserRequest  true  "User data"
// @Success      200  {object}  interface{}  "User updated successfully"
// @Failure      400  {object}  interface{}  "Invalid user ID or request payload"
// @Failure      404  {object}  interface{}  "User not found"
// @Failure      409  {object}  interface{}  "User already exists"
// @Failure      500  {object}  interface{}  "Failed to update user"
// @Router       /api/v1/users/{id} [patch]
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

// Delete godoc
// @Summary      Delete user
// @Description  Deletes a user by their unique identifier.
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  interface{}  "User deleted successfully"
// @Failure      400  {object}  interface{}  "Invalid user ID"
// @Failure      404  {object}  interface{}  "User not found"
// @Failure      500  {object}  interface{}  "Failed to delete user"
// @Router       /api/v1/users/{id} [delete]
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

// ChangePassword godoc
// @Summary      Change user password
// @Description  Changes the password of an existing user.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path  int                         true  "User ID"
// @Param        request  body  dto.UserChangePasswordRequest  true  "New password"
// @Success      200  {object}  interface{}  "Password changed successfully"
// @Failure      400  {object}  interface{}  "Invalid user ID or request payload"
// @Failure      500  {object}  interface{}  "Failed to change password"
// @Router       /api/v1/users/{id}/password [patch]
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

// ChangeRole godoc
// @Summary      Change user role
// @Description  Changes the role of an existing user.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path  int                       true  "User ID"
// @Param        request  body  dto.UserChangeRoleRequest true  "New user role"
// @Success      200  {object}  interface{}  "Role changed successfully"
// @Failure      400  {object}  interface{}  "Invalid user ID or request payload"
// @Failure      500  {object}  interface{}  "Failed to change role"
// @Router       /api/v1/users/{id}/role [patch]
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
