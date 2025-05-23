Feature: delete the current branch that has an uncommitted file

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
    And an uncommitted file
    When I run "git-town delete"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                                                   |
      | current | git fetch --prune --tags                                  |
      |         | git push origin :current                                  |
      |         | git add -A                                                |
      |         | git commit -m "Committing open changes on deleted branch" |
      |         | git checkout other                                        |
      | other   | git branch -D current                                     |
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
      | BRANCH  | COMMAND                                                                  |
      | other   | git push origin {{ sha 'current commit' }}:refs/heads/current            |
      |         | git branch current {{ sha 'Committing open changes on deleted branch' }} |
      |         | git checkout current                                                     |
      | current | git reset --soft HEAD~1                                                  |
    And the initial commits exist now
    And the initial branches and lineage exist now
    And the uncommitted file still exists
