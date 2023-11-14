package campaign

import (
	"gorm.io/gorm"
	"time"
)

// Campaign represents a campaign entity
type Campaign struct {
	CampaignID uint      `gorm:"primaryKey;autoIncrement:true" json:"campaign_id"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	Name       string    `gorm:"not null" json:"name"`
	Image      string    `json:"image"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	FirstName   string         `gorm:"not null" json:"first_name"`
	LastName    string         `gorm:"not null" json:"last_name"`
	Email       string         `gorm:"unique;not null" json:"email"`
	CustomerKey string         `json:"customer_key"`
	Password    string         `gorm:"not null" json:"password"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type CampaignService interface {
	FetchAllCampaigns(userId uint) ([]Campaign, error)
	CreateNewCampaign(newCampaign Campaign) (*Campaign, error)
	GetSingleCampaign(campaignId int, userId uint) (Campaign, error)
	UpdateCampaign(campaignUpdatePayload Campaign, campaignId int, userId uint) error
	DeleteCampaign(campaignId int) error
	InitCampaign(customerKey string) (Campaign, error)
}
