package entity

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	URL        string
	CampaignID uuid.UUID
	Campaign   Campaign
	ItemID     uuid.UUID
	Item       Item
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
