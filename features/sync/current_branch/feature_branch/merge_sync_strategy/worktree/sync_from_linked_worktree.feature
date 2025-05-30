Feature: sync a branch whose parent is active in another worktree

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE    | PARENT | LOCATIONS     |
      | parent | feature | main   | local, origin |
      | child  | feature | parent | local, origin |
    And the commits
      | BRANCH | LOCATION | MESSAGE              |
      | main   | local    | local main commit    |
      |        | origin   | origin main commit   |
      | parent | local    | local parent commit  |
      |        | origin   | origin parent commit |
      | child  | local    | local child commit   |
      |        | origin   | origin child commit  |
    And the current branch is "parent"
    And branch "child" is active in another worktree
    When I run "git-town sync" in the other worktree

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | child  | git fetch --prune --tags                          |
      |        | git checkout main                                 |
      | main   | git -c rebase.updateRefs=false rebase origin/main |
      |        | git push                                          |
      |        | git checkout child                                |
      | child  | git merge --no-edit --ff origin/parent            |
      |        | git merge --no-edit --ff origin/child             |
      |        | git push                                          |
    And these commits exist now
      | BRANCH | LOCATION                | MESSAGE                                                 |
      | main   | local, origin, worktree | origin main commit                                      |
      |        |                         | local main commit                                       |
      | child  | origin, worktree        | local child commit                                      |
      |        |                         | Merge remote-tracking branch 'origin/parent' into child |
      |        |                         | origin child commit                                     |
      |        |                         | Merge remote-tracking branch 'origin/child' into child  |
      |        | worktree                | origin parent commit                                    |
      | parent | local                   | local parent commit                                     |
      |        | origin                  | origin parent commit                                    |
    When I run "git-town undo" in the other worktree
    Then Git Town runs the commands
      | BRANCH | COMMAND                                                                            |
      | child  | git reset --hard {{ sha 'local child commit' }}                                    |
      |        | git push --force-with-lease origin {{ sha-in-origin 'origin child commit' }}:child |
    And the current branch in the other worktree is still "child"
    And these commits exist now
      | BRANCH | LOCATION                | MESSAGE              |
      | main   | local, origin, worktree | origin main commit   |
      |        |                         | local main commit    |
      | child  | origin                  | origin child commit  |
      |        | worktree                | local child commit   |
      | parent | local                   | local parent commit  |
      |        | origin                  | origin parent commit |
