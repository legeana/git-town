Feature: deleting a branch whose tracking branch is gone

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | current | feature | main   | local, origin |
      | other   | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION      | MESSAGE        |
      | current | local, origin | current commit |
      | other   | local, origin | other commit   |
    And the current branch is "current"
    And origin deletes the "current" branch
    And the current branch is "current" and the previous branch is "other"
    When I run "git-town delete"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                  |
      | current | git fetch --prune --tags |
      |         | git checkout other       |
      | other   | git branch -D current    |
    And no uncommitted files exist now
    And the branches are now
      | REPOSITORY    | BRANCHES    |
      | local, origin | main, other |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE      |
      | other  | local, origin | other commit |
    And this lineage exists now
      | BRANCH | PARENT |
      | other  | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                       |
      | other  | git branch current {{ sha 'current commit' }} |
      |        | git checkout current                          |
    And these commits exist now
      | BRANCH  | LOCATION      | MESSAGE        |
      | current | local         | current commit |
      | other   | local, origin | other commit   |
    And the initial branches and lineage exist now
