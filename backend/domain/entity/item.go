package entity

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string
	Description string
	Quantity   int
	CampaignID uuid.UUID
	Campaign   Campaign
	Medias     []Media
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
