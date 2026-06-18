package models

import (
	"time"
	"github.com/google/uuid"
)

type List struct {
	InternalID      int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID        uuid.UUID `json:"public_id" db:"public_id" gorm:"not null"`
	BoardPublicID   uuid.UUID `json:"board_public_id" db:"board_public_id" gorm:"not null;column:board_public_id"`
	Title           string    `json:"title" db:"title" gorm:"not null;column:title"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	BoardInternalID int64     `json:"-" db:"board_internal_id"`

	// Relasi
	Board           Board     `json:"-" gorm:"foreignKey:BoardInternalID;references:InternalID"`
}