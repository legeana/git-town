Feature: cannot ship not-up-to-date feature branches using the fast-forward strategy

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION      | MESSAGE        |
      | feature | local, origin | feature commit |
      | main    | local, origin | main commit    |
    And the current branch is "feature"
    And Git setting "git-town.ship-strategy" is "fast-forward"
    When I run "git-town ship"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                     |
      | feature | git fetch --prune --tags    |
      |         | git checkout main           |
      | main    | git merge --ff-only feature |
      |         | git merge --abort           |
      |         | git checkout feature        |
    And Git Town prints the error:
      """
      aborted because merge exited with error
      """
    And the initial branches and lineage exist now
    And the initial commits exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the initial commits exist now
    And the initial branches and lineage exist now
