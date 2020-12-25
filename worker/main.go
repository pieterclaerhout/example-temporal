package main

import (
	"os"

	"github.com/pieterclaerhout/example-temporal"
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewClient(client.Options{})
	log.CheckError(err)
	defer c.Close()

	log.InfoDump(os.Args, "os.Args")
	if len(os.Args) != 2 {
		log.Fatal("No arg specified: withdraw | greeting")
	}

	switch os.Args[1] {
	case "withdraw":
		runWorkerWithdraw(c)
	case "greeting":
		runWorkerGreeting(c)
	default:
		log.Fatal("Unknown argument:", os.Args[1])
	}

}

func runWorkerWithdraw(c client.Client) {

	w := worker.New(c, example.TransferMoneyTaskQueue, worker.Options{})

	w.RegisterWorkflow(example.TransferMoney)
	w.RegisterActivity(example.Withdraw)
	w.RegisterActivity(example.Deposit)

	err := w.Run(worker.InterruptCh())
	log.CheckError(err)

}

func runWorkerGreeting(c client.Client) {

	w := worker.New(c, example.GreetingTaskQueue, worker.Options{})

	w.RegisterWorkflow(example.GreetingWorkflow)
	w.RegisterActivity(example.ComposeGreeting)

	err := w.Run(worker.InterruptCh())
	log.CheckError(err)

}
