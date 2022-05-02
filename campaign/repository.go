package campaign

import "gorm.io/gorm"

type CampaignRepository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(Id int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllIMageAsFalse(CampaignId int) (bool, error)
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

func (c *CampaignRepositoryImpl) Save(campaign Campaign) (Campaign, error) {
	err := c.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (c *CampaignRepositoryImpl) Update(campaign Campaign) (Campaign, error) {
	err := c.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (c *CampaignRepositoryImpl) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := c.db.Create(&campaignImage).Error

	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (c *CampaignRepositoryImpl) MarkAllIMageAsFalse(CampaignId int) (bool, error) {
	err := c.db.Model(&CampaignImage{}).Where("campaign_id = ?", CampaignId).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
