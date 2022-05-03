package handler

import (
	"kitabantu-api/helper"
	"kitabantu-api/transaction"
	"kitabantu-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService transaction.TransactionService
}

func NewTranscationHandler(transactionService transaction.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (t *TransactionHandler) GetTransactionByCampaign(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := t.transactionService.GetTransactionByCampaignId(input)

	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	response := helper.ApiResponse("Success to get campaign transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

	return
}
