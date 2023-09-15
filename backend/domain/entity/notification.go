package entity

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Content    string
	UserID     uuid.UUID
	ReadStatus bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
