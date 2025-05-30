package flags_test

import (
	"testing"

	"github.com/git-town/git-town/v20/internal/cli/flags"
	"github.com/shoenig/test/must"
	"github.com/spf13/cobra"
)

func TestDryRun(t *testing.T) {
	t.Parallel()
	cmd := cobra.Command{}
	addFlag, readFlag := flags.DryRun()
	addFlag(&cmd)
	err := cmd.ParseFlags([]string{"--dry-run"})
	must.NoError(t, err)
	have, err := readFlag(&cmd)
	must.NoError(t, err)
	must.EqOp(t, true, have)
}
