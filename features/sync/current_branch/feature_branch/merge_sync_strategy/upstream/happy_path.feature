Feature: with upstream repo

  Background:
    Given a Git repo with origin
    And an upstream repo
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION | MESSAGE         |
      | main    | upstream | upstream commit |
      | feature | local    | local commit    |
    And the current branch is "feature"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                             |
      | feature | git fetch --prune --tags                            |
      |         | git checkout main                                   |
      | main    | git fetch upstream main                             |
      |         | git -c rebase.updateRefs=false rebase upstream/main |
      |         | git push                                            |
      |         | git checkout feature                                |
      | feature | git merge --no-edit --ff main                       |
      |         | git merge --no-edit --ff origin/feature             |
      |         | git push                                            |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH  | LOCATION                | MESSAGE                          |
      | main    | local, origin, upstream | upstream commit                  |
      | feature | local, origin           | local commit                     |
      |         |                         | Merge branch 'main' into feature |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                                               |
      | feature | git reset --hard {{ sha 'local commit' }}                             |
      |         | git push --force-with-lease origin {{ sha 'initial commit' }}:feature |
    And these commits exist now
      | BRANCH  | LOCATION                | MESSAGE         |
      | main    | local, origin, upstream | upstream commit |
      | feature | local                   | local commit    |
    And the initial branches and lineage exist now
