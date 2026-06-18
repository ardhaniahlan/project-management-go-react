package models

import (
	"project-management-be/models/types"

	"github.com/google/uuid"
)

type CardPosition struct {
	InternalID     int64           `json:"internal_id" db:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID       uuid.UUID       `json:"public_id" db:"public_id" gorm:"type:uuid;column:public_id;not null"`
	ListInternalID int64           `json:"list_internal_id" db:"list_internal_id" gorm:"column:list_internal_id;not null"`
	CardOrder      types.UUIDArray `json:"card_order" db:"card_order" gorm:"type:uuid[]"`
}
