package app_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.skymeyer.dev/app"
)

func backupInfo(tb testing.TB) func(tb testing.TB) {
	var info = struct {
		name, version, commit, date string
	}{
		name:    app.Name,
		version: app.Version,
		commit:  app.BuildCommit,
		date:    app.BuildDate,
	}

	return func(tb testing.TB) {
		app.Name = info.name
		app.Version = info.version
		app.BuildCommit = info.commit
		app.BuildDate = info.date
	}
}

func TestInfo(t *testing.T) {
	teardown := backupInfo(t)
	defer teardown(t)

	for _, d := range []struct {
		name, version, commit, date string
		expected                    string
	}{
		{
			"foo", "1.0.2", "d240853866f20", "2022-02-11",
			"foo 1.0.2 %s %s/%s d240853866f20 2022-02-11",
		},
		{
			"foo", "1.0.2", "", "",
			"foo 1.0.2 %s %s/%s",
		},
	} {
		app.Name = d.name
		app.Version = d.version
		app.BuildCommit = d.commit
		app.BuildDate = d.date

		assert.Equal(t, fmt.Sprintf(
			d.expected,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		), app.Info())
	}
}
