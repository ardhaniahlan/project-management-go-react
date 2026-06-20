package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
}

type ResponsePaginate struct {
	Status       string         `json:"status"`
	ResponseCode int            `json:"response_code"`
	Message      string         `json:"message,omitempty"`
	Data         interface{}    `json:"data,omitempty"`
	Meta         PaginationMeta `json:"meta,omitempty"`
	Error        string         `json:"error,omitempty"`
}

type PaginationMeta struct {
	Page       int64  `json:"page" example:"1"`
	Limit      int64  `json:"limit" example:"10"`
	Total      int64  `json:"total" example:"100"`
	TotalPages int64  `json:"total_pages" example:"10"`
	Filter     string `json:"filter" example:"nama=ardhan"`
	Sort       string `json:"sort" example:"-id"`
}

func SuccessPaginate(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(ResponsePaginate{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func NotFoundPaginate(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusNotFound).JSON(ResponsePaginate{
		Status:       "not found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:       "success",
		ResponseCode: fiber.StatusCreated,
		Message:      message,
		Data:         data,
	})
}

func BadRequest(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:       "error bad request",
		ResponseCode: fiber.StatusBadRequest,
		Message:      message,
		Error:        err,
	})
}

func NotFound(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:       "error not found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Error:        err,
	})
}

func InternalServerError(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:       "error",
		ResponseCode: fiber.StatusInternalServerError,
		Message:      message,
		Error:        err,
	})
}

func Unauthorized(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:       "error",
		ResponseCode: fiber.StatusUnauthorized,
		Message:      message,
		Error:        err,
	})
}
