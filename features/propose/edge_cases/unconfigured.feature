@messyoutput
Feature: ask for missing configuration

  Background:
    Given a Git repo with origin
    And Git Town is not configured
    And the origin is "https://github.com/git-town/git-town.git"
    When I run "git-town propose" and enter into the dialog:
      | DIALOG      | KEYS  |
      | main branch | enter |

  Scenario: result
    Then the main branch is now "main"
    And Git Town prints the error:
      """
      cannot propose the main branch
      """
