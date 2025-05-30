Feature: active parked branches get synced like normal feature branches

  Background:
    Given a Git repo with origin
    And the branches
      | NAME   | TYPE   | PARENT | LOCATIONS     |
      | parked | parked | main   | local, origin |
    And the commits
      | BRANCH | LOCATION | MESSAGE              |
      | main   | local    | local main commit    |
      |        | origin   | origin main commit   |
      | parked | local    | local parked commit  |
      |        | origin   | origin parked commit |
    And the current branch is "parked"
    When I run "git-town sync"

  Scenario: result
    Then Git Town runs the commands
      | BRANCH | COMMAND                                           |
      | parked | git fetch --prune --tags                          |
      |        | git checkout main                                 |
      | main   | git -c rebase.updateRefs=false rebase origin/main |
      |        | git push                                          |
      |        | git checkout parked                               |
      | parked | git merge --no-edit --ff main                     |
      |        | git merge --no-edit --ff origin/parked            |
      |        | git push                                          |
    And all branches are now synchronized
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE                                                  |
      | main   | local, origin | origin main commit                                       |
      |        |               | local main commit                                        |
      | parked | local, origin | local parked commit                                      |
      |        |               | Merge branch 'main' into parked                          |
      |        |               | origin parked commit                                     |
      |        |               | Merge remote-tracking branch 'origin/parked' into parked |

  Scenario: undo
    When I run "git-town undo"
    Then Git Town runs the commands
      | BRANCH | COMMAND                                                                              |
      | parked | git reset --hard {{ sha 'local parked commit' }}                                     |
      |        | git push --force-with-lease origin {{ sha-in-origin 'origin parked commit' }}:parked |
    And these commits exist now
      | BRANCH | LOCATION      | MESSAGE              |
      | main   | local, origin | origin main commit   |
      |        |               | local main commit    |
      | parked | local         | local parked commit  |
      |        | origin        | origin parked commit |
    And the initial branches and lineage exist now
