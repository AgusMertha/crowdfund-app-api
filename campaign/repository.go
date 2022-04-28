package campaign

import "gorm.io/gorm"

type CampaignRepository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(Id int) (Campaign, error)
}

type CampaignRepositoryImpl struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepositoryImpl {
	return &CampaignRepositoryImpl{db}
}

func (c *CampaignRepositoryImpl) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := c.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (c *CampaignRepositoryImpl) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	err := c.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (c *CampaignRepositoryImpl) FindById(Id int) (Campaign, error) {
	campaign := Campaign{}

	err := c.db.Preload("User").Preload("CampaignImages").Where("id = ?", Id).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
