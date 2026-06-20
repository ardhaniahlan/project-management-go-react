package dto

import "time"

type CreateBoardRequest struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	DueDate     *time.Time `json:"due_date,omitempty" validate:"omitempty"`
}