package payment

import (
	"kitabantu-api/transaction"
	"kitabantu-api/user"
	"os"

	"github.com/veritrans/go-midtrans"
)

type PaymentService interface {
	GetPaymentURL(transaction transaction.Transaction, user user.User) (string, error)
}

func NewPaymentService() *PaymentServiceImpl {
	return &PaymentServiceImpl{}
}

type PaymentServiceImpl struct {
}

func (p *PaymentServiceImpl) GetPaymentURL(transaction transaction.Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)

	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
