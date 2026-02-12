package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	payload "github.com/marveldo/gogin/internal/application/payloads"
)

type Taskhandler struct{}

func (t *Taskhandler) ProcessTask(ctx context.Context, task *asynq.Task) error {

	switch task.Type() {
	case "add":
		var data payload.AdditionPayload
		if err := json.Unmarshal(task.Payload(), &data); err != nil {
			return err
		}
		fmt.Printf("Background Addition of payload figures is : %v \n", data.Num_1+data.Num_2)
	case "email":
		var data payload.EmailPayload
		if err := json.Unmarshal(task.Payload(), &data); err != nil {
			return err
		}
		err := SendEmail(&data)
		if err != nil {
			fmt.Printf("Error is : %v", err.Error())
			return err
		}
	default:
		return fmt.Errorf("unexpected task type %q \n", task.Type())
	}
	return nil

}
