package workflows

import (
	"testing"

	"github.com/pieterclaerhout/example-temporal/activities"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_WorkflowWithdraw(t *testing.T) {

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := activities.TransferDetails{
		Amount:      1.00,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: "1234",
	}

	env.OnActivity(activities.Withdraw, mock.Anything, testDetails).Return(nil)
	env.OnActivity(activities.Deposit, mock.Anything, testDetails).Return(nil)
	env.ExecuteWorkflow(TransferMoney, testDetails)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

}
