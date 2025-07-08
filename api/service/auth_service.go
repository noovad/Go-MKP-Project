package service

import (
	"errors"
	"go-gin-project/api/repository"
	"go-gin-project/data"
	"go-gin-project/helper"
	"go-gin-project/model"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(request data.RegisterRequest) (data.AuthResponse, error)
	Login(request data.LoginRequest) (data.AuthResponse, error)
	GetUserById(id uuid.UUID) (data.UserResponse, error)
}

type authServiceImpl struct {
	authRepository repository.AuthRepository
	validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &authServiceImpl{
		authRepository: authRepository,
		validate:       validate,
	}
}

func (s *authServiceImpl) Register(request data.RegisterRequest) (data.AuthResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return data.AuthResponse{}, helper.WrapValidation(err)
		
	}

	_, err = s.authRepository.FindByUsername(request.Username)
	if err == nil {
		return data.AuthResponse{}, helper.ErrUniqueViolation
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return data.AuthResponse{}, err
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return data.AuthResponse{}, err
	}

	user := model.User{
		Id:        uuid.New(),
		Username:  request.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = s.authRepository.CreateUser(user)
	if err != nil {
		return data.AuthResponse{}, err
	}

	token, err := helper.GenerateToken(user.Id, user.Username)
	if err != nil {
		return data.AuthResponse{}, err
	}

	response := data.AuthResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}

	return response, nil
}

func (s *authServiceImpl) Login(request data.LoginRequest) (data.AuthResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return data.AuthResponse{}, helper.WrapValidation(err)
	}

	user, err := s.authRepository.FindByUsername(request.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data.AuthResponse{}, helper.ErrNotFound
		}
		return data.AuthResponse{}, err
	}

	err = helper.CheckPassword(user.Password, request.Password)
	if err != nil {
		return data.AuthResponse{}, helper.ErrNotFound
	}

	token, err := helper.GenerateToken(user.Id, user.Username)
	if err != nil {
		return data.AuthResponse{}, err
	}

	response := data.AuthResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}

	return response, nil
}

func (s *authServiceImpl) GetUserById(id uuid.UUID) (data.UserResponse, error) {
	user, err := s.authRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data.UserResponse{}, helper.ErrNotFound
		}
		return data.UserResponse{}, err
	}

	response := data.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return response, nil
}
