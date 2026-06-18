package models

import "time"

type BoardMember struct {
	BoardInternalID uint64    `json:"board_internal_id" db:"board_internal_id" gorm:"not null;primaryKey;column:board_internal_id"`
	UserInternalID  uint64    `json:"user_internal_id" db:"user_internal_id" gorm:"not null;primaryKey;column:user_internal_id"`
	JoinAt          time.Time `json:"join_at" db:"join_at"`

	// Relasi
	Board           Board     `json:"-" gorm:"foreignKey:BoardInternalID;references:InternalID"`
	User            User      `json:"-" gorm:"foreignKey:UserInternalID;references:InternalID"`
}