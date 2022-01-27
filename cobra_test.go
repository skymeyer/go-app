package app_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.skymeyer.dev/app"
)

func TestCobraVersion(t *testing.T) {
	teardown := backupInfo(t)
	defer teardown(t)

	app.Name = "foo"
	app.Version = "1.2.3"
	app.BuildCommit = "abcdef"
	app.BuildDate = "2022-08-05"

	cmd := app.NewVersionCmd()
	out := bytes.NewBufferString("")
	cmd.SetOut(out)

	// Long version
	cmd.Execute()
	res, err := ioutil.ReadAll(out)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf(
		"foo 1.2.3 %s %s/%s abcdef 2022-08-05\n",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	), string(res))

	// Short version
	cmd.SetArgs([]string{"--short"})
	cmd.Execute()
	res, err = ioutil.ReadAll(out)
	assert.NoError(t, err)
	assert.Equal(t, "1.2.3\n", string(res))
}
