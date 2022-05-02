package handler

import (
	"fmt"
	"kitabantu-api/campaign"
	"kitabantu-api/helper"
	"kitabantu-api/user"
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

func (h *CampaignHandler) GetCampaignById(c *gin.Context) {
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

func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Create campaign failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)

	if err != nil {
		errorResponse := helper.ApiResponse("Create campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiResponse("Create campaign success", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UpdateCampaign(c *gin.Context) {
	var inputId campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputId)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Failed update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Failed update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.campaignService.UpdateCampaign(inputData, inputId)

	if err != nil {
		errorResponse := helper.ApiResponse("Update campaign failed", http.StatusBadRequest, "error", gin.H{"errors": err.Error()})
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiResponse("Update campaign success", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}
		errorResponse := helper.ApiResponse("Failed update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	file, err := c.FormFile("campaign_image")

	if err != nil {

		errorMessage := gin.H{"is_uploaded": false}
		errorResponse := helper.ApiResponse("Failed to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	path := fmt.Sprintf("images/campaign/%d-%s", input.CampaignId, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {

		errorMessage := gin.H{"is_uploaded": false}
		errorResponse := helper.ApiResponse("Failed to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	_, err = h.campaignService.UploadCampaignImage(input, path)

	if err != nil {

		errorMessage := gin.H{"is_uploaded": false}
		errorResponse := helper.ApiResponse("Failed to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorResponse := helper.ApiResponse("Success to upload campaign image", http.StatusOK, "success", gin.H{"is_uploaded": true})
	c.JSON(http.StatusOK, errorResponse)
}
