package transaction

import "bwastartup/user"

type GetCampaignTransactionInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int    `json:"amount" binding:"required"`
	CampaignID string `json:"campaign_id" binding:"required"`
	User       user.User
}
