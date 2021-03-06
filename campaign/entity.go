package campaign

import (
	"kitabantu-api/user"
	"time"
)

type Campaign struct {
	Id               int             `json:"id"`
	UserId           int             `json:"user_id"`
	Name             string          `json:"name"`
	ShortDescription string          `json:"short_description"`
	Description      string          `json:"description"`
	GoalAmount       int             `json:"goal_amount"`
	CurrentAmount    int             `json:"current_amount"`
	Perks            string          `json:"perks"`
	BackerCount      string          `json:"backer_count"`
	Slug             string          `json:"slug"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	CampaignImages   []CampaignImage `json:"campaign_images"`
	User             user.User       `json:"user"`
}

type CampaignImage struct {
	Id         int       `json:"id"`
	CampaignId int       `json:"campaign_id"`
	FileName   string    `json:"file_name"`
	IsPrimary  bool      `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
