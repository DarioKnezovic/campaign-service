package service

import (
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/DarioKnezovic/campaign-service/internal/campaign/repository"
)

type CampaignService struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(campaignRepository repository.CampaignRepository) *CampaignService {
	return &CampaignService{
		campaignRepository: campaignRepository,
	}
}

func (s *CampaignService) FetchAllCampaigns() ([]campaign.Campaign, error) {

	return nil, nil
}
