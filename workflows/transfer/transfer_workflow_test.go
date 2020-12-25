package transfer_test

import (
	"testing"

	"github.com/pieterclaerhout/example-temporal/workflows/transfer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_WorkflowWithdraw(t *testing.T) {

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := transfer.TransferDetails{
		Amount:      1.00,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: "1234",
	}

	env.OnActivity(transfer.Withdraw, mock.Anything, testDetails).Return(nil)
	env.OnActivity(transfer.Deposit, mock.Anything, testDetails).Return(nil)
	env.ExecuteWorkflow(transfer.TransferMoney, testDetails)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

}
