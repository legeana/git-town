Feature: hack a new branch while the main branch is active in another worktree

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE    | PARENT | LOCATIONS     |
      | existing | feature | main   | local, origin |
    And the current branch is "existing"
    And the commits
      | BRANCH   | LOCATION | MESSAGE            |
      | main     | origin   | origin main commit |
      |          | local    | local main commit  |
      | existing | local    | existing commit    |
    And branch "main" is active in another worktree
    And an uncommitted file
    When I run "git-town hack new"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH   | COMMAND                     |
      | existing | git add -A                  |
      |          | git stash -m "Git Town WIP" |
      |          | git checkout -b new main    |
      | new      | git stash pop               |
    And the current branch is now "new"
    And the uncommitted file still exists
    And these commits exist now
      | BRANCH   | LOCATION        | MESSAGE            |
      | main     | origin          | origin main commit |
      |          | worktree        | local main commit  |
      | existing | local, worktree | existing commit    |
    And this lineage exists now
      | BRANCH   | PARENT |
      | existing | main   |
      | new      | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH   | COMMAND                     |
      | new      | git add -A                  |
      |          | git stash -m "Git Town WIP" |
      |          | git checkout existing       |
      | existing | git branch -D new           |
      |          | git stash pop               |
    And the current branch is now "existing"
    And these commits exist now
      | BRANCH   | LOCATION | MESSAGE            |
      | main     | origin   | origin main commit |
      |          | worktree | local main commit  |
      | existing | local    | existing commit    |
    And the initial branches and lineage exist now
    And the uncommitted file still exists
