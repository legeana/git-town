@messyoutput
Feature: remove an existing code hosting override

  Background:
    Given a Git repo clone
    And local Git Town setting "hosting-platform" is "github"
    When I run "git-town config setup" and enter into the dialog:
      | DIALOG                      | KEYS           | DESCRIPTION                                 |
      | welcome                     | enter          |                                             |
      | aliases                     | enter          |                                             |
      | main branch                 | down enter     |                                             |
      | perennial branches          |                | no input here since the dialog doesn't show |
      | perennial regex             | enter          |                                             |
      | hosting platform            | up up up enter |                                             |
      | origin hostname             | enter          |                                             |
      | sync-feature-strategy       | enter          |                                             |
      | sync-perennial-strategy     | enter          |                                             |
      | sync-upstream               | enter          |                                             |
      | push-new-branches           | enter          |                                             |
      | push-hook                   | enter          |                                             |
      | ship-delete-tracking-branch | enter          |                                             |
      | sync-before-ship            | enter          |                                             |
      | save config to Git metadata | down enter     |                                             |

  Scenario: result
    Then it runs the commands
      | COMMAND                                      |
      | git config --unset git-town.hosting-platform |
    And local Git Town setting "hosting-platform" now doesn't exist

  Scenario: undo
    When I run "git-town undo"
    And local Git Town setting "hosting-platform" is now "github"
