package activities

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func (a Activities) AddProduct(ctx context.Context, req *models.TemporalProductRequest) (interface{}, error) {
	fmt.Println("data request add product : ", req)

	/*
		Raw Json : {
			"product_name": "Bag",
			"price": "100"
		}
	*/

	req.Data.ProductID = uuid.New()
	res := a.DB.Create(&req.Data)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil, res.Error
	}

	_, _, result, err := a.FindByID(req.Data.ProductID, "product")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return result, nil
}
