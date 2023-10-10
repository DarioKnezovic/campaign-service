package handlers

import (
	"github.com/DarioKnezovic/campaign-service/config"
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/DarioKnezovic/campaign-service/pkg/util"
	"log"
	"net/http"
	"strings"
)

type CampaignHandler struct {
	CampaignService campaign.CampaignService
}

func (h *CampaignHandler) GetAllCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	cfg := config.LoadConfig()

	// TODO: In the future add GRPC call to User service to check token
	claims, err := util.VerifyJWT(token[1], []byte(cfg.JWTSecretKey))
	if err != nil {
		log.Printf("Error during verifying JWT: %v", err)
		responseError := map[string]string{
			"error": "Error during verifying JWT",
		}
		util.SendJSONResponse(w, http.StatusUnauthorized, responseError)
	}

	userId := claims.Id
	if userId == 0 {
		log.Print("User ID is not available from JWT token")
		responseError := map[string]string{
			"error": "Undefined User ID from Authorization token",
		}
		util.SendJSONResponse(w, http.StatusUnauthorized, responseError)
	}

	campaigns, err := h.CampaignService.FetchAllCampaigns(userId)
	if err != nil {
		log.Printf("Error during fetching campaigns %e", err)
		responseError := map[string]string{
			"error": "Internal server error",
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
	}

	util.SendJSONResponse(w, http.StatusOK, campaigns)
}
