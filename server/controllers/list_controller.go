package controllers

import (
	"project-management-be/dto"
	"project-management-be/models"
	"project-management-be/services"
	"project-management-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ListController struct {
	listService services.ListService
}

func NewListController(listService services.ListService) *ListController {
	return &ListController{listService: listService}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error {
	req := new(dto.NewListRequest)
	if err := ctx.BodyParser(req); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	boardPublicID, err := uuid.Parse(req.BoardPublicID)
	if err != nil {
		return utils.BadRequest(ctx, "Invalid board public ID", err.Error())
	}

	list := &models.List{
		Title:         req.Title,
    BoardPublicID: boardPublicID,
	}

	if err := c.listService.Create(list); err != nil {
		return utils.InternalServerError(ctx, "Failed to create list", err.Error())
	}
	
	return utils.Success(ctx, "List created successfully", list)
}
	