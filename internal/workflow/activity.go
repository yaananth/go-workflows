package workflow

import (
	"github.com/cschleiden/go-dt/internal/command"
	"github.com/cschleiden/go-dt/internal/converter"
	"github.com/cschleiden/go-dt/internal/sync"
	"github.com/pkg/errors"
)

type Activity interface{}

// ExecuteActivity schedules the given activity to be executed
func ExecuteActivity(ctx sync.Context, name string, args ...interface{}) (sync.Future, error) {
	wfState := getWfState(ctx)

	inputs, err := converter.ArgsToInputs(converter.DefaultConverter, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert activity input")
	}

	// TOOO: Validate arguments against activity registration
	eventID := wfState.eventID
	wfState.eventID++

	command := command.NewScheduleActivityTaskCommand(eventID, name, "", inputs)
	wfState.addCommand(command)

	f := sync.NewFuture()
	wfState.pendingFutures[eventID] = f

	return f, nil
}
