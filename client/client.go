package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func RunPaymentFail(req *models.TemporalOrderRequest) error {
	client := resty.New()

	url := req.Task + req.Data.PaymentID + "/" + req.Data.MemberID
	_, err := client.R().EnableTrace().Patch(url)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
