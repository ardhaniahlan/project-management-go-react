package models

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	InternalID     int64      `json:"internal_id" db:"internal_id" gorm:"primaryKey;column:internal_id;autoIncrement"`
	PublicID       uuid.UUID  `json:"public_id" db:"public_id" gorm:"not null;column:public_id"`
	ListInternalID int64      `json:"list_internal_id" db:"list_internal_id" gorm:"not null;column:list_internal_id"`
	Title          string     `json:"title" db:"title" gorm:"not null;column:title"`
	Description    string     `json:"description" db:"description" gorm:"column:description"`
	DueDate        *time.Time `json:"due_date,omitempty" db:"due_date" gorm:"column:due_date"`
	Position       int        `json:"position" db:"position" gorm:"column:position"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at" gorm:"column:created_at"`
}
