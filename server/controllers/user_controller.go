package controllers

import (
	"math"
	"project-management-be/models"
	"project-management-be/services"
	"project-management-be/utils"
	"strconv"

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

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx, "User not found", err.Error())
	}

	userResponse := models.MapToUserResponse(user)
	return utils.Success(ctx, "User found", userResponse)
}

func (c *UserController) GetUsersPaginate(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	filter := ctx.Query("filter")
	sort := ctx.Query("sort")

	users, count, err := c.service.GetAllPaginate(filter, sort, limit, offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Failed to fetch users", err.Error())
	}

	userResponse := make([]models.UserResponse, 0)
	for _, user := range users {
		userResponse = append(userResponse, models.MapToUserResponse(&user))
	}

	meta := utils.PaginationMeta{
		Page:       int64(page),
		Limit:      int64(limit),
		Total:      count,
		TotalPages: int64(math.Ceil(float64(count) / float64(limit))),
		Filter:     filter,
		Sort:       sort,
	}

	return utils.SuccessPaginate(ctx, "Users fetched successfully", userResponse, meta)
}