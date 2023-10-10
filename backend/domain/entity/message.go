package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID uuid.UUID
	Content    string
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	ReadStatus bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
