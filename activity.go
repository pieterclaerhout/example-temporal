package example

import (
	"context"
	"fmt"

	"github.com/pieterclaerhout/go-log"
)

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

func ComposeGreeting(name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
