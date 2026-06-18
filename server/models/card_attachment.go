package models

import (
	"time"
	"github.com/google/uuid"
)

type CardAttachment struct {
	InternalID     int64     `json:"internal_id" db:"internal_id" gorm:"column:internal_id;primaryKey;autoIncrement"`
	PublicID       uuid.UUID `json:"public_id" db:"public_id" gorm:"column:public_id;not null"`
	CardInternalID int64     `json:"card_internal_id" db:"card_internal_id" gorm:"column:card_internal_id;not null"`
	File           string    `json:"file" db:"file" gorm:"column:file;not null"`
	UserInternalID int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id;not null"`
	CreatedAt      time.Time `json:"created_at" db:"created_at" gorm:"column:created_at;not null"`

	// Relasi
	Card           Card      `json:"-" gorm:"foreignKey:CardInternalID;references:InternalID"`
	User           User      `json:"-" gorm:"foreignKey:UserInternalID;references:InternalID"`
}