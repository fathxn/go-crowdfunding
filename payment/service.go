package payment

import (
	"go-crowdfunding/transaction"
	"go-crowdfunding/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetToken(transaction transaction.Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func GetToken(transaction transaction.Transaction, user user.User) (string, error) {
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
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}