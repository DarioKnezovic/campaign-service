package api

import (
	"github.com/DarioKnezovic/campaign-service/api/handlers"
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, campaignService campaign.CampaignService) {
	userHandler := &handlers.CampaignHandler{
		CampaignService: campaignService,
	}

	// Define routes with parameters using "{parameter}" syntax
	router.HandleFunc("/api/campaign/all", userHandler.GetAllCampaignsHandler).Methods("GET")
	router.HandleFunc("/api/campaign/create", userHandler.CreateNewCampaignHandler).Methods("POST")
	router.HandleFunc("/api/campaign/single/{id}", userHandler.GetSingleCampaignHandler).Methods("GET")
	router.HandleFunc("/api/campaign/update/{id}", userHandler.UpdateCampaignHandler).Methods("PUT")
	router.HandleFunc("/api/campaign/delete/{id}", userHandler.DeleteCampaignHandler).Methods("DELETE")
	router.HandleFunc("/api/user/{id}/campaigns", userHandler.GetUserCampaignsHandler).Methods("GET")
}
