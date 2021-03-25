package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse("error to get campaigns", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("List of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("failed to get campaign detail", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.ApiResponse("failed to get campaign detail", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CraeteCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("failed to create campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}
