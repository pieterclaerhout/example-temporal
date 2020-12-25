package main

import (
	"github.com/pieterclaerhout/example-temporal"
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewClient(client.Options{})
	log.CheckError(err)
	defer c.Close()

	w := worker.New(c, example.TransferMoneyTaskQueue, worker.Options{})
	w.RegisterWorkflow(example.TransferMoney)
	w.RegisterActivity(example.Withdraw)
	w.RegisterActivity(example.Deposit)

	err = w.Run(worker.InterruptCh())
	log.CheckError(err)

}
