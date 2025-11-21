package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RealImage/team-48gb/internal/dtos"
	"github.com/RealImage/team-48gb/internal/services"
)

type CampaignHandler struct {
	campaignService *services.CampaignService
}

func NewCampaignHandler(campaignService *services.CampaignService) *CampaignHandler {
	return &CampaignHandler{
		campaignService: campaignService,
	}
}

func (h *CampaignHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateCampaignRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	campaign, err := h.campaignService.CreateCampaign(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(campaign)
}
