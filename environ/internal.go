package environ

import (
	"os"
)

func getString(name string, defaultVal string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return defaultVal
}
