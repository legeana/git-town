Feature: make multiple other branches contribution branches

  Background:
    Given a Git repo with origin
    And the branches
      | NAME      | TYPE    | PARENT | LOCATIONS     |
      | feature-1 | feature | main   | local, origin |
      | feature-2 | feature | main   | local, origin |
      | feature-3 | feature | main   | local, origin |
    And an uncommitted file
    When I run "git-town contribute feature-1 feature-2 feature-3"

  Scenario: result
    Then Git Town runs no commands
    And Git Town prints:
      """
      branch "feature-1" is now a contribution branch
      """
    And branch "feature-1" now has type "contribution"
    And Git Town prints:
      """
      branch "feature-2" is now a contribution branch
      """
    And branch "feature-2" now has type "contribution"
    And Git Town prints:
      """
      branch "feature-3" is now a contribution branch
      """
    And branch "feature-3" now has type "contribution"
    And the current branch is still "main"
    And the uncommitted file still exists

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND       |
      | main   | git add -A    |
      |        | git stash     |
      |        | git stash pop |
    And there are now no contribution branches
    And the current branch is still "main"
    And the uncommitted file still exists
