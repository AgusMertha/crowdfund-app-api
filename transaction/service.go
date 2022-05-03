package transaction

import (
	"errors"
	"kitabantu-api/campaign"
)

type TransactionService interface {
	GetTransactionByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type TransactionServiceImpl struct {
	transactionRepository TransactionRepository
	campaignRepository    campaign.CampaignRepository
}

func NewTransactionService(transactionRepository TransactionRepository, campaignRepository campaign.CampaignRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{transactionRepository, campaignRepository}
}

func (t *TransactionServiceImpl) GetTransactionByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := t.campaignRepository.FindById(input.CampaignId)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.Id {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := t.transactionRepository.GetTransactionByCampaignId(input.CampaignId)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
