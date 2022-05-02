package campaign

import "strings"

type CampaignFormatter struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"shor_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserId           int    `json:"user_id"`
	Slug             string `json:"slug"`
	Perks            string `json:"perks"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		Id:               campaign.Id,
		Title:            campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserId:           campaign.UserId,
		Perks:            campaign.Perks,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	Id               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserId           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.Id = campaign.Id
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageUrl = ""
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserId = campaign.UserId
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		imageFormatter := CampaignImageFormatter{}
		imageFormatter.ImageUrl = image.FileName
		imageFormatter.IsPrimary = image.IsPrimary

		images = append(images, imageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
