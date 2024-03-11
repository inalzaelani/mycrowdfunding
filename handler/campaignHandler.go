package handler

import (
	"github.com/gin-gonic/gin"
	"latihanGo/campaign"
	"latihanGo/helper"
	"net/http"
	"strconv"
)

type campaignHanler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHanler {
	return &campaignHanler{service}
}

func (h *campaignHanler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse("Failed get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaign", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}
