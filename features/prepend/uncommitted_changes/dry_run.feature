Feature: dry-run prepending a branch to a feature branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME | TYPE    | PARENT | LOCATIONS     |
      | old  | feature | main   | local, origin |
    And the current branch is "old"
    And the commits
      | BRANCH | LOCATION      | MESSAGE    |
      | old    | local, origin | old commit |
    And an uncommitted file
    When I run "git-town prepend parent --dry-run"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                     |
      | old    | git add -A                  |
      |        | git stash -m "Git Town WIP" |
      |        | git checkout -b parent main |
      | parent | git stash pop               |
    And the current branch is still "old"
    And the uncommitted file still exists
    And the initial commits exist now
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the current branch is still "old"
    And the uncommitted file still exists
    And the initial commits exist now
    And the initial lineage exists now
