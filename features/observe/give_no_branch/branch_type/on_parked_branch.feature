Feature: observing the current parked branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE   | PARENT | LOCATIONS     |
      | parked | parked | main   | local, origin |
    And the current branch is "parked"
    When I run "git-town observe"

  Scenario: result
    Then Git Town runs no commands
    And Git Town prints:
      """
      branch "parked" is now an observed branch
      """
    And the current branch is still "parked"
    And branch "parked" now has type "observed"
    And there are now no parked branches

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And the current branch is still "parked"
    And branch "parked" now has type "parked"
    And there are now no observed branches