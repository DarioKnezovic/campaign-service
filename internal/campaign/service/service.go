package service

import (
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/DarioKnezovic/campaign-service/internal/campaign/repository"
	"log"
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

func (s *CampaignService) GetSingleCampaign(campaignId int, userId uint) (campaign.Campaign, error) {

	campaignReceived, err := s.campaignRepository.FetchCampaignById(campaignId, userId)
	return campaignReceived, err
}

func (s *CampaignService) UpdateCampaign(campaignUpdatePayload campaign.Campaign, campaignId int, userId uint) error {
	return s.campaignRepository.UpdateCampaignById(campaignUpdatePayload, campaignId, userId)
}

func (s *CampaignService) DeleteCampaign(campaignId int) error {
	selectedCampaign, err := s.campaignRepository.FetchCampaignById(campaignId, 0)
	if err != nil {
		return err
	}

	return s.campaignRepository.DeleteCampaignById(selectedCampaign)
}

func (s *CampaignService) InitCampaign(customerKey string) (campaign.Campaign, error) {
	userId, err := s.campaignRepository.FindUserUsingCustomerKey(customerKey)
	if err != nil {
		return campaign.Campaign{}, err
	}

	log.Println("Fetch campaign for customer_id=", userId)

	fetchedCampaign, err := s.campaignRepository.FindCampaignByCustomerID(userId)
	if err != nil {
		return campaign.Campaign{}, err
	}

	return fetchedCampaign, nil
}
