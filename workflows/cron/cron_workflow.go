package cron

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SampleCronWorkflow(ctx workflow.Context) (*SampleCronResult, error) {

	log := workflow.GetLogger(ctx)

	log.Info("Cron workflow started.", "StartTime", workflow.Now(ctx))

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	startTime := time.Time{}
	if workflow.HasLastCompletionResult(ctx) {
		var lastResult SampleCronResult
		if err := workflow.GetLastCompletionResult(ctx, &lastResult); err == nil {
			startTime = lastResult.EndTime
		}
	}

	endTime := workflow.Now(ctx)

	err := workflow.ExecuteActivity(ctx1, SampleCronActivity, startTime, endTime).Get(ctx, nil)
	if err != nil {
		log.Error("Cron job failed.", "Error", err)
		return nil, err
	}

	return &SampleCronResult{EndTime: endTime}, nil

}
