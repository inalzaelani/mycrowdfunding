package campaign

import "latihanGo/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name        string `json:"name" binding:"required"`
	ShortDesc   string `json:"short_desc" binding:"required"`
	Description string `json:"description" binding:"required"`
	GoalAmount  int    `json:"goal_amount" binding:"required"`
	Perks       string `json:"perks" binding:"required"`
	User        user.User
}
