package campaign

import (
	"time"
)

// Campaign represents a campaign entity
type Campaign struct {
	CampaignID uint      `gorm:"primaryKey;autoIncrement:true" json:"campaign_id"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	Name       string    `gorm:"not null" json:"name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type CampaignService interface {
	FetchAllCampaigns(userId uint) ([]Campaign, error)
	CreateNewCampaign(newCampaign Campaign) (*Campaign, error)
	GetSingleCampaign(campaignId int, userId uint) (Campaign, error)
	UpdateCampaign(campaignUpdatePayload Campaign, campaignId int, userId uint) error
}
