package services

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"talacert-api/internal/constants"
	"talacert-api/internal/dto"
	"talacert-api/internal/models"
	"talacert-api/internal/repositories"
)

var (
	ErrEmailExists    = errors.New("email already exists")
	ErrUsernameExists = errors.New("username already exists")
	ErrUserNotFound   = errors.New("user not found")
	ErrHashPassword   = errors.New("failed to hash password")
)

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(
	userRepository repositories.UserRepositoryInterface,
) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) GetAll(
	ctx context.Context,
) ([]dto.UserResponse, error) {

	users, err := s.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, s.toUserResponse(&user))
	}

	return response, nil
}

func (s *UserService) Create(
	ctx context.Context,
	req dto.CreateUserRequest,
) error {

	existingEmail, err := s.UserRepository.FindByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return err
	}

	if existingEmail != nil {
		return ErrEmailExists
	}

	existingUsername, err := s.UserRepository.FindByUsername(
		ctx,
		req.Username,
	)

	if err != nil {
		return err
	}

	if existingUsername != nil {
		return ErrUsernameExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return ErrHashPassword
	}

	role := constants.RoleUser

	if req.Role != "" {
		role = constants.Role(strings.ToUpper(req.Role))
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
	}

	return s.UserRepository.Create(ctx, user)
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id uint,
) (*dto.UserResponse, error) {

	user, err := s.UserRepository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	response := s.toUserResponse(user)

	return &response, nil
}

func (s *UserService) Update(
	ctx context.Context,
	id uint,
	req dto.UpdateUserRequest,
) error {

	user, err := s.UserRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	if req.Username != "" &&
		req.Username != user.Username {

		existing, err := s.UserRepository.FindByUsername(
			ctx,
			req.Username,
		)

		if err != nil {
			return err
		}

		if existing != nil {
			return ErrUsernameExists
		}

		user.Username = req.Username
	}

	if req.Email != "" &&
		req.Email != user.Email {

		existing, err := s.UserRepository.FindByEmail(
			ctx,
			req.Email,
		)

		if err != nil {
			return err
		}

		if existing != nil {
			return ErrEmailExists
		}

		user.Email = req.Email
	}

	if req.Role != "" {
		user.Role = constants.Role(strings.ToUpper(req.Role))
	}

	return s.UserRepository.Update(ctx, user)
}

func (s *UserService) Delete(
	ctx context.Context,
	id uint,
) error {

	user, err := s.UserRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	return s.UserRepository.Delete(ctx, id)
}

func (s *UserService) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	return s.UserRepository.FindByEmail(ctx, email)
}

func (s *UserService) GetUserByUsername(
	ctx context.Context,
	username string,
) (*models.User, error) {

	return s.UserRepository.FindByUsername(ctx, username)
}

func (s *UserService) GetUserByEmailOrUsername(
	ctx context.Context,
	identifier string,
) (*models.User, error) {

	user, err := s.UserRepository.FindByEmail(
		ctx,
		identifier,
	)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return s.UserRepository.FindByUsername(
		ctx,
		identifier,
	)
}

func (s *UserService) ChangeRole(
	ctx context.Context,
	id uint,
	role constants.Role,
) error {

	user, err := s.UserRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	user.Role = role

	return s.UserRepository.Update(ctx, user)
}

func (s *UserService) ChangePassword(
	ctx context.Context,
	id uint,
	newPassword string,
) error {

	user, err := s.UserRepository.FindByID(ctx, id)

	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return ErrHashPassword
	}

	user.Password = string(hashedPassword)

	return s.UserRepository.Update(ctx, user)
}

func (s *UserService) toUserResponse(
	user *models.User,
) dto.UserResponse {

	return dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
	}
}
