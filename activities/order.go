package activities

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/libs"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
	"go.temporal.io/sdk/client"
)

func (a Activities) Order(ctx context.Context, req *models.TemporalOrderRequest) (interface{}, error) {
	fmt.Println("ini datanya kan : ", req)
	member, _, _, err := a.FindByID(req.Data.MemberID, "member")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	_, _, product, err := a.FindByID(req.Data.ProductID, "product")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	req.Data.MemberName = member.MemberName
	req.Data.ProductName = product.ProductName

	req.Data.Price = product.Price

	fullPrice := req.Data.Price * float64(req.Data.Qty)

	req.Data.FullPrice = fullPrice

	errInsert := a.DB.Create(req.Data).Error
	if errInsert != nil {
		logrus.Error(errInsert.Error())
		return nil, err
	}

	_, result, _, err := a.FindByID(req.Data.MemberID, "payment")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	req.Data.PaymentID = result.PaymentID
	ErrExpiredWorkflow := executeExpiredWorkflow(ctx, a.Client, req)
	if ErrExpiredWorkflow != nil {
		logrus.Error(ErrExpiredWorkflow.Error())
	}

	return result, nil
}

func executeExpiredWorkflow(ctx context.Context, temporalClient client.Client, req *models.TemporalOrderRequest) (err error) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                 req.Data.PaymentID.String(),
		TaskQueue:          libs.ExpiredWorkflow,
		WorkflowRunTimeout: time.Minute * 10,
	}

	workflowRun, err := temporalClient.ExecuteWorkflow(ctx, workflowOptions, "ExpiredWorkflow", req)
	if err != nil {
		return err
	}

	logrus.Println("Started workflow", "WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	return
}
