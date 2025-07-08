package controller

import (
	"errors"
	"go-gin-project/api/service"
	"go-gin-project/data"
	"go-gin-project/helper"
	"go-gin-project/helper/responsejson"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	registerRequest := data.RegisterRequest{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	authResponse, err := c.authService.Register(registerRequest)
	if err != nil {
		if errors.Is(err, helper.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, helper.ErrUniqueViolation) {
			responsejson.Conflict(ctx, err, "Username already exists")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to register user")
		return
	}

	ctx.SetCookie("access_token", authResponse.Token, 24*60*60, "/", "", false, true)

	responsejson.Created(ctx, authResponse, "User registered successfully")
}

func (c *AuthController) Login(ctx *gin.Context) {
	loginRequest := data.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		responsejson.BadRequest(ctx, err, "Invalid request body")
		return
	}

	authResponse, err := c.authService.Login(loginRequest)
	if err != nil {
		if errors.Is(err, helper.ErrValidation) {
			responsejson.BadRequest(ctx, err, "Validation error")
			return
		}
		if errors.Is(err, helper.ErrNotFound) {
			responsejson.Unauthorized(ctx, "Invalid username or password")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to login")
		return
	}

	ctx.SetCookie("access_token", authResponse.Token, 24*60*60, "/", "", false, true)

	responsejson.Success(ctx, authResponse, "Login successful")
}

func (c *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "", false, true)
	responsejson.Success(ctx, nil, "Logout successful")
}

func (c *AuthController) Profile(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		responsejson.Unauthorized(ctx, "User not authenticated")
		return
	}

	userUUID, ok := userId.(uuid.UUID)
	if !ok {
		responsejson.InternalServerError(ctx, nil, "Invalid user ID format")
		return
	}

	userResponse, err := c.authService.GetUserById(userUUID)
	if err != nil {
		if errors.Is(err, helper.ErrNotFound) {
			responsejson.NotFound(ctx, "User not found")
			return
		}
		responsejson.InternalServerError(ctx, err, "Failed to get user profile")
		return
	}

	responsejson.Success(ctx, userResponse, "Profile retrieved successfully")
}
