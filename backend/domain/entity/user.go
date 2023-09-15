package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name           string
	Email          string
	Password       string
	ProfilePicture string
	Roles          []Role     `gorm:"many2many:user_roles;"`
	Campaigns      []Campaign `gorm:"foreignKey:CreatorID"`
	Messages       []Message
	Notifications  []Notification
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
