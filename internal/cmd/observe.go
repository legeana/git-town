package cmd

import (
	"errors"
	"fmt"

	"github.com/git-town/git-town/v17/internal/cli/flags"
	"github.com/git-town/git-town/v17/internal/cmd/cmdhelpers"
	"github.com/git-town/git-town/v17/internal/config/configdomain"
	"github.com/git-town/git-town/v17/internal/execute"
	"github.com/git-town/git-town/v17/internal/git/gitdomain"
	"github.com/git-town/git-town/v17/internal/messages"
	configInterpreter "github.com/git-town/git-town/v17/internal/vm/interpreter/config"
	. "github.com/git-town/git-town/v17/pkg/prelude"
	"github.com/spf13/cobra"
)

const observeDesc = "Stop your contributions to some feature branches"

const observeHelp = `
Marks the given local branches as observed.
If no branch is provided, observes the current branch.

Observed branches are useful when you assist other developers
and make local changes to try out ideas,
but want the other developers to implement and commit all official changes.

On an observed branch, "git town sync"
- pulls down updates from the tracking branch (always via rebase)
- does not push your local commits to the tracking branch
- does not pull updates from the parent branch
`

func observeCmd() *cobra.Command {
	addVerboseFlag, readVerboseFlag := flags.Verbose()
	cmd := cobra.Command{
		Use:     "observe [branches]",
		Args:    cobra.ArbitraryArgs,
		GroupID: "types",
		Short:   observeDesc,
		Long:    cmdhelpers.Long(observeDesc, observeHelp),
		RunE: func(cmd *cobra.Command, args []string) error {
			verbose, err := readVerboseFlag(cmd)
			if err != nil {
				return err
			}
			return executeObserve(args, verbose)
		},
	}
	addVerboseFlag(&cmd)
	return &cmd
}

func executeObserve(args []string, verbose configdomain.Verbose) error {
	repo, err := execute.OpenRepo(execute.OpenRepoArgs{
		DryRun:           false,
		PrintBranchNames: false,
		PrintCommands:    true,
		ValidateGitRepo:  true,
		ValidateIsOnline: false,
		Verbose:          verbose,
	})
	if err != nil {
		return err
	}
	data, err := determineObserveData(args, repo)
	if err != nil {
		return err
	}
	err = validateObserveData(data, repo)
	if err != nil {
		return err
	}
	branchNames := data.branchesToObserve.Keys()
	if err = repo.UnvalidatedConfig.NormalConfig.SetBranchTypeOverride(configdomain.BranchTypeObservedBranch, branchNames...); err != nil {
		return err
	}
	printObservedBranches(branchNames)
	if checkout, hasCheckout := data.checkout.Get(); hasCheckout {
		if err = repo.Git.CheckoutBranch(repo.Frontend, checkout, false); err != nil {
			return err
		}
	}
	return configInterpreter.Finished(configInterpreter.FinishedArgs{
		Backend:               repo.Backend,
		BeginBranchesSnapshot: Some(data.branchesSnapshot),
		BeginConfigSnapshot:   repo.ConfigSnapshot,
		Command:               "observe",
		CommandsCounter:       repo.CommandsCounter,
		FinalMessages:         repo.FinalMessages,
		Git:                   repo.Git,
		RootDir:               repo.RootDir,
		TouchedBranches:       branchNames.BranchNames(),
		Verbose:               verbose,
	})
}

type observeData struct {
	branchInfos       gitdomain.BranchInfos
	branchesSnapshot  gitdomain.BranchesSnapshot
	branchesToObserve configdomain.BranchesAndTypes
	checkout          Option[gitdomain.LocalBranchName]
}

func printObservedBranches(branches gitdomain.LocalBranchNames) {
	for _, branch := range branches {
		fmt.Printf(messages.ObservedBranchIsNowObserved, branch)
	}
}

func determineObserveData(args []string, repo execute.OpenRepoResult) (observeData, error) {
	branchesSnapshot, err := repo.Git.BranchesSnapshot(repo.Backend)
	if err != nil {
		return observeData{}, err
	}
	branchesToObserve, branchToCheckout, err := execute.BranchesToMark(args, branchesSnapshot, repo.UnvalidatedConfig)
	return observeData{
		branchInfos:       branchesSnapshot.Branches,
		branchesSnapshot:  branchesSnapshot,
		branchesToObserve: branchesToObserve,
		checkout:          branchToCheckout,
	}, err
}

func validateObserveData(data observeData, repo execute.OpenRepoResult) error {
	for branchName, branchType := range data.branchesToObserve {
		switch branchType {
		case configdomain.BranchTypeMainBranch:
			return errors.New(messages.MainBranchCannotObserve)
		case configdomain.BranchTypePerennialBranch:
			return errors.New(messages.PerennialBranchCannotObserve)
		case configdomain.BranchTypeObservedBranch:
			return fmt.Errorf(messages.BranchIsAlreadyObserved, branchName)
		case
			configdomain.BranchTypeFeatureBranch,
			configdomain.BranchTypeContributionBranch,
			configdomain.BranchTypeParkedBranch,
			configdomain.BranchTypePrototypeBranch:
		}
		hasLocalBranch := data.branchInfos.HasLocalBranch(branchName)
		hasRemoteBranch := data.branchInfos.HasMatchingTrackingBranchFor(branchName, repo.UnvalidatedConfig.NormalConfig.DevRemote)
		if !hasLocalBranch && !hasRemoteBranch {
			return fmt.Errorf(messages.BranchDoesntExist, branchName)
		}
		if hasLocalBranch && !hasRemoteBranch {
			return fmt.Errorf(messages.ObserveBranchIsLocal, branchName)
		}
	}
	return nil
}
