package activities

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/client"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func (a Activities) Expired(ctx context.Context, req *models.TemporalOrderRequest, paid bool) error {
	fmt.Println("Will be Update Data to failed : ", req)

	if paid {
		logrus.Info(req.Data.PaymentID + " Has been Paid")
	} else {
		err := client.RunPaymentFail(req)
		if err != nil {
			logrus.Error(err.Error())
			return err
		}
		logrus.Info(req.Data.PaymentID + " failed to pay")
	}

	return nil
}
