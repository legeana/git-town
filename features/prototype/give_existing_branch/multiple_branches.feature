Feature: prototype multiple other branches

  Background:
    Given a Git repo with origin
    And the branches
      | NAME         | TYPE         | PARENT | LOCATIONS     |
      | feature      | feature      | main   | local, origin |
      | contribution | contribution |        | local, origin |
      | observed     | observed     |        | local, origin |
      | parked       | parked       | main   | local, origin |
    When I run "git-town prototype feature contribution observed parked"

  Scenario: result
    Then Git Town runs no commands
    And Git Town prints:
      """
      branch "feature" is now a prototype branch
      """
    And branch "feature" now has type "prototype"
    And Git Town prints:
      """
      branch "contribution" is now a prototype branch
      """
    And branch "contribution" now has type "prototype"
    And Git Town prints:
      """
      branch "observed" is now a prototype branch
      """
    And branch "observed" now has type "prototype"
    And Git Town prints:
      """
      branch "parked" is now a prototype branch
      """
    And branch "parked" now has type "prototype"
    And the current branch is still "main"

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs no commands
    And there are now no prototype branches
    And branch "contribution" now has type "contribution"
    And branch "observed" now has type "observed"
    And branch "parked" now has type "parked"
    And the current branch is still "main"
    And the initial branches and lineage exist now
