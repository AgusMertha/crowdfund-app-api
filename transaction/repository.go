package transaction

import "gorm.io/gorm"

type TransactionRepository interface {
	GetTransactionByCampaignId(campaignId int) ([]Transaction, error)
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db}
}

func (t *TransactionRepositoryImpl) GetTransactionByCampaignId(campaignId int) ([]Transaction, error) {
	var transactions []Transaction
	err := t.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
