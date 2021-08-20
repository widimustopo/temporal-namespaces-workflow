package activities

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
	"go.temporal.io/sdk/client"

	"gorm.io/gorm"
)

type Activities struct {
	*gorm.DB
	client.Client
}

func HandlerActivities(db *gorm.DB, temporalClient client.Client) *Activities {
	return &Activities{
		db,
		temporalClient,
	}
}

func (a Activities) Register(ctx context.Context, req *models.TemporalMemberRequest) (interface{}, error) {

	fmt.Println("ini datanya kan : ", req)

	res := a.DB.Create(&req.Data)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil, res.Error
	}

	return &req.Data, nil
}
