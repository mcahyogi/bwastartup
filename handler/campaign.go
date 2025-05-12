package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID := c.Query("user_id")

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
