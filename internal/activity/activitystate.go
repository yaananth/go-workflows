package activity

import (
	"context"

	"github.com/cschleiden/go-workflows/log"
	"github.com/cschleiden/go-workflows/workflow"
)

type ActivityState struct {
	ActivityID string
	Instance   *workflow.Instance
	Logger     log.Logger
}

func NewActivityState(activityID string, instance *workflow.Instance, logger log.Logger) *ActivityState {
	return &ActivityState{
		activityID,
		instance,
		logger.With(
			log.ActivityIDKey, activityID,
			log.InstanceIDKey, instance.InstanceID,
			log.ExecutionIDKey, instance.ExecutionID,
		)}
}

type key int

var activityCtxKey key

func WithActivityState(ctx context.Context, as *ActivityState) context.Context {
	return context.WithValue(ctx, activityCtxKey, as)
}

func GetActivityState(context context.Context) *ActivityState {
	return context.Value(activityCtxKey).(*ActivityState)
}
