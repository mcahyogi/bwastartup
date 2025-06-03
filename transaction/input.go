package transaction

import "bwastartup/user"

type GetCampaignTransactionInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}
