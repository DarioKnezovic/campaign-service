package repository

import (
	"errors"
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	GetAllCampaigns(userId uint) ([]campaign.Campaign, error)
	InsertNewCampaign(newCampaign campaign.Campaign) (*campaign.Campaign, error)
	FetchCampaignById(campaignId int) (campaign.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepository{
		db: db,
	}
}

func (c *campaignRepository) GetAllCampaigns(userId uint) ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	err := c.db.Where("customer_id = ?", userId).Find(&campaigns).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return campaigns, nil
}

func (c *campaignRepository) InsertNewCampaign(newCampaign campaign.Campaign) (*campaign.Campaign, error) {

	err := c.db.Create(&newCampaign).Error
	if err != nil {
		return nil, err
	}

	return &newCampaign, nil
}

func (c *campaignRepository) FetchCampaignById(campaignId int) (campaign.Campaign, error) {
	var foundedCampaign campaign.Campaign
	err := c.db.Where("campaign_id = ?", campaignId).Find(&foundedCampaign).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return campaign.Campaign{}, nil
		}
		return campaign.Campaign{}, err
	}

	return foundedCampaign, nil
}
