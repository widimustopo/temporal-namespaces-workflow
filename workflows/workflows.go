package workflows

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-workflow/libs"
	"github.com/widimustopo/temporal-namespaces-workflow/models"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func RegisterWorkflow(ctx workflow.Context, req *models.TemporalMemberRequest) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On Register Activity Workflow")
		return
	}

	ctx = withActivityOptions(ctx, req.WorkflowName)
	err = workflow.ExecuteActivity(ctx, libs.ActivityRegisterMember, req).Get(ctx, &resp)

	return

}

func OrderWorkflow(ctx workflow.Context, req *models.TemporalOrderRequest) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On Order Activity Workflow")
		return
	}

	ctx = withActivityOptions(ctx, libs.OrderWorkflow)
	err = workflow.ExecuteActivity(ctx, libs.ActivityOrder, req).Get(ctx, &resp)

	return

}

func PaymentWorkflow(ctx workflow.Context, req *models.TemporalPaymentRequest) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On Payment Activity Workflow")
		return
	}

	ctx = withActivityOptions(ctx, libs.PaymentWorkflow)
	err = workflow.ExecuteActivity(ctx, libs.ActivityPayment, req).Get(ctx, &resp)

	return

}

func PaymentFailWorkflow(ctx workflow.Context, req *models.TemporalPaymentRequest) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On PaymentFail Activity Workflow")
		return
	}

	ctx = withActivityOptions(ctx, libs.PaymentFailWorkflow)
	err = workflow.ExecuteActivity(ctx, libs.ActivityPaymentFail, req).Get(ctx, &resp)

	return

}

func withActivityOptions(ctx workflow.Context, queueName string) workflow.Context {
	ao := workflow.ActivityOptions{
		TaskQueue:              queueName,
		ScheduleToStartTimeout: 24 * time.Hour,
		StartToCloseTimeout:    24 * time.Hour,
		HeartbeatTimeout:       time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        time.Minute * 5,
			NonRetryableErrorTypes: []string{"BusinessError"},
		},
	}
	ctxOut := workflow.WithActivityOptions(ctx, ao)
	return ctxOut
}

func ExpiredWorkflow(ctx workflow.Context, req *models.TemporalOrderRequest, queueName string) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On Expired Workflow")
		return
	}

	statusPayment := false
	//  cancel payment expired
	selector := workflow.NewSelector(ctx)
	selectorCh := workflow.GetSignalChannel(ctx, "paymentsignal_"+req.Data.PaymentID.String())
	selector.AddReceive(selectorCh, func(ch workflow.ReceiveChannel, _ bool) {
		var isPaidSubSignal bool
		ch.Receive(ctx, &isPaidSubSignal)
		statusPayment = isPaidSubSignal
	})

	fmt.Println(statusPayment)
	ctx = actvityOptionExpired(ctx, req)

	for {

		if time.Now().Local().After(req.Times) {
			fmt.Println("Habis waktu")
			err = workflow.ExecuteActivity(ctx, libs.ActivityExpired, req, statusPayment).Get(ctx, &resp)

			break
		}

		workflow.AwaitWithTimeout(ctx, time.Second*10, func() bool {
			return statusPayment
		})

		for selector.HasPending() {
			selector.Select(ctx)
		}

		if statusPayment {
			fmt.Println("Sudah di bayar")
			err = workflow.ExecuteActivity(ctx, libs.ActivityExpired, req, statusPayment).Get(ctx, &resp)
			break
		}

	}

	return
}

func actvityOptionExpired(ctx workflow.Context, req *models.TemporalOrderRequest) workflow.Context {
	ao := workflow.ActivityOptions{
		TaskQueue:              req.WorkflowName,
		ScheduleToStartTimeout: 24 * time.Hour,
		StartToCloseTimeout:    24 * time.Hour,
		HeartbeatTimeout:       time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        req.Interval,
			MaximumInterval:        req.MaxInterval,
			MaximumAttempts:        req.Attempt,
			NonRetryableErrorTypes: []string{"BusinessError"},
		},
	}
	ctxOut := workflow.WithActivityOptions(ctx, ao)
	return ctxOut
}
