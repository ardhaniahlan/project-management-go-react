package models

import (
	"project-management-be/models/types"

	"github.com/google/uuid"
)

type ListPosition struct {
	InternalID      int64           `json:"internal_id" db:"internal_id" gorm:"primaryKey;column:internal_id;autoIncrement"`
	PublicID        uuid.UUID       `json:"public_id" db:"public_id" gorm:"not null;column:public_id"`
	BoardInternalID int64           `json:"-" db:"board_internal_id" gorm:"not null;column:board_internal_id"`
	ListOrder       types.UUIDArray `json:"list_order" db:"list_order" gorm:"type:uuid[]"`

	// Relasi
	Board           Board           `json:"-" gorm:"foreignKey:BoardInternalID;references:InternalID"`
}