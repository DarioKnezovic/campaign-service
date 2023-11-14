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
		log.Print(err)
		util.SendJSONResponse(w, http.StatusBadRequest, util.ResponseMessages[400])
		return
	}

	newCampaign.CustomerID = userId
	savedCampaign, err := h.CampaignService.CreateNewCampaign(newCampaign)
	if err != nil {
		log.Printf("Error during creating new campaign: %e", err)
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[500])
		return
	}

	util.SendJSONResponse(w, http.StatusOK, savedCampaign)
}

func (h *CampaignHandler) GetSingleCampaignHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[600])
		return
	}

	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		log.Printf("Error during getting user id from token: %e", err)
		responseError := map[string]string{
			"error": err.Error(),
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	receivedCampaign, err := h.CampaignService.GetSingleCampaign(id, userId)
	if err != nil {
		log.Printf("Error during getting single campaign: %e", err)
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[500])
		return
	}

	util.SendJSONResponse(w, http.StatusOK, receivedCampaign)
}

func (h *CampaignHandler) UpdateCampaignHandler(w http.ResponseWriter, r *http.Request) {
	var campaignUpdate campaign.Campaign

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[600])
		return
	}

	err = json.NewDecoder(r.Body).Decode(&campaignUpdate)
	if err != nil {
		log.Print(err)
		util.SendJSONResponse(w, http.StatusBadRequest, util.ResponseMessages[400])
		return
	}

	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		log.Printf("Error during getting user id from token: %e", err)
		responseError := map[string]string{
			"error": err.Error(),
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, responseError)
		return
	}

	err = h.CampaignService.UpdateCampaign(campaignUpdate, id, userId)
	if err != nil {
		log.Printf("Error during updating campaign %d: %e", id, err)
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[500])
		return
	}

	util.SendJSONResponse(w, http.StatusOK, util.ResponseMessages[200])
}

func (h *CampaignHandler) DeleteCampaignHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[600])
		return
	}

	err = h.CampaignService.DeleteCampaign(id)
	if err != nil {
		log.Printf("Error during deleting campaign %d: %e", id, err)
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[500])
		return
	}

	util.SendJSONResponse(w, http.StatusOK, util.ResponseMessages[200])
}

func (h *CampaignHandler) GetUserCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[600])
		return
	}

	campaigns, err := h.CampaignService.FetchAllCampaigns(uint(id))
	if err != nil {
		log.Printf("Error during fetching user's campaign %d: %e", id, err)
		util.SendJSONResponse(w, http.StatusInternalServerError, util.ResponseMessages[500])
		return
	}

	util.SendJSONResponse(w, http.StatusOK, campaigns)
}

func (h *CampaignHandler) CampaignInitHandler(w http.ResponseWriter, r *http.Request) {
	customerKey, exists := mux.Vars(r)["customerKey"]
	if !exists {
		util.SendJSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": util.ResponseMessages[400],
		})
		return
	}

	fetchedCampaign, err := h.CampaignService.InitCampaign(customerKey)
	if err != nil {
		log.Print("Error during campaign initialization: ", err)
		if err.Error() == "record not found" {
			util.SendJSONResponse(w, http.StatusNotFound, map[string]string{
				"error": util.ResponseMessages[404],
			})
			return
		}
		util.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": util.ResponseMessages[500],
		})
		return
	}

	util.SendJSONResponse(w, http.StatusOK, fetchedCampaign)
}
