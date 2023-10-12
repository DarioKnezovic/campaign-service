package handlers

import (
	"encoding/json"
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"github.com/DarioKnezovic/campaign-service/pkg/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type CampaignHandler struct {
	CampaignService campaign.CampaignService
}

func (h *CampaignHandler) GetAllCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		log.Printf("Error during getting user id from token")
		responseError := map[string]string{
			"error": err.Error(),
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	campaigns, err := h.CampaignService.FetchAllCampaigns(userId)
	if err != nil {
		log.Printf("Error during fetching campaigns %e", err)
		responseError := map[string]string{
			"error": "Internal server error",
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	util.SendJSONResponse(w, http.StatusOK, campaigns)
}

func (h *CampaignHandler) CreateNewCampaignHandler(w http.ResponseWriter, r *http.Request) {
	var newCampaign campaign.Campaign
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		log.Printf("Error during getting user id from token")
		responseError := map[string]string{
			"error": err.Error(),
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newCampaign)
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusBadRequest, nil)
		return
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

	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		log.Printf("Error during getting user id from token")
		responseError := map[string]string{
			"error": err.Error(),
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	receivedCampaign, err := h.CampaignService.GetSingleCampaign(id, userId)
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	util.SendJSONResponse(w, http.StatusOK, receivedCampaign)
}

func (h *CampaignHandler) UpdateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	var campaignUpdate campaign.Campaign

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&campaignUpdate)
	if err != nil {
		log.Print(err)
		util.SendJSONResponse(w, http.StatusBadRequest, nil)
		return
	}

	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		log.Printf("Error during getting user id from token")
		responseError := map[string]string{
			"error": err.Error(),
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	err = h.CampaignService.UpdateCampaign(campaignUpdate, id, userId)
	if err != nil {
		// TODO: add response body
		util.SendJSONResponse(w, http.StatusInternalServerError, nil)
		return
	}

	util.SendJSONResponse(w, http.StatusOK, nil)
}
