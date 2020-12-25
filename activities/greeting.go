package activities

import (
	"fmt"
)

const GreetingTaskQueue = "GREETING_TASK_QUEUE"

func ComposeGreeting(name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
