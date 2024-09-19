package messages

const (
	UndoContinueGuidance              = "\n\nTo continue after having resolved conflicts, run \"git town continue\".\nTo go back to where you started, run \"git town undo\".\n"
	AliasedCommands                   = "Aliased commands: %s\n"
	ArgumentUnknown                   = "unknown argument: %q"
	APIParentBranchLookupStart        = "Looking for parent of %s ... "
	APIProposalLookupStart            = "Looking for proposal online ... "
	APIProposalUpdateStart            = "updating proposal online ... "
	APIUpdateProposalBase             = "updating base branch of proposal %s to %s ... "
	APIUpdateProposalHead             = "updating head branch of proposal %s to %s ... "
	APIGitHubCannotUpdateHeadBranch   = "GitHub cannot update the head branch of pull requests.\nOnce Git Town deleted the old tracking branch,\nGitHub has therefore closed the existing pull request for this branch (#%d).\nYou have to create a new one.\n"
	APIGitLabCannotUpdateHeadBranch   = "GitLab cannot update the source branch of merge requests:\nhttps://gitlab.com/gitlab-org/gitlab-foss/-/issues/47020\nGitLab has therefore closed the existing merge request for this branch (!%d).\nYou have to create a new one."
	APIGiteaCannotUpdateHeadBranch    = "Gitea cannot update the head branch of pull requests.\nOnce Git Town deleted the old tracking branch,\nGitea has therefore closed the existing pull request for this branch (#%d).\nYou have to create a new one.\n"
	BranchAlreadyExistsLocally        = "there is already a branch %q"
	BranchAlreadyExistsRemotely       = "there is already a branch %q at the \"origin\" remote"
	BranchAuthorMultiple              = "\nMultiple people authored the %q branch.\n\n"
	BranchCheckoutProblem             = "cannot check out branch %q: %w"
	BranchCurrentProblem              = "cannot determine current branch: %w"
	BranchDeleted                     = "deleted branch %q"
	BranchDeletedHasUnmergedChanges   = "Branch %q was deleted at the remote but the local branch contains unshipped changes.\nI am therefore not removing this branch. You can see the unshipped changes by running \"git town diff-parent\"."
	BranchDiffProblem                 = "cannot determine if branch %q has unmerged commits: %w"
	BranchDoesntContainCommit         = "branch %q does not contain commit %q. Found commits %s"
	BranchDoesntExist                 = "there is no branch %q"
	BranchHasWrongSHA                 = "cannot reset branch %q to %q because it received additional commits in the meantime. It should have SHA %q but has %q"
	BranchInfoNotFound                = "cannot find branch info for %q"
	BranchIsAlreadyContribution       = "branch %q is already a contribution branch"
	BranchIsAlreadyObserved           = "branch %q is already observed"
	BranchIsAlreadyPrototype          = "branch %q is already a prototype branch"
	BranchIsAlreadyParked             = "branch %q is already parked"
	BranchLocalSHAProblem             = "cannot determine SHA of local branch %q: %w"
	BranchLocalProblem                = "cannot determine whether the local branch %q exists: %w"
	BranchParentChanged               = "branch %q is now a child of %q"
	BrowserOpen                       = "Please open in a browser: %s\n"
	CacheUnitialized                  = "using a cached value before initialization"
	CodeHosting                       = "Code hosting: %s\n"
	CommandsRun                       = "Ran %d shell commands."
	CommitMessageProblem              = "cannot determine last commit message: %w"
	CompressUnsynced                  = "please sync branch %q before compressing it"
	CompressIsPerennial               = "better not compress perennial branches"
	CompressAlreadyOneCommit          = "branch %q has already just one commit"
	CompressBranchNoParent            = "cannot compress branch %q because it has no parent"
	CompressContributionBranch        = "you are merely contributing to branch %q and should leave compressing it to the branch owner"
	CompressNoBranchInfo              = "no branch info for branch %q"
	CompressNoCommits                 = "branch %q has no commits"
	CompressObservedBranch            = "you are merely observing branch %q and should leave compressing it to the branch owner"
	CompressParkedBranch              = "branch %q and should not compress it"
	CompletionTypeUnknown             = "unknown completion type: %q"
	ConfigFileCannotRead              = "cannot read the configuration file %q: %w"
	ConfigFileInvalidContent          = "the configuration file %q does not contain TOML-formatted content: %w"
	ConfigLineageParentIsChild        = "removing lineage entry for %q because the parent is the child"
	ConfigLineageEmptyChild           = "removing empty lineage entry"
	ConfigMainbranchInConfigFile      = "please configure the main branch in the config file"
	ConfigNeeded                      = "Git Town needs to be configured\n\n"
	ConfigScopeUnhandled              = "unhandled config scope"
	ConfigStorage                     = "Config storage: %s\n"
	ConfigShipStrategyUnknown         = "unknown ship strategy: %q"
	ConfigSyncStrategyUnknown         = "unknown sync strategy: %q"
	ConfigRemoveError                 = "unexpected error while removing the 'git-town' section from the Git configuration: %w"
	ContinueMessage                   = `You can run "git town continue" to finish it.`
	ContinueSkipGuidance              = "To continue by skipping the current branch, run \"git town skip\"."
	ContributeBranchIsNowContribution = "branch %q is now a contribution branch\n"
	ContributeBranchIsLocal           = "branch %q is local only - branches you want to contribute to must have a remote branch because they are per definition other people's branches"
	ContributionBranchCannotPark      = "cannot park contribution branches"
	ContributionBranchCannotPropose   = "cannot propose contribution branches"
	ContributionBranchCannotShip      = "cannot ship contribution branches"
	CreatePrototypeBranches           = "Create prototype branches:"
	DefaultBranchType                 = "Default branch type: %s\n"
	DiffConflictWithMain              = "conflicts between your uncommmitted changes and the main branch"
	DryRun                            = "In dry run mode. No commands will be run. When run in normal mode, the command output will appear beneath the command. Some commands will only be run if necessary. For example: 'git push' will run if and only if there are local commits not on origin."
	ValueInvalid                      = "invalid value for %s: %q. Please provide either \"yes\" or \"no\""
	ValueGlobalInvalid                = "invalid value for global %s: %q. Please provide either \"true\" or \"false\""
	ConflictDetectionProblem          = "cannot determine conflicts: %w"
	ContinueNothingToDo               = "nothing to continue"
	ContinueUnresolvedConflicts       = "you must resolve the conflicts before continuing"
	ContinueUntrackedChanges          = "please stage or commit the untracked changes first"
	CurrentBranchCannotDetermine      = "cannot determine the current branch"
	DialogUnexpectedResponse          = "unexpected response: %s"
	DiffParentNoFeatureBranch         = "you can only diff-parent feature branches"
	DiffProblem                       = "cannot list diff of %q and %q: %w"
	DirCurrentProblem                 = "cannot determine the current directory"
	FeatureRegex                      = "Feature regex: %s\n"
	FileContentInvalidJSON            = "cannot parse JSON content of file %q: %w"
	FileDeleteProblem                 = "cannot delete file %q: %w"
	FileReadProblem                   = "cannot read file %q: %w"
	FileStatProblem                   = "cannot check file %q: %w"
	FileWriteProblem                  = "cannot write file %q: %w"
	GiteaToken                        = "Gitea token: %s\n"
	GitAnotherProcessIsRunningRetry   = "another git process seems to be running in this repository, retrying in 1 sec ..."
	GitHubEnterpriseInitializeError   = "cannot initialize GitHub Enterprise client: %s"
	GitHubToken                       = "GitHub token: %s\n"
	GitLabToken                       = "GitLab token: %s\n"
	GitOutputIrregular                = `
ERROR: Encountered irregular Git output

PLEASE REPORT THE OUTPUT BELOW AT https://github.com/git-town/git-town/issues/new

roblematic line: %q

BEGIN OUTPUT FROM 'git branch -vva'
%s
END OUTPUT FROM 'git branch -vva'
`
	GitUserEmailMissing                   = `please set the Git user email by running: git config --global user.email "<your email>"`
	GitUserNameMissing                    = `please set the Git user name by running: git config --global user.name "<your name>"`
	GitURLCannotParse                     = "cannot parse Git URL %q"
	GitVersionMajorNotNumber              = "cannot convert major version %q to int: %w"
	GitVersionMinorNotNumber              = "cannot convert minor version %q to int: %w"
	GitVersionProblem                     = "cannot determine Git version: %w"
	GitVersionUnexpectedOutput            = "'git version' returned unexpected output: %q.\nPlease open an issue and supply the output of running 'git version'"
	GitVersionTooLow                      = "this app requires Git 2.30 or higher"
	HackTooManyArguments                  = "please provide only one branch to create"
	HackBranchIsAlreadyFeature            = "branch %q is already a feature branch"
	HackBranchIsNowFeature                = "branch %q is now a feature branch\n"
	HackCannotFeatureMainBranch           = "you are trying to convert the main branch to a feature branch. That's not possible. If you want to create a feature branch, did you forget to add the branch name?"
	HackCannotFeaturePerennialBranch      = "branch %q is a perennial branch and therefore be a feature branch"
	HostingBitBucketNotImplemented        = "shipping pull requests via the Bitbucket API is currently not supported. If you need this functionality, please vote for it by opening a ticket at https://github.com/git-town/git-town/issues"
	HostingGitlabMergingViaAPI            = "Merging MR !%d ... "
	HostingGitlabUpdateMRViaAPI           = "Updating target branch for MR !%d to %q ... "
	HostingGiteaNotImplemented            = "shipping pull requests via the Gitea API is currently not supported. If you need this functionality, please vote for it by opening a ticket at https://github.com/git-town/git-town/issues"
	HostingGiteaUpdatePRViaAPI            = "Gitea API: Updating base branch for PR #%d to #%s"
	HostingGithubMergingViaAPI            = "GitHub API: merging PR %s ... "
	HostingPlatformUnknown                = "unknown hosting platform: %q"
	InputAddOrRemove                      = `invalid argument %q. Please provide either "add" or "remove"`
	InputYesOrNo                          = `invalid argument: %q. Please provide either "yes" or "no".\n`
	KillBranchOtherWorktree               = `branch %q is active in another worktree`
	KillCannotKillMainBranch              = "you cannot kill the main branch"
	KillCannotKillPerennialBranches       = "you cannot kill perennial branches"
	MainBranch                            = "Main branch: %s\n"
	MainBranchCannotMakeContribution      = "cannot make the main branch a contribution branch"
	MainBranchCannotObserve               = "cannot observe the main branch"
	MainBranchCannotPark                  = "cannot park the main branch"
	MainBranchCannotPropose               = "cannot propose the main branch"
	MainBranchCannotPrototype             = "cannot prototype the main branch"
	MainBranchCannotShip                  = "cannot ship the main branch"
	ObservedBranchCannotPark              = "cannot park observed branches"
	ObservedBranchCannotPropose           = "cannot propose observed branches"
	ObservedBranchCannotShip              = "cannot ship observed branches"
	ObserveBranchIsLocal                  = "branch %q is local only - branches you want to observe must have a remote branch because they are per definition other people's branches"
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
	PerennialBranchCannotPrototype        = "cannot prototype perennial branches"
	PerennialBranchCannotShip             = "cannot ship perennial branches"
	PerennialBranches                     = "Perennial branches: %s\n"
	PerennialBranchRemovedParentEntry     = "Removed parent entry for perennial branch %q\n"
	PerennialRegex                        = "Perennial regex: %s\n"
	PreviousCommandFinished               = "The previous Git Town command (%s) finished successfully.\n"
	PreviousCommandProblem                = "The last Git Town command (%s) hit a problem %v ago.\n"
	ProposalMultipleFromToFound           = "found %d proposals from branch %q to branch %q"
	ProposalMultipleFromFound             = "found %d proposals for branch %q"
	ProposalNoNumberGiven                 = "no proposal number given"
	ProposalNoParent                      = "branch %q has no parent and can therefore not be proposed"
	ProposalNotFoundForBranch             = "cannot determine proposal for branch %q: %w"
	ProposalTargetBranchUpdateProblem     = "cannot update the target branch of proposal %d via the API"
	ProposalURLProblem                    = "cannot determine proposal URL from %q to %q: %w"
	PrototypeBranchIsNowPrototype         = "branch %q is now a prototype branch\n"
	PrototypeRemoved                      = "branch %q is no longer a prototype branch"
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
	SettingDeprecatedGlobalMessage = "Upgrading deprecated global setting %q to %q."
	SettingGlobalCannotRemove      = "ERROR: cannot remove global Git setting %q: %v"
	SettingGlobalCannotWrite       = "ERROR: cannot write global Git setting %q: %v"
	SettingIgnoreInvalid           = "Notice: ignoring invalid dialog input setting %q\n"
	SettingLocalDeprecatedMessage  = "Upgrading deprecated local setting %q to %q."
	SettingLocalCannotRemove       = "ERROR: cannot remove local Git setting %q: %v"
	SettingLocalCannotWrite        = "ERROR: cannot write local Git setting %q: %v"
	SettingSunsetDeleted           = "Deleting obsolete setting %q"
	ShipBranchDeletedAtRemote      = "branch %q was deleted at the remote"
	ShipBranchIsInOtherWorktree    = "branch %q is checked out in another worktree, please ship from there"
	ShipBranchNotInSync            = "branch %q is not in sync"
	ShipAbortedMergeError          = "aborted because merge exited with error"
	ShipAPIConnectorRequired       = "shipping via the API requires a connector"
	ShipBranchOtherWorktree        = "branch %q is active in another worktree"
	ShipBranchHasNoParent          = "branch %q has no parent to ship into"
	ShipBranchNothingToDo          = "the branch %q has no shippable changes"
	ShipChildBranch                = "shipping this branch would ship %s as well,\nplease ship %q first"
	ShipDeletesTrackingBranches    = "Ship deletes tracking branches: %s\n"
	ShipAPINoProposal              = "cannot ship branch %q via API because it has no proposal"
	ShipAPINoRemoteBranch          = "cannot ship branch %q via API because it has no remote branch"
	ShipMessageWithFastForward     = "shipping with the fast-forward strategy does not use the given commit message"
	ShipOpenChanges                = "you have uncommitted changes. Did you mean to commit them before shipping?"
	ShippableChangesProblem        = "cannot determine whether branch %q has shippable changes: %w"
	SkipBranchHasConflicts         = "cannot skip branch that resulted in conflicts"
	SkipMessage                    = `You can run "git town skip" to skip the currently failing operation.`
	SkipNothingToDo                = "nothing to skip"
	SkipNoInitialBranchInfo        = "found no information about branch %q in the initial snapshot"
	SkipNoFinalBranchInfo          = "found no information about branch %q in the final snapshot"
	SkipNoFinalSnapshot            = "found no final snapshot"
	SquashCannotReadFile           = "cannot read squash message file %q: %w"
	SquashCommitAuthorQuery        = "Please choose an author for the squash commit:"
	SquashCommitAuthorProblem      = "error getting squash commit author: %w"
	SquashCommitAuthorSelection    = "Selected squash commit author: %s\n"
	SquashMessageProblem           = "cannot comment out the squash commit message: %w"
	StatusFileNotFound             = "No status file found for this repository."
	SwitchNoBranches               = "no branches to switch to"
	SwitchUncommittedChanges       = "uncommitted changes"
	SyncFeatureBranches            = "Sync feature branches: %s\n"
	SyncPerennialBranches          = "Sync perennial branches: %s\n"
	SyncStatusNotRecognized        = "cannot determine the sync status for Git remote %q and branch name %q"
	SyncTags                       = "Sync tags: %s\n"
	SyncWithUpstream               = "Sync with upstream: %s\n"
	UndoCreateOpcodeProblem        = "cannot create undo operations for %q: %w"
	UndoMessage                    = `You can run "git town undo" to go back to where you started.`
	UndoNothingToDo                = "nothing to undo"
	UnfinishedCommandHandle        = "Handle unfinished command: %s\n"
	UnfinishedRunStateContinue     = "Continue the \"%s\" command after having resolved conflicts"
	UnfinishedRunStateDiscard      = "Discard the unfinished state and run the new command"
	UnfinishedRunStateQuit         = "Quit without running anything"
	UnfinishedRunStateSkip         = "Skip the current branch and continue the \"%s\" command on the next branch"
	UnfinishedRunStateUndo         = "Undo the previous \"%s\" command"
)
