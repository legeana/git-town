Feature: auto-push the new branch without running Git push hooks

  Background:
    Given a Git repo with origin
    And Git setting "git-town.push-new-branches" is "true"
    And Git setting "git-town.push-hook" is "true"
    And the commits
      | BRANCH | LOCATION | MESSAGE       |
      | main   | origin   | origin commit |
    And the current branch is "main"
    When I run "git-town hack new"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                 |
      | main   | git fetch --prune --tags                |
      |        | git rebase origin/main --no-update-refs |
      |        | git checkout -b new                     |
      | new    | git push -u origin new                  |
    And the current branch is now "new"
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE       |
      | main   | local, origin | origin commit |
    And this lineage exists now
      | BRANCH | PARENT |
      | new    | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                     |
      | new    | git checkout main                           |
      | main   | git reset --hard {{ sha 'initial commit' }} |
      |        | git branch -D new                           |
      |        | git push origin :new                        |
    And the current branch is now "main"
    And the initial commits exist now
    And no lineage exists now