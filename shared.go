package example

const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"
const GreetingTaskQueue = "GREETING_TASK_QUEUE"

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}
