package payment

import (
	"github.com/veritrans/go-midtrans"
	"latihanGo/campaign"
	"latihanGo/user"
	"strconv"
)

type service struct {
	campaignRepository campaign.Repository
}
type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService(campaignRepository campaign.Repository) *service {
	return &service{campaignRepository}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-yJVdabMak7uk89oi-yy55Z1B"
	midclient.ClientKey = "SB-Mid-client-afKI8K6N8Gpz15Pc"
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
