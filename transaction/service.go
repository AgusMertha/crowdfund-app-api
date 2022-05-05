package transaction

import (
	"errors"
	"kitabantu-api/campaign"
	"strconv"
	"time"
)

type TransactionService interface {
	GetTransactionByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
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

	if campaign.Id == 0 {
		return []Transaction{}, errors.New("Campaign not found")
	}

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

func (t *TransactionServiceImpl) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transactions, err := t.transactionRepository.GetTransactionByUser(userId)

	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (t *TransactionServiceImpl) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}

	transaction.Amount = input.Amount
	transaction.CampaignId = input.CampaignId
	transaction.UserId = input.User.Id
	transaction.Status = "pending"

	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	transaction.Code = "order-" + strconv.Itoa(input.CampaignId) + strconv.Itoa(input.User.Id) + timestamp

	newTransaction, err := t.transactionRepository.Save(transaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
