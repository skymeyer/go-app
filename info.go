package app

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	Name        = "app"   // Application name
	Version     = "0.1.0" // Semantic version
	BuildCommit = ""      // Optional build commit
	BuildDate   = ""      // Optional build date
)

// Info returns the application version/runtime string.
func Info() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s %s %s/%s %s %s",
		Name,
		Version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		BuildCommit,
		BuildDate,
	))
}
