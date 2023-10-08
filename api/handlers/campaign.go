package handlers

import (
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/DarioKnezovic/campaign-service/pkg/util"
	"net/http"
)

type CampaignHandler struct {
	CampaignService campaign.CampaignService
}

func (h *CampaignHandler) GetAllCampaignsHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"response": "puši ga",
	}
	util.SendJSONResponse(w, http.StatusOK, response)
}
