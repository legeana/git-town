Feature: sync all feature branches

  Background:
    Given my repo has the feature branches "alpha" and "beta"
    And my repo has the perennial branches "production" and "qa"
    And my repo contains the commits
      | BRANCH     | LOCATION      | MESSAGE                  |
      | main       | remote        | main commit              |
      | alpha      | local, remote | alpha commit             |
      | beta       | local, remote | beta commit              |
      | production | local         | local production commit  |
      |            | remote        | remote production commit |
      | qa         | local         | qa local commit          |
      |            | remote        | qa remote commit         |
    And I am on the "alpha" branch
    When I run "git-town sync --all"

  Scenario: result
    Then it runs the commands
      | BRANCH     | COMMAND                          |
      | alpha      | git fetch --prune --tags         |
      |            | git checkout main                |
      | main       | git rebase origin/main           |
      |            | git checkout alpha               |
      | alpha      | git merge --no-edit origin/alpha |
      |            | git merge --no-edit main         |
      |            | git push                         |
      |            | git checkout beta                |
      | beta       | git merge --no-edit origin/beta  |
      |            | git merge --no-edit main         |
      |            | git push                         |
      |            | git checkout production          |
      | production | git rebase origin/production     |
      |            | git push                         |
      |            | git checkout qa                  |
      | qa         | git rebase origin/qa             |
      |            | git push                         |
      |            | git checkout alpha               |
      | alpha      | git push --tags                  |
    And I am still on the "alpha" branch
    And all branches are now synchronized