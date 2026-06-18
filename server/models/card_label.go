package models

type CardLabel struct {
	CardInternalID  int64 `json:"card_internal_id" db:"card_internal_id" gorm:"column:card_internal_id;not null;primaryKey"`
	LabelInternalID int64 `json:"label_internal_id" db:"label_internal_id" gorm:"column:label_internal_id;not null;primaryKey"`

	// Relasi
	Card            Card  `json:"-" gorm:"foreignKey:CardInternalID;references:InternalID"`
	Label           Label `json:"-" gorm:"foreignKey:LabelInternalID;references:InternalID"`
}