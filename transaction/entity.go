package transaction

import (
	"kitabantu-api/campaign"
	"kitabantu-api/user"
	"time"
)

type Transaction struct {
	Id         int       `json:"id"`
	CampaignId int       `json:"campaign_id"`
	UserId     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       user.User
	Campaign   campaign.Campaign
}
