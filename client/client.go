package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func RunPaymentFail(req *models.TemporalOrderRequest) error {
	client := resty.New()

	url := req.Task + req.Data.PaymentID.String() + "/" + req.Data.MemberID
	fmt.Println(url, "ini url post")
	resp, err := client.R().EnableTrace().Patch(url)
	if err != nil {
		logrus.Error(err)
		return err
	}

	fmt.Println(resp.Body(), "data nya")
	return nil
}
