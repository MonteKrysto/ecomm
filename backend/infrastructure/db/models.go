package db

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email        string     `gorm:"type:varchar(100);uniqueIndex"`
	Password     string     `gorm:"size:100"`
	CreatorProfile CreatorProfile
	ContributorProfile ContributorProfile
	CustomerProfile CustomerProfile
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreatorProfile struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID  `gorm:"type:uuid"`
	Avatar     string
	Bio        string     `gorm:"size:500"`
	Campaigns  []Campaign
}

type ContributorProfile struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID  `gorm:"type:uuid"`
	Avatar     string
	Bio        string     `gorm:"size:500"`
	Campaigns  []Campaign // Campaigns the contributor is part of
}

type CustomerProfile struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID  `gorm:"type:uuid"`
}

type Campaign struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title           string     `gorm:"size:100"`
	Description     string     `gorm:"size:500"`
	IsActive        bool
	CreatorID       uuid.UUID  `gorm:"type:uuid"`
	Contributors    []ContributorProfile `gorm:"many2many:campaign_contributors;"`
	Items           []Item
	Photos          []Photo
	Videos          []Video
	StartDate       time.Time
	EndDate         time.Time
	Private         bool
	Invitations     []Invitation
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Item struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CampaignID      uuid.UUID  `gorm:"type:uuid"`
	Title           string     `gorm:"size:100"`
	Description     string     `gorm:"size:500"`
	Price           float64
	Quantity        int
	Photos          []Photo
	Videos          []Video
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Photo struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	URL        string
}

type Video struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	URL       string
}

type Message struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	SenderID  uuid.UUID  `gorm:"type:uuid"`
	ReceiverID uuid.UUID `gorm:"type:uuid"`
	Content   string     `gorm:"size:500"`
	Read      bool
	CreatedAt time.Time
}

type Notification struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID  `gorm:"type:uuid"`
	Content   string     `gorm:"size:500"`
	Read      bool
	CreatedAt time.Time
}

type Invitation struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CampaignID  uuid.UUID  `gorm:"type:uuid"`
	Email       string     `gorm:"type:varchar(100)"`
	InvitedAt   time.Time
	AcceptedAt  time.Time
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&CreatorProfile{},
		&ContributorProfile{},
		&CustomerProfile{},
		&Campaign{},
		&Item{},
		&Photo{},
		&Video{},
		&Message{},
		&Notification{},
		&Invitation{},
	)
	return err
}
