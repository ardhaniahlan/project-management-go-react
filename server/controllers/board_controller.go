package controllers

import (
	"project-management-be/dto"
	"project-management-be/models"
	"project-management-be/services"
	"project-management-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type BoardController struct {
	boardService services.BoardService
}

func NewBoardController(bService services.BoardService) *BoardController {
	return &BoardController{boardService: bService}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	req := new(dto.CreateBoardRequest)
	if err := ctx.BodyParser(req); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	claims, ok := ctx.Locals("user").(jwt.MapClaims)
	if !ok {
		return utils.Unauthorized(ctx, "Unauthorized", "Gagal membaca token")
	}

	userPublicID, ok := claims["public_id"].(string)
	if !ok || userPublicID == "" {
		return utils.Unauthorized(ctx, "Unauthorized", "ID tidak valid")
	}

	board := &models.Board{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
	}

	if err := c.boardService.Create(board, userPublicID); err != nil {
		return utils.InternalServerError(ctx, "Failed to create board", err.Error())
	}

	return utils.Created(ctx, "Board created successfully", board)
}
