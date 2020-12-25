package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/pieterclaerhout/example-temporal"
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
)

func main() {

	c, err := client.NewClient(client.Options{})
	log.CheckError(err)
	defer c.Close()

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
