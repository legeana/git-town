package config

import (
	"fmt"

	"github.com/git-town/git-town/v11/src/cli/flags"
	"github.com/git-town/git-town/v11/src/cli/format"
	"github.com/git-town/git-town/v11/src/cli/io"
	"github.com/git-town/git-town/v11/src/cmd/cmdhelpers"
	"github.com/git-town/git-town/v11/src/execute"
	"github.com/git-town/git-town/v11/src/git"
	"github.com/git-town/git-town/v11/src/git/gitdomain"
	"github.com/git-town/git-town/v11/src/messages"
	"github.com/spf13/cobra"
)

const mainbranchDesc = "Displays or sets your main development branch"

const mainbranchHelp = `
The main branch is the Git branch from which new feature branches are cut.`

func mainbranchConfigCmd() *cobra.Command {
	addVerboseFlag, readVerboseFlag := flags.Verbose()
	cmd := cobra.Command{
		Use:   "main-branch [<branch>]",
		Args:  cobra.MaximumNArgs(1),
		Short: mainbranchDesc,
		Long:  cmdhelpers.Long(mainbranchDesc, mainbranchHelp),
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeConfigMainBranch(args, readVerboseFlag(cmd))
		},
	}
	addVerboseFlag(&cmd)
	return &cmd
}

func executeConfigMainBranch(args []string, verbose bool) error {
	repo, err := execute.OpenRepo(execute.OpenRepoArgs{
		Verbose:          verbose,
		DryRun:           false,
		OmitBranchNames:  true,
		PrintCommands:    true,
		ValidateIsOnline: false,
		ValidateGitRepo:  true,
	})
	if err != nil {
		return err
	}
	if len(args) > 0 {
		newMainBranch := gitdomain.NewLocalBranchName(args[0])
		return setMainBranch(newMainBranch, repo.Runner)
	}
	printMainBranch(repo.Runner)
	return nil
}

func printMainBranch(run *git.ProdRunner) {
	io.Println(format.StringSetting(run.Config.MainBranch.String()))
}

func setMainBranch(branch gitdomain.LocalBranchName, run *git.ProdRunner) error {
	if !run.Backend.HasLocalBranch(branch) {
		return fmt.Errorf(messages.BranchDoesntExist, branch)
	}
	return run.Config.SetMainBranch(branch)
}