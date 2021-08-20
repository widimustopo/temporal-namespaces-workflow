package activities

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"
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

	fmt.Println("Data Register : ", req)

	/*
		Raw Json : {
			"member_name" : "Widi Mustopo"
		}
	*/

	req.Data.MemberID = uuid.New()
	res := a.DB.Create(&req.Data)
	if res.Error != nil {
		logrus.Error(res.Error)
		return nil, res.Error
	}

	result, _, _, err := a.FindByID(req.Data.MemberID, "member")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return result, nil
}
