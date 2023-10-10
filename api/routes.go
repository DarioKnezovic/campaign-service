package api

import (
	"github.com/DarioKnezovic/campaign-service/api/handlers"
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"net/http"
)

func RegisterRoutes(campaignService campaign.CampaignService) {
	userHandler := &handlers.CampaignHandler{
		CampaignService: campaignService,
	}

	http.HandleFunc("/api/campaign/all", userHandler.GetAllCampaignsHandler)
	http.HandleFunc("/api/campaign/create", userHandler.CreateNewCampaignHandler)
}
