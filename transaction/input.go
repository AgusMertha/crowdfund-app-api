package transaction

import "kitabantu-api/user"

type GetCampaignTransactionsInput struct {
	CampaignId int `uri:"id" binding:"required"`
	User       user.User
}
