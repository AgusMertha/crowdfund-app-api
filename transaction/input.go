package transaction

import "kitabantu-api/user"

type GetCampaignTransactionsInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}
