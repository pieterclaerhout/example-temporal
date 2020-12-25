package example

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_WorkflowWithdraw(t *testing.T) {

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := TransferDetails{
		Amount:      1.00,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: "1234",
	}

	env.OnActivity(Withdraw, mock.Anything, testDetails).Return(nil)
	env.OnActivity(Deposit, mock.Anything, testDetails).Return(nil)
	env.ExecuteWorkflow(TransferMoney, testDetails)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

}

func Test_WorkflowGreeting(t *testing.T) {

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	env.OnActivity(ComposeGreeting, mock.Anything).Return("Hello World!", nil)
	env.ExecuteWorkflow(GreetingWorkflow, "World")

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var greeting string
	require.NoError(t, env.GetWorkflowResult(&greeting))
	require.Equal(t, "Hello World!", greeting)

}
