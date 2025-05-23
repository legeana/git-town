Feature: sync the current observed branch in a local repo

  Background:
    Given a local Git repo
    And the branches
      | NAME  | TYPE     | LOCATIONS |
      | other | observed | local     |
    And the commits
      | BRANCH | LOCATION | MESSAGE      | FILE NAME  |
      | main   | local    | main commit  | main_file  |
      | other  | local    | local commit | local_file |
    And the current branch is "other"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND |
    And all branches are now synchronized
    And the initial commits exist now
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND |
    And the initial commits exist now
    And the initial branches and lineage exist now
