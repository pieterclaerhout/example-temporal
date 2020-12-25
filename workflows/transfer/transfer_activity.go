package transfer

import (
	"context"

	"github.com/pieterclaerhout/go-log"
)

const TaskQueue = "TRANSFER_MONEY_TASK_QUEUE"

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}

func Withdraw(ctx context.Context, transferDetails TransferDetails) error {
	log.Infof(
		"\nWithdrawing $%f from account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ReferenceID,
	)
	return nil
}

func Deposit(ctx context.Context, transferDetails TransferDetails) error {
	log.Infof(
		"\nDepositing $%f into account %s. ReferenceId: %s\n",
		transferDetails.Amount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	return nil
}
