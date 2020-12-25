package greeting_test

import (
	"testing"

	"github.com/pieterclaerhout/example-temporal/workflows/greeting"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_WorkflowGreeting(t *testing.T) {

	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	env.OnActivity(greeting.ComposeGreeting, mock.Anything).Return("Hello World!", nil)
	env.ExecuteWorkflow(greeting.GreetingWorkflow, "World")

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var greeting string
	require.NoError(t, env.GetWorkflowResult(&greeting))
	require.Equal(t, "Hello World!", greeting)

}
