Feature: does not merge contribution branches

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE         | PARENT | LOCATIONS |
      | parent  | feature      | main   | local     |
      | current | contribution | parent | local     |
    And the current branch is "current"
    When I run "git-town merge"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH  | COMMAND                  |
      | current | git fetch --prune --tags |
    And Git Town prints the error:
      """
      cannot merge contribution branches
      """

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the initial commits exist now
    And the initial lineage exists now
