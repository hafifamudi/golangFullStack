package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var input campaign.CreateCampaignInput

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

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("failed to update campaign detail", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(input, inputData)
	if err != nil {
		response := helper.ApiResponse("failed to update campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("succesfully update campaign", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("failed to upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("Failed to upload campaign image!", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uplaoded": false}

		response := helper.ApiResponse("Failed to upload Avatar image!", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uplaoded": false}

		response := helper.ApiResponse("Failed to upload campaign image!", http.StatusUnprocessableEntity, "error", data)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uplaoded": true}

	response := helper.ApiResponse("campaign image successfully uploaded", http.StatusOK, "error", data)

	c.JSON(http.StatusOK, response)
}
