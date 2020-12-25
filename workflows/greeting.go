package workflows

import (
	"time"

	"github.com/pieterclaerhout/example-temporal/activities"
	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, activities.ComposeGreeting, name).Get(ctx, &result)
	return result, err

}
