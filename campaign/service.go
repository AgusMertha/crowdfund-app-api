package campaign

type CampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
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
