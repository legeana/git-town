Feature: skip deleting the remote branch when shipping another branch using the always-merge strategy

  Background:
    Given a Git repo with origin
    And the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
      | other   | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION      | MESSAGE        |
      | feature | local, origin | feature commit |
      | other   | local         | other commit   |
    And the current branch is "other"
    And Git setting "git-town.ship-delete-tracking-branch" is "false"
    And Git setting "git-town.ship-strategy" is "always-merge"
    When I run "git-town ship feature" and close the editor
    And origin deletes the "feature" branch

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                             |
      | other  | git fetch --prune --tags            |
      |        | git checkout main                   |
      | main   | git merge --no-ff --edit -- feature |
      |        | git push                            |
      |        | git checkout other                  |
      | other  | git branch -D feature               |
    And the branches are now
      | REPOSITORY    | BRANCHES    |
      | local, origin | main, other |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                |
      | main   | local, origin | feature commit         |
      |        |               | Merge branch 'feature' |
      | other  | local         | other commit           |
    And this lineage exists now
      | BRANCH | PARENT |
      | other  | main   |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                       |
      | other  | git branch feature {{ sha 'feature commit' }} |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                |
      | main   | local, origin | feature commit         |
      |        |               | Merge branch 'feature' |
      | other  | local         | other commit           |
    And these branches exist now
      | REPOSITORY | BRANCHES             |
      | local      | main, feature, other |
      | origin     | main, other          |
    And the initial lineage exists now
