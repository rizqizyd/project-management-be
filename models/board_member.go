package models

import "time"

type BoardMember struct {
	BoardID  int64     `json:"board_internal_id" db:"board_internal_id" gorm:"primaryKey;column:board_internal_id"`
	UserID   int64     `json:"user_internal_id" db:"user_internal_id" gorm:"primaryKey;column:user_internal_id"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}

/**
 * Composite primary key consists of BoardID and UserID
 * Composite primary keys are defined using gorm tags on the struct fields.
 * Composite keys help ensure that each user can only be a member of a board once.
 */
