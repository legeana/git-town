package messages

const (
	UndoContinueGuidance               = "\n\nTo continue after having resolved conflicts, run \"git-town continue\".\nTo go back to where you started, run \"git-town undo\".\n"
	AliasedCommands                    = "Aliased commands: %s\n"
	ArgumentUnknown                    = "unknown argument: %q"
	BranchAlreadyExistsLocally         = "there is already a branch %q"
	BranchAlreadyExistsRemotely        = "there is already a branch %q at the \"origin\" remote"
	BranchAuthorMultiple               = "\nMultiple people authored the %q branch.\n\n"
	BranchCheckoutProblem              = "cannot check out branch %q: %w"
	BranchCurrentProblem               = "cannot determine current branch: %w"
	BranchDeleted                      = "deleted branch %q"
	BranchDeletedHasUnmergedChanges    = "Branch %q was deleted at the remote but the local branch contains unshipped changes.\nI am therefore not removing this branch. You can see the unshipped changes by running \"git town diff-parent\"."
	BranchDiffProblem                  = "cannot determine if branch %q has unmerged commits: %w"
	BranchDoesntContainCommit          = "branch %q does not contain commit %q. Found commits %s"
	BranchDoesntExist                  = "there is no branch %q"
	BranchHasWrongSHA                  = "cannot reset branch %q to %q because it received additional commits in the meantime. It should have SHA %q but has %q"
	BranchIsAlreadyContribution        = "branch %q is already a contribution branch"
	BranchIsAlreadyObserved            = "branch %q is already observed"
	BranchIsAlreadyParked              = "branch %q is already parked"
	BranchLocalSHAProblem              = "cannot determine SHA of local branch %q: %w"
	BranchLocalProblem                 = "cannot determine whether the local branch %q exists: %w"
	BranchParentChanged                = "branch %q is now a child of %q"
	BrowserOpen                        = "Please open in a browser: %s\n"
	CacheUnitialized                   = "using a cached value before initialization"
	CodeHosting                        = "Code hosting: %s\n"
	CommandsRun                        = "Ran %d shell commands."
	CommitMessageProblem               = "cannot determine last commit message: %w"
	CompletionTypeUnknown              = "unknown completion type: %q"
	ConfigFileCannotRead               = "cannot read the configuration file %q: %w"
	ConfigFileInvalidData              = "the configuration file %q does not contain TOML-formatted content: %w"
	ConfigMainbranchInConfigFile       = "please configure the main branch in the config file"
	ConfigNeeded                       = "Git Town needs to be configured\n\n"
	ConfigStorage                      = "Config storage: %s\n"
	ConfigSyncFeatureStrategyUnknown   = "unknown sync-feature strategy: %q"
	ConfigSyncPerennialStrategyUnknown = "unknown sync-perennial strategy: %q"
	ConfigRemoveError                  = "unexpected error while removing the 'git-town' section from the Git configuration: %w"
	ContinueMessage                    = `You can run "git town continue" to finish it.`
	ContinueSkipGuidance               = "To continue by skipping the current branch, run \"git-town skip\"."
	ContributeBranchIsNowContribution  = "branch %q is now a contribution branch\n"
	ContributionBranchCannotPark       = "cannot park contribution branches"
	ContributionBranchCannotPropose    = "cannot propose contribution branches"
	ContributionBranchCannotShip       = "cannot ship contribution branches"
	DiffConflictWithMain               = "conflicts between your uncommmitted changes and the main branch"
	DryRun                             = "In dry run mode. No commands will be run. When run in normal mode, the command output will appear beneath the command. Some commands will only be run if necessary. For example: 'git push' will run if and only if there are local commits not on origin."
	ValueInvalid                       = "invalid value for %s: %q. Please provide either \"yes\" or \"no\""
	ValueGlobalInvalid                 = "invalid value for global %s: %q. Please provide either \"true\" or \"false\""
	ConflictDetectionProblem           = "cannot determine conflicts: %w"
	ContinueNothingToDo                = "nothing to continue"
	ContinueUnresolvedConflicts        = "you must resolve the conflicts before continuing"
	ContinueUntrackedChanges           = "please stage or commit the untracked changes first"
	DialogUnexpectedResponse           = "unexpected response: %s"
	DiffParentNoFeatureBranch          = "you can only diff-parent feature branches"
	DiffProblem                        = "cannot list diff of %q and %q: %w"
	DirCurrentProblem                  = "cannot determine the current directory"
	FileContentInvalidJSON             = "cannot parse JSON content of file %q: %w"
	FileDeleteProblem                  = "cannot delete file %q: %w"
	FileReadProblem                    = "cannot read file %q: %w"
	FileStatProblem                    = "cannot check file %q: %w"
	FileWriteProblem                   = "cannot write file %q: %w"
	GiteaToken                         = "Gitea token: %s\n"
	GitHubToken                        = "GitHub token: %s\n"
	GitLabToken                        = "GitLab token: %s\n"
	GitOutputIrregular                 = `
ERROR: Encountered irregular Git output

PLEASE REPORT THE OUTPUT BELOW AT https://github.com/git-town/git-town/issues/new

roblematic line: %q

BEGIN OUTPUT FROM 'git branch -vva'
%s
END OUTPUT FROM 'git branch -vva'
`
	GitUserEmailMissing                   = `please set the Git user email by running: git config --global user.email "<your email>"`
	GitUserNameMissing                    = `please set the Git user name by running: git config --global user.name "<your name>"`
	GitVersionMajorNotNumber              = "cannot convert major version %q to int: %w"
	GitVersionMinorNotNumber              = "cannot convert minor version %q to int: %w"
	GitVersionProblem                     = "cannot determine Git version: %w"
	GitVersionUnexpectedOutput            = "'git version' returned unexpected output: %q.\nPlease open an issue and supply the output of running 'git version'"
	GitVersionTooLow                      = "this app requires Git 2.7.0 or higher"
	HackTooManyArguments                  = "please provide only one branch to create"
	HackBranchIsAlreadyFeature            = "branch %q is already a feature branch"
	HackBranchIsNowFeature                = "branch %q is now a feature branch\n"
	HackCannotFeatureMainBranch           = "cannot make the main branch a feature branch"
	HackCannotFeaturePerennialBranch      = "branch %q is a perennial branch and therefore be a feature branch"
	HostingBitBucketNotImplemented        = "shipping pull requests via the Bitbucket API is currently not supported. If you need this functionality, please vote for it by opening a ticket at https://github.com/git-town/git-town/issues"
	HostingGitlabMergingViaAPI            = "GitLab API: Merging MR !%d ... "
	HostingGitlabUpdateMRViaAPI           = "GitLab API: Updating target branch for MR !%d to %q ... "
	HostingGiteaNotImplemented            = "shipping pull requests via the Gitea API is currently not supported. If you need this functionality, please vote for it by opening a ticket at https://github.com/git-town/git-town/issues"
	HostingGiteaUpdatePRViaAPI            = "Gitea API: Updating base branch for PR #%d to #%s"
	HostingGithubMergingViaAPI            = "GitHub API: merging PR #%d ... "
	HostingGithubUpdatePRViaAPI           = "GitHub API: updating base branch for PR #%d ... "
	HostingPlatformUnknown                = "unknown hosting platform: %q"
	InputAddOrRemove                      = `invalid argument %q. Please provide either "add" or "remove"`
	InputYesOrNo                          = `invalid argument: %q. Please provide either "yes" or "no".\n`
	KillBranchOtherWorktree               = `branch %q is active in another worktree`
	KillOnlyFeatureBranches               = "you can only kill feature branches"
	MainBranch                            = "Main branch: %s\n"
	MainBranchCannotMakeContribution      = "cannot make the main branch a contribution branch"
	MainBranchCannotObserve               = "cannot observe the main branch"
	MainBranchCannotPark                  = "cannot park the main branch"
	MainBranchCannotPropose               = "cannot propose the main branch"
	MainBranchCannotShip                  = "cannot ship the main branch"
	ObservedBranchCannotPark              = "cannot park observed branches"
	ObservedBranchCannotPropose           = "cannot propose observed branches"
	ObservedBranchCannotShip              = "cannot ship observed branches"
	ObservedBranchIsNowObserved           = "branch %q is now an observed branch\n"
	OfflineNotAllowed                     = "this command requires an active internet connection"
	OpcodeUnknown                         = "unknown opcode: %q, run \"git town status reset\" to reset it"
	OpenChangesProblem                    = "cannot determine open changes: %w"
	OriginHostname                        = "Origin hostname: %s\n"
	ParentDialogSelected                  = "Selected parent branch for %q: %s\n"
	ParkedBranchIsNowParked               = "branch %q is now parked\n"
	PerennialBranchCannotMakeContribution = "cannot make perennial branches contribution branches"
	PerennialBranchCannotObserve          = "cannot observe perennial branches"
	PerennialBranchCannotPark             = "cannot park perennial branches"
	PerennialBranchCannotPropose          = "cannot propose perennial branches"
	PerennialBranchCannotShip             = "cannot ship perennial branches"
	PerennialBranches                     = "Perennial branches: %s\n"
	PerennialRegex                        = "Perennial regex: %s\n"
	PreviousCommandFinished               = "The previous Git Town command (%s) finished successfully.\n"
	PreviousCommandProblem                = "The last Git Town command (%s) hit a problem %v ago.\n"
	ProposalMultipleFound                 = "found %d proposals from branch %q to branch %q"
	ProposalNoNumberGiven                 = "no proposal number given"
	ProposalNotFoundForBranch             = "cannot determine proposal for branch %q: %w"
	ProposalTargetBranchUpdateProblem     = "cannot update the target branch of proposal %d via the API"
	ProposalURLProblem                    = "cannot determine proposal URL from %q to %q: %w"
	PullRequestDeprecation                = `DEPRECATION NOTICE

This command has been renamed to "git town propose"
nd will be removed in future versions of Git Town.`
	PushHook                       = "Push hook: %s\n"
	PushNewBranches                = "Push new branches: %s\n"
	RebaseProblem                  = "cannot determine rebase in progress: %w"
	RemoteExistsProblem            = "cannot determine if remote %q exists: %w"
	RemotesProblem                 = "cannot determine remotes: %w"
	RenameBranchNotInSync          = "%q is not in sync with its tracking branch, please sync the branches before renaming"
	RenameMainBranch               = "the main branch cannot be renamed"
	RenamePerennialBranchWarning   = "%q is a perennial branch. Renaming a perennial branch typically requires other updates. If you are sure you want to do this, use '--force'"
	RenameToSameName               = "cannot rename branch to current name"
	RepoOutside                    = "this is not a Git repository"
	RunAutoUndo                    = "%s\nAuto-undo... "
	RunCommandProblem              = "error running command %q: %w"
	RunstateDeleted                = "Runstate file deleted."
	RunstateDeleteProblem          = "cannot delete previous run state: %w"
	RunstateLoadProblem            = "cannot load previous run state: %w"
	RunstateSerializeProblem       = "cannot encode run-state: %w"
	RunstatePathProblem            = "cannot determine the runstate file path: %w"
	RunstateSaveProblem            = "cannot save run state: %w"
	SetParentNoFeatureBranch       = "the branch %q is not a feature branch. Only feature branches can have parent branches"
	SettingDeprecatedGlobalMessage = `
I found the deprecated global setting %q.
I am upgrading this setting to the new format %q.
`
	SettingGlobalCannotRemove     = "ERROR: cannot remove global Git setting %q: %v"
	SettingGlobalCannotWrite      = "ERROR: cannot write global Git setting %q: %v"
	SettingIgnoreInvalid          = "Notice: ignoring invalid dialog input setting %q\n"
	SettingLocalDeprecatedMessage = `
I found the deprecated local setting %q.
I am upgrading this setting to the new format %q.
`
	SettingLocalCannotRemove    = "ERROR: cannot remove local Git setting %q: %v"
	SettingLocalCannotWrite     = "ERROR: cannot write local Git setting %q: %v"
	ShipAbortedMergeError       = "aborted because commit exited with error"
	ShipBranchOtherWorktree     = "branch %q is active in another worktree"
	ShipBranchNothingToDo       = "the branch %q has no shippable changes"
	ShipChildBranch             = "shipping this branch would ship %s as well,\nplease ship %q first"
	ShipDeletesTrackingBranches = "Ship deletes tracking branches: %s\n"
	ShipOpenChanges             = "you have uncommitted changes. Did you mean to commit them before shipping?"
	ShippableChangesProblem     = "cannot determine whether branch %q has shippable changes: %w"
	SkipBranchHasConflicts      = "cannot skip branch that resulted in conflicts"
	SkipMessage                 = `You can run "git town skip" to skip the currently failing operation.`
	SkipNothingToDo             = "nothing to skip"
	SquashCannotReadFile        = "cannot read squash message file %q: %w"
	SquashCommitAuthorQuery     = "Please choose an author for the squash commit:"
	SquashCommitAuthorProblem   = "error getting squash commit author: %w"
	SquashCommitAuthorSelection = "Selected squash commit author: %s\n"
	SquashMessageProblem        = "cannot comment out the squash commit message: %w"
	StatusFileNotFound          = "No status file found for this repository."
	SyncBeforeShip              = "Sync before ship: %s\n"
	SyncFeatureBranches         = "Sync feature branches: %s\n"
	SyncPerennialBranches       = "Sync perennial branches: %s\n"
	SyncStatusNotRecognized     = "cannot determine the sync status for Git remote %q and branch name %q"
	SyncWithUpstream            = "Sync with upstream: %s\n"
	UndoCreateOpcodeProblem     = "cannot create undo operations for %q: %w"
	UndoMessage                 = `You can run "git town undo" to go back to where you started.`
	UndoNothingToDo             = "nothing to undo"
	UnfinishedCommandHandle     = "Handle unfinished command: %s\n"
	UnfinishedRunStateContinue  = "Continue the \"%s\" command after having resolved conflicts"
	UnfinishedRunStateDiscard   = "Discard the unfinished state and run the new command"
	UnfinishedRunStateQuit      = "Quit without running anything"
	UnfinishedRunStateSkip      = "Skip the current branch and continue the \"%s\" command on the next branch"
	UnfinishedRunStateUndo      = "Undo the previous \"%s\" command"
)
