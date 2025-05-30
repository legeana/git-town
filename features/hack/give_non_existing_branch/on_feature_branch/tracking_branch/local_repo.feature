Feature: local repo

  Background:
    Given a local Git repo
    And the branches
      | NAME     | TYPE    | PARENT | LOCATIONS |
      | existing | feature | main   | local     |
    And the commits
      | BRANCH | LOCATION | MESSAGE     |
      | main   | local    | main commit |
    And the current branch is "existing"
    When I run "git-town hack new"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH   | COMMAND                  |
      | existing | git checkout -b new main |
    And the initial commits exist now
    And this lineage exists now
      | BRANCH   | PARENT |
      | existing | main   |
      | new      | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH   | COMMAND               |
      | new      | git checkout existing |
      | existing | git branch -D new     |
    And the initial commits exist now
    And this lineage exists now
      | BRANCH   | PARENT |
      | existing | main   |
