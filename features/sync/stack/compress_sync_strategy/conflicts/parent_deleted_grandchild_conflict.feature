Feature: syncing a grandchild branch with conflicts using the "compress" strategy while its parent was deleted remotely

  Background:
    Given a Git repo with origin
    And the branches
      | NAME       | TYPE    | PARENT | LOCATIONS     |
      | child      | feature | main   | local, origin |
      | grandchild | feature | child  | local, origin |
    And the commits
      | BRANCH     | LOCATION | MESSAGE                       | FILE NAME        | FILE CONTENT       |
      | main       | local    | conflicting main commit       | conflicting_file | main content       |
      | grandchild | local    | conflicting grandchild commit | conflicting_file | grandchild content |
    And the current branch is "child" and the previous branch is "grandchild"
    And Git setting "git-town.sync-feature-strategy" is "compress"
    And origin deletes the "child" branch
    And wait 1 second to ensure new Git timestamps
    When I run "git-town sync --all"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH     | COMMAND                                           |
      | child      | git fetch --prune --tags                          |
      |            | git checkout main                                 |
      | main       | git -c rebase.updateRefs=false rebase origin/main |
      |            | git push                                          |
      |            | git branch -D child                               |
      |            | git checkout grandchild                           |
      | grandchild | git merge --no-edit --ff main                     |
    And Git Town prints the error:
      """
      git merge conflict
      """
    And Git Town prints the error:
      """
      To continue after having resolved conflicts, run "git town continue".
      To go back to where you started, run "git town undo".
      To continue by skipping the current branch, run "git town skip".
      """
    And a merge is now in progress

  Scenario: skip the grandchild merge conflict and delete the grandchild branch
    When I run "git-town skip"
    Then Git Town runs the commands
      | BRANCH     | COMMAND           |
      | grandchild | git merge --abort |
      |            | git push --tags   |
    When I run "git-town delete"
    Then Git Town runs the commands
      | BRANCH     | COMMAND                     |
      | grandchild | git fetch --prune --tags    |
      |            | git push origin :grandchild |
      |            | git checkout main           |
      | main       | git branch -D grandchild    |
