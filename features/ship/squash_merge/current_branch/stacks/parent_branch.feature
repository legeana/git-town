Feature: ship a parent branch

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE    | PARENT | LOCATIONS     |
      | parent | feature | main   | local, origin |
      | child  | feature | parent | local, origin |
    And the commits
      | BRANCH | LOCATION      | MESSAGE       |
      | parent | local, origin | parent commit |
      | child  | local, origin | child commit  |
    And the current branch is "parent"
    And Git setting "git-town.ship-strategy" is "squash-merge"
    When I run "git-town ship -m 'parent done'"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                        |
      | parent | git fetch --prune --tags       |
      |        | git checkout main              |
      | main   | git merge --squash --ff parent |
      |        | git commit -m "parent done"    |
      |        | git push                       |
      |        | git push origin :parent        |
      |        | git branch -D parent           |
    And Git Town prints:
      """
      branch "child" is now a child of "main"
      """
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE      |
      | main   | local, origin | parent done  |
      | child  | local, origin | child commit |
    And this lineage exists now
      | BRANCH | PARENT |
      | child  | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                     |
      | main   | git revert {{ sha 'parent done' }}          |
      |        | git push                                    |
      |        | git branch parent {{ sha 'parent commit' }} |
      |        | git push -u origin parent                   |
      |        | git checkout parent                         |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE              |
      | main   | local, origin | parent done          |
      |        |               | Revert "parent done" |
      | child  | local, origin | child commit         |
      | parent | local, origin | parent commit        |
    And the initial branches and lineage exist now
