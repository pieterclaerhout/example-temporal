package environ

import (
	"github.com/pieterclaerhout/go-log"
	"go.temporal.io/sdk/client"
)

const hostPortName = "HOST_PORT"
const hostPortDefault = "localhost:7233"

// NewClient returns a new client based on the environment variables
func NewClient() (client.Client, error) {

	logger := NewLogger()

	hostPort := getString(hostPortName, hostPortDefault)
	log.Debug("Client:", hostPort)

	return client.NewClient(client.Options{
		HostPort: hostPort,
		Logger:   logger,
	})

}
