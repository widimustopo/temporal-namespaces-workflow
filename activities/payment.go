package activities

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func (a Activities) Payment(ctx context.Context, req *models.TemporalPaymentRequest) (interface{}, error) {
	fmt.Println("ini data payment : ", req)

	errSignal := a.SignalWorkflow(context.Background(), req.Data.PaymentID, "", "paymentsignal_"+req.Data.PaymentID, true)
	if errSignal != nil {
		logrus.Error(errSignal)
		return "Expired Paid Request", nil
	}

	var payments *models.Payment

	newUuid, _ := uuid.FromString(req.Data.PaymentID)
	payments = &models.Payment{
		PaymentID:     newUuid,
		StatusPayment: req.Data.StatusPayment,
	}

	err := a.DB.Model(&payments).Where("payment_id = ?", payments.PaymentID.String()).Update("status_payment", payments.StatusPayment).Error
	if err != nil {
		logrus.Error("error ", err)
		return nil, err
	}

	_, result, _, err := a.FindByID(req.Data.PaymentID, "payment")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return result, nil
}

func (a Activities) PaymentFail(ctx context.Context, req *models.TemporalPaymentRequest) (interface{}, error) {
	fmt.Println("ini data payment fail : ", req)

	var payments *models.Payment

	newUuid, _ := uuid.FromString(req.Data.PaymentID)
	payments = &models.Payment{
		PaymentID:     newUuid,
		StatusPayment: req.Data.StatusPayment,
	}

	err := a.DB.Model(&payments).Where("payment_id = ?", payments.PaymentID.String()).Update("status_payment", payments.StatusPayment).Error
	if err != nil {
		logrus.Error("error ", err)
		return nil, err
	}

	logrus.Info(req.Data.PaymentID, " : Payment has been Failed")

	return fmt.Sprintf("Payment has been Failed for ID : %v", req.Data.PaymentID), nil
}
