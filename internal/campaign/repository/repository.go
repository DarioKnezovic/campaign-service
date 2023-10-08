package repository

import (
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	GetAllCampaigns() ([]campaign.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepository{
		db: db,
	}
}

func (c campaignRepository) GetAllCampaigns() ([]campaign.Campaign, error) {
	//TODO implement me
	panic("implement me")
}
