# Sync with upstream

This setting configures whether to pull in updates from the `upstream` remote.
This is intended for codebases that are forks of other codebases and want to
stay in sync with the codebase they are forked from.

## options

When set to `true` (the default value), `git town sync` also updates the local
[main-branch](main-branch.md) with changes from its counterpart in the
`upstream` remote. When set to `false`, `git town sync` does not pull in updates
from upstream even if that remote exists.

## in config file

In the [config file](../configuration-file.md) syncing with upstream can be set
like this:

```toml
[sync]
upstream = true
```

## in Git metadata

To manually configure syncing with upstream in Git, run this command:

```wrap
git config [--global] git-town.sync-upstream <true|false>
```

The optional `--global` flag applies this setting to all Git repositories on
your machine. Without it, the setting applies only to the current repository.
