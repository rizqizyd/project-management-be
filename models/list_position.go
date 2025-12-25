package models

import (
	"github.com/google/uuid"
	"github.com/rizqizyd/project-management-be/models/types"
)

type ListPosition struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id"`
	BoardID    int64     `json:"board_internal_id" db:"board_internal_id" gorm:"column:board_internal_id"`
	// ListOrder  UUIDArray `json:"list_order" db:"list_order" gorm:"type:uuid[]"`
	ListOrder types.UUIDArray `json:"list_order" db:"list_order"`
}
