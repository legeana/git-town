Feature: sync the current prototype branch in a local repo

  Background:
    Given a local Git repo
    And the branches
      | NAME      | TYPE      | PARENT | LOCATIONS |
      | prototype | prototype | main   | local     |
    And the commits
      | BRANCH    | LOCATION | MESSAGE      | FILE NAME  |
      | main      | local    | main commit  | main_file  |
      | prototype | local    | local commit | local_file |
    And the current branch is "prototype"
    And Git setting "git-town.sync-feature-strategy" is "rebase"
    And an uncommitted file
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH    | COMMAND                          |
      | prototype | git add -A                       |
      |           | git stash -m "Git Town WIP"      |
      |           | git rebase main --no-update-refs |
      |           | git stash pop                    |
    And all branches are now synchronized
    And the current branch is still "prototype"
    And the uncommitted file still exists
    And these commits exist now
      | BRANCH    | LOCATION | MESSAGE      |
      | main      | local    | main commit  |
      | prototype | local    | local commit |
    And the initial branches and lineage exist now

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH    | COMMAND                                              |
      | prototype | git add -A                                           |
      |           | git stash -m "Git Town WIP"                          |
      |           | git reset --hard {{ sha-before-run 'local commit' }} |
      |           | git stash pop                                        |
    And the current branch is still "prototype"
    And the initial commits exist now
    And the initial branches and lineage exist now
