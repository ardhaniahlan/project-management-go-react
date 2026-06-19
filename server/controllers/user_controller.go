package controllers

import (
	"project-management-be/models"
	"project-management-be/services"
	"project-management-be/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Invalid Request Body", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Failed Register user", err.Error())
	}

	userResponse := models.MapToUserResponse(user)
	return utils.Created(ctx, "User registered successfully", userResponse)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	user, err := c.service.Login(req.Email, req.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Invalid credentials", err.Error())
	}

	token, _ := utils.GenerateTokenJWT(user.InternalID, user.Role, user.Email, user.PublicID) 
	refreshToken, _ := utils.GenerateRefreshTokenJWT(user.InternalID)

	userResponse := models.MapToUserResponse(user)
	return utils.Success(ctx, "Login successful", fiber.Map{
		"access_token": token,
		"refresh_token": refreshToken,
		"user": userResponse,
	})
}