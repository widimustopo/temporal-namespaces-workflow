package activities

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func (a Activities) Counter(ctx context.Context, req *models.Product) (interface{}, error) {
	fmt.Println("product : ", req)

	_, _, result, err := a.FindByID(req.ProductID, "product")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	result.CounterOrder = result.CounterOrder + 1

	errUp := a.DB.Model(&req).Where("product_id = ?", req.ProductID).Update("counter_order", result.CounterOrder).Error
	if errUp != nil {
		logrus.Error("error ", errUp)
		return nil, errUp
	}

	return fmt.Sprintf("Product's Orders For ID : %v is : %v ", req.ProductID, result.CounterOrder), nil
}
