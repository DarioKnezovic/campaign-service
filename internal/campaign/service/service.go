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

func (s *CampaignService) FetchAllCampaigns(userId uint) ([]campaign.Campaign, error) {

	campaigns, err := s.campaignRepository.GetAllCampaigns(userId)
	return campaigns, err
}

func (s *CampaignService) CreateNewCampaign(newCampaign campaign.Campaign) (*campaign.Campaign, error) {
	savedCampaign, err := s.campaignRepository.InsertNewCampaign(newCampaign)
	if err != nil {
		return nil, err
	}

	return savedCampaign, nil
}

func (s *CampaignService) GetSingleCampaign(campaignId int) (campaign.Campaign, error) {

	campaignReceived, err := s.campaignRepository.FetchCampaignById(campaignId)
	return campaignReceived, err
}
