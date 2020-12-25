package main

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/pieterclaerhout/example-temporal"
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
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
		runWithdraw(c)
	case "greeting":
		runGreeting(c)
	default:
		log.Fatal("Unknown argument:", os.Args[1])
	}

}

func runWithdraw(c client.Client) {

	options := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: example.TransferMoneyTaskQueue,
	}

	transferDetails := example.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, example.TransferMoney, transferDetails)
	log.CheckError(err)

	log.InfoDump(transferDetails, we.GetID()+"|"+we.GetRunID())

}

func runGreeting(c client.Client) {

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: example.GreetingTaskQueue,
	}
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, example.GreetingWorkflow, name)
	log.CheckError(err)

	var greeting string
	err = we.Get(context.Background(), &greeting)
	log.CheckError(err)

	log.InfoDump(greeting, we.GetID()+"|"+we.GetRunID())

}
