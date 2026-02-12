package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type Taskhandler struct{}

func (t *Taskhandler) ProcessTask(ctx context.Context, task *asynq.Task) error {

	switch task.Type() {
	case "add":
		var data AdditionPayload
		if err := json.Unmarshal(task.Payload(), &data); err != nil {
			return err
		}
		fmt.Printf("Background Addition of payload figures is : %v", data.Num_1+data.Num_2)
	default:
		return fmt.Errorf("unexpected task type %q", task.Type())
	}
	return nil

}
