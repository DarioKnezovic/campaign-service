package handlers

import (
	"encoding/json"
	"github.com/DarioKnezovic/campaign-service/config"
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/DarioKnezovic/campaign-service/pkg/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CampaignHandler struct {
	CampaignService campaign.CampaignService
}

func (h *CampaignHandler) GetAllCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	cfg := config.LoadConfig()

	if len(token) != 2 {
		responseError := map[string]string{
			"error": "Invalid Authorization header",
		}
		util.SendJSONResponse(w, http.StatusUnauthorized, responseError)
		return
	}

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

func (h *CampaignHandler) CreateNewCampaignHandler(w http.ResponseWriter, r *http.Request) {
	var newCampaign campaign.Campaign
	token := strings.Split(r.Header.Get("Authorization"), " ")
	cfg := config.LoadConfig()

	err := json.NewDecoder(r.Body).Decode(&newCampaign)
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusBadRequest, nil)
		return
	}

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

	newCampaign.CustomerID = userId
	savedCampaign, err := h.CampaignService.CreateNewCampaign(newCampaign)
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	util.SendJSONResponse(w, http.StatusOK, savedCampaign)
}

func (h *CampaignHandler) GetSingleCampaignHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	receivedCampaign, err := h.CampaignService.GetSingleCampaign(id)
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	util.SendJSONResponse(w, http.StatusOK, receivedCampaign)
}
