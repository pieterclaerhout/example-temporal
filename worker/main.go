package main

import (
	"os"

	"github.com/pieterclaerhout/example-temporal/environ"
	"github.com/pieterclaerhout/example-temporal/workflows/cron"
	"github.com/pieterclaerhout/example-temporal/workflows/greeting"
	"github.com/pieterclaerhout/example-temporal/workflows/transfer"
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	log.PrintColors = true

	c, err := environ.NewClient()
	log.CheckError(err)
	defer c.Close()

	if len(os.Args) != 2 {
		log.Fatal("No arg specified: transfer | greeting")
	}

	var w worker.Worker

	switch os.Args[1] {
	case "transfer":
		w = workerTransfer(c)
	case "greeting":
		w = workerGreeting(c)
	case "cron":
		w = workerCron(c)
	default:
		log.Fatal("Unknown argument:", os.Args[1])
	}

	err = w.Run(worker.InterruptCh())
	log.CheckError(err)

}

func workerTransfer(c client.Client) worker.Worker {
	w := worker.New(c, transfer.TaskQueue, worker.Options{})
	w.RegisterWorkflow(transfer.TransferMoney)
	w.RegisterActivity(transfer.Withdraw)
	w.RegisterActivity(transfer.Deposit)
	return w
}

func workerGreeting(c client.Client) worker.Worker {
	w := worker.New(c, greeting.TaskQueue, worker.Options{})
	w.RegisterWorkflow(greeting.GreetingWorkflow)
	w.RegisterActivity(greeting.ComposeGreeting)
	return w
}

func workerCron(c client.Client) worker.Worker {
	w := worker.New(c, "cron", worker.Options{})
	w.RegisterWorkflow(cron.SampleCronWorkflow)
	w.RegisterActivity(cron.SampleCronActivity)
	return w
}
