Feature: the branch to delete has a deleted tracking branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME  | TYPE    | PARENT | LOCATIONS     |
      | old   | feature | main   | local, origin |
      | other | feature | main   | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE      |
      | old    | local, origin | old commit   |
      | other  | local, origin | other commit |
    And the current branch is "old"
    And origin deletes the "old" branch
    And the current branch is "old" and the previous branch is "other"
    When I run "git-town delete"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                  |
      | old    | git fetch --prune --tags |
      |        | git checkout other       |
      | other  | git branch -D old        |
    And no uncommitted files exist now
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE      |
      | other  | local, origin | other commit |
    And the branches are now
      | REPOSITORY    | BRANCHES    |
      | local, origin | main, other |
    And this lineage exists now
      | BRANCH | PARENT |
      | other  | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                               |
      | other  | git branch old {{ sha 'old commit' }} |
      |        | git checkout old                      |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE      |
      | old    | local         | old commit   |
      | other  | local, origin | other commit |
    And the initial branches and lineage exist now
