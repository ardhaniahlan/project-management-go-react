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
	
func (c *ListController) UpdateList(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	
	req := new(dto.UpdateListRequest)

	if err := ctx.BodyParser(req); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid list public ID", err.Error())
	}

	existingList, err := c.listService.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List not found", "List not found")
	}

	list := &models.List{
		Title: req.Title,
	}

	list.InternalID = existingList.InternalID
	list.PublicID = existingList.PublicID
	
	if err := c.listService.Update(list); err != nil {
		return utils.InternalServerError(ctx, "Failed to update list", err.Error())
	}

	updatedList, err := c.listService.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List not found", "List not found")
	}
	
	return utils.Success(ctx, "List updated successfully", updatedList)
}

func (c *ListController) GetListOnBoard(ctx *fiber.Ctx) error {
	boardPublicID := ctx.Params("board_id")
	if _, err := uuid.Parse(boardPublicID); err != nil {
		return utils.BadRequest(ctx, "Invalid board public ID", err.Error())
	}

	list, err := c.listService.GetByBoardID(boardPublicID)
	if err != nil {
		return utils.InternalServerError(ctx, "Failed to get lists", err.Error())
	}
	
	return utils.Success(ctx, "Lists retrieved successfully", list)
}

func (c *ListController) DeleteList(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid list public ID", err.Error())
	}

	list, err := c.listService.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List not found", "List not found")
	}

	if err := c.listService.Delete(uint(list.InternalID)); err != nil {
		return utils.InternalServerError(ctx, "Failed to delete list", err.Error())
	}
	return utils.Success(ctx, "List deleted successfully", publicID)
}

func (c *ListController) UpdateListPosition(ctx *fiber.Ctx) error {
	boardPublicID := ctx.Params("board_id")

	if _, err := uuid.Parse(boardPublicID); err != nil {
		return utils.BadRequest(ctx, "Invalid board public ID", err.Error())
	}

	var positionUUID []uuid.UUID
	if err := ctx.BodyParser(&positionUUID); err != nil {
		var positionString []string
		if err := ctx.BodyParser(&positionString); err != nil {
			return utils.BadRequest(ctx, "Invalid position format", err.Error())
		}
		for _, position := range positionString {
			if _, err := uuid.Parse(position); err != nil {
				return utils.BadRequest(ctx, "Invalid position", err.Error())
			}
			positionUUID = append(positionUUID, uuid.MustParse(position))
		}
	}

	if err := c.listService.UpdatePosition(boardPublicID, positionUUID); err != nil {
		return utils.InternalServerError(ctx, "Failed to update list position", err.Error())
	}
	return utils.Success(ctx, "List position updated successfully", nil)
}