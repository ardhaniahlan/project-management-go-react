package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	InternalID uint64         `json:"internal_id" db:"internal_id" gorm:"primarykey"`
	PublicID   uuid.UUID      `json:"public_id" db:"public_id"`
	Name       string         `json:"name" db:"name" gorm:"not null"`
	Email      string         `json:"email" db:"email" gorm:"not null;unique"`
	Password   string         `json:"password" db:"password" gorm:"not null;column:password"`
	Role       string         `json:"role" db:"role" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}