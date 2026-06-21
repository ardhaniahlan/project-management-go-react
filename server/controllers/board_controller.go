package controllers

import (
	"project-management-be/dto"
	"project-management-be/models"
	"project-management-be/services"
	"project-management-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func (c *BoardController) UpdateBoard(ctx *fiber.Ctx) error {
	claims, _ := ctx.Locals("user").(jwt.MapClaims)
	userPublicID := claims["public_id"].(string)

	boardID := ctx.Params("id")
	if _, err := uuid.Parse(boardID); err != nil {
		return utils.BadRequest(ctx, "Invalid public ID", err.Error())
	}

	req := new(dto.UpdateBoardRequest)
	if err := ctx.BodyParser(req); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	board := &models.Board{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
	}

	if err := c.boardService.Update(boardID, board, userPublicID); err != nil {
		return utils.InternalServerError(ctx, "Failed to update board", err.Error())
	}

	return utils.Success(ctx, "Board updated successfully", board)
}

func (c *BoardController) AddBoardMembers(ctx *fiber.Ctx) error {
	claims, _ := ctx.Locals("user").(jwt.MapClaims)
	actorPublicID := claims["public_id"].(string)

	boardPublicID := ctx.Params("id")

	req := new(dto.AddMembersRequest)
	if err := ctx.BodyParser(req); err != nil {
		return utils.BadRequest(ctx, "Format request tidak valid", err.Error())
	}

	if err := c.boardService.AddMembers(boardPublicID, req.UserIDs, actorPublicID); err != nil {
		return utils.InternalServerError(ctx, "Gagal menambahkan anggota", err.Error())
	}

	return utils.Success(ctx, "Anggota berhasil ditambahkan", nil)
}
