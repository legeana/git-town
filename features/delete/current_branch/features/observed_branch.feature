Feature: delete the current observed branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME     | TYPE     | PARENT | LOCATIONS     |
      | observed | observed |        | local, origin |
      | feature  | feature  | main   | local, origin |
    And the commits
      | BRANCH   | LOCATION      | MESSAGE         |
      | feature  | local, origin | feature commit  |
      | observed | local, origin | observed commit |
    And the current branch is "observed"
    And an uncommitted file
    And the current branch is "observed" and the previous branch is "feature"
    When I run "git-town delete"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH   | COMMAND                                                   |
      | observed | git fetch --prune --tags                                  |
      |          | git add -A                                                |
      |          | git commit -m "Committing open changes on deleted branch" |
      |          | git checkout feature                                      |
      | feature  | git branch -D observed                                    |
    And the current branch is now "feature"
    And no uncommitted files exist now
    And the branches are now
      | REPOSITORY | BRANCHES                |
      | local      | main, feature           |
      | origin     | main, feature, observed |
    And these commits exist now
      | BRANCH   | LOCATION      | MESSAGE         |
      | feature  | local, origin | feature commit  |
      | observed | origin        | observed commit |
    And this lineage exists now
      | BRANCH  | PARENT |
      | feature | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH   | COMMAND                                                                   |
      | feature  | git branch observed {{ sha 'Committing open changes on deleted branch' }} |
      |          | git checkout observed                                                     |
      | observed | git reset --soft HEAD~1                                                   |
    And the current branch is now "observed"
    And the uncommitted file still exists
    And the initial commits exist now
    And the initial branches and lineage exist now
    And branch "observed" now has type "observed"
