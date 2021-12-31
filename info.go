package app

import (
	"fmt"
	"runtime"
)

var (
	Name    = "app"
	Version = ""
	Commit  = ""
)

func Info() string {
	return fmt.Sprintf("%s %s %s %s/%s",
		Name,
		VersionWithMeta(),
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

func VersionWithMeta() string {
	if Commit == "" {
		return Version
	}
	return fmt.Sprintf("%s+commit.%s", Version, Commit)
}
