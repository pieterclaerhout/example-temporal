package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/pieterclaerhout/example-temporal/environ"
	"github.com/pieterclaerhout/example-temporal/workflows/cron"
	"github.com/pieterclaerhout/example-temporal/workflows/greeting"
	"github.com/pieterclaerhout/example-temporal/workflows/transfer"
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
)

func main() {

	log.PrintColors = true

	c, err := environ.NewClient()
	log.CheckError(err)
	defer c.Close()

	if len(os.Args) != 2 {
		log.Fatal("No arg specified: transfer | greeting")
	}

	switch os.Args[1] {
	case "transfer":
		runTransfer(c)
	case "greeting":
		runGreeting(c)
	case "cron":
		runCron(c)
	default:
		log.Fatal("Unknown argument:", os.Args[1])
	}

}

func runTransfer(c client.Client) {

	options := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: transfer.TaskQueue,
	}

	transferDetails := transfer.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, transfer.TransferMoney, transferDetails)
	log.CheckError(err)

	log.InfoDump(transferDetails, we.GetID()+"|"+we.GetRunID())

}

func runGreeting(c client.Client) {

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: greeting.TaskQueue,
	}
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.GreetingWorkflow, name)
	log.CheckError(err)

	var greeting string
	err = we.Get(context.Background(), &greeting)
	log.CheckError(err)

	log.InfoDump(greeting, we.GetID()+"|"+we.GetRunID())

}

func runCron(c client.Client) {

	workflowID := "cron_" + uuid.New().String()
	workflowOptions := client.StartWorkflowOptions{
		ID:           workflowID,
		TaskQueue:    "cron",
		CronSchedule: "* * * * *",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, cron.SampleCronWorkflow)
	log.CheckError(err)

	log.Info("Started workflow", "WorkflowID", we.GetID()+"|"+we.GetRunID())

}
