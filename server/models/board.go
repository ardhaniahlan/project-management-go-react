package models

import (
	"time"
	"github.com/google/uuid"
)

type Board struct {
	InternalID      uint64         `json:"internal_id" db:"internal_id" gorm:"primarykey;autoIncrement"`
	PublicID        uuid.UUID      `json:"public_id" db:"public_id" gorm:"not null"`
	Title           string         `json:"title" db:"title" gorm:"not null"`
	Description     string         `json:"description" db:"description" gorm:"not null"`
	OwnerInternalID uint64         `json:"owner_internal_id" db:"owner_internal_id" gorm:"not null;column:owner_internal_id"`
	OwnerPublicID   uuid.UUID      `json:"owner_public_id" db:"owner_public_id" gorm:"not null"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	DueDate         *time.Time     `json:"due_date,omitempty" db:"due_date"`
}