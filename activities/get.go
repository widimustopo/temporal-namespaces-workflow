package activities

import (
	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
)

func (a Activities) FindByID(id string, types string) (*models.Member, *models.Payment, *models.Product, error) {

	switch types {
	case "member":
		var member models.Member

		err := a.DB.First(&member, "member_id = ?", id).Error
		if err != nil {
			logrus.Fatal("err ", err)
			return nil, nil, nil, err
		}
		return &member, nil, nil, nil
	case "payment":
		var payment models.Payment

		err := a.DB.First(&payment, "member_id = ? OR payment_id = ?", id, id).Error
		if err != nil {
			logrus.Fatal("err ", err)
			return nil, nil, nil, err
		}
		return nil, &payment, nil, nil
	case "product":
		var product models.Product
		err := a.DB.First(&product, "product_id = ?", id).Error
		if err != nil {
			logrus.Fatal("err ", err)
			return nil, nil, nil, err
		}
		return nil, nil, &product, nil

	}

	return nil, nil, nil, nil
}
