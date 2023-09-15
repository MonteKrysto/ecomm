package entity

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Active      bool
	CreatorID   uuid.UUID
	Creator     User
	Items       []Item
	Medias      []Media
	Contributors []User `gorm:"many2many:campaign_contributors;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
