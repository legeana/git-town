Feature: display the parent of a top-level feature branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
    When I run "git-town config get-parent --verbose"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | TYPE    | COMMAND                       |
      |        | backend | git version                   |
      |        | backend | git rev-parse --show-toplevel |
      |        | backend | git config -lz --global       |
      |        | backend | git config -lz --local        |
      |        | backend | git config -lz                |
      |        | backend | git branch --show-current     |
    And Git Town prints:
      """
      Ran 6 shell commands.
      """
