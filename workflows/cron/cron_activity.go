package cron

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

type SampleCronResult struct {
	EndTime time.Time
}

func SampleCronActivity(ctx context.Context, beginTime, endTime time.Time) error {
	log := activity.GetLogger(ctx)
	log.Info("Cron job running.", "beginTime_exclude", beginTime, "endTime_include", endTime)
	return nil
}
