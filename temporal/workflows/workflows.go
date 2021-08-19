package workflows

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/widimustopo/temporal-namespaces-module/entities"
	"github.com/widimustopo/temporal-namespaces-module/libs"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func RegisterWorkflow(ctx workflow.Context, req *entities.TemporalRequest) (resp interface{}, err error) {
	if req == nil {
		logrus.Fatal("There is no request to proccess On Start Activity Workflow")
		return
	}

	ctx = withActivityOptions(ctx)
	err = workflow.ExecuteActivity(ctx, libs.ActivityRegisterMember, req).Get(ctx, &resp)

	return

}

func withActivityOptions(ctx workflow.Context) workflow.Context {
	ao := workflow.ActivityOptions{
		TaskQueue:              libs.RegisterWorkflow,
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
