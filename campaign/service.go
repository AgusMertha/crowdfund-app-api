package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(input CreateCampaignInput, inputId GetCampaignDetailInput) (Campaign, error)
	UploadCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type CampaignServiceImpl struct {
	campaignRepository CampaignRepository
}

func NewCampaignService(campaignRepository CampaignRepository) *CampaignServiceImpl {
	return &CampaignServiceImpl{campaignRepository}
}

func (s *CampaignServiceImpl) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.campaignRepository.FindByUserId(userId)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.campaignRepository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *CampaignServiceImpl) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.campaignRepository.FindById(input.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignServiceImpl) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserId = input.User.Id

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.campaignRepository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *CampaignServiceImpl) UpdateCampaign(input CreateCampaignInput, inputId GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.campaignRepository.FindById(inputId.Id)

	if err != nil {
		return campaign, err
	}

	if campaign.Id == 0 {
		return campaign, errors.New("Campaign not found")
	}

	if campaign.UserId != input.User.Id {
		return campaign, errors.New("Not an owner of this campaign")
	}

	campaign.Name = input.Name
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)

	updatedCampaign, err := s.campaignRepository.Update(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *CampaignServiceImpl) UploadCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.campaignRepository.FindById(input.CampaignId)

	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.Id == 0 {
		return CampaignImage{}, errors.New("Campaign not found")
	}

	if campaign.UserId != input.User.Id {
		return CampaignImage{}, errors.New("Not an owner of this campaign")
	}

	if input.IsPrimary {
		_, err := s.campaignRepository.MarkAllIMageAsFalse(input.CampaignId)

		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignId = input.CampaignId
	campaignImage.IsPrimary = input.IsPrimary
	campaignImage.FileName = fileLocation

	image, err := s.campaignRepository.CreateImage(campaignImage)

	if err != nil {
		return image, err
	}

	return image, nil
}
