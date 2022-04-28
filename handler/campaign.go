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

func (h CampaignHandler) GetCampaignById(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Failed get campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignById(input)

	if err != nil {
		if err != nil {
			errorResponse := helper.ApiResponse("Failed get campaign", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, errorResponse)
			return
		}
	}

	if campaignDetail.Id == 0 {
		errorMessage := gin.H{"errors": "Campaign not found"}
		errorResponse := helper.ApiResponse("Failed get campaign", http.StatusNotFound, "error", errorMessage)
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}

	response := helper.ApiResponse("Success get campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
	return
}
