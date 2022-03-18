package payment

import (
	"bwastartup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type service struct{}

func PaymentTransactionService() *service {
	return &service{}
}

// Get midtrans payment URL
func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
	midclient.ClientKey = ""
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	// create midtrands payment redirect URL
	return snapTokenResp.RedirectURL, nil
}
