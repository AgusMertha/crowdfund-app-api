package handler

import (
	"kitabantu-api/campaign"
	"kitabantu-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService campaign.CampaignService
}

func NewCampaignHandler(campaignService campaign.CampaignService) *CampaignHandler {
	return &CampaignHandler{campaignService}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userId)

	if err != nil {
		errorResponse := helper.ApiResponse("Failed get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
		c.JSON(http.StatusOK, response)
		return
}
