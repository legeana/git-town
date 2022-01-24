Feature: git town-rename-branch: errors when renaming the main branch

  Background:
    Given I am on the "main" branch

  Scenario: error when trying to rename
    When I run "git-town rename-branch main renamed-main"
    Then it runs no commands
    And it prints the error:
      """
      the main branch cannot be renamed
      """
    And I am still on the "main" branch

  Scenario: error when trying to force rename
    When I run "git-town rename-branch main renamed-main --force"
    Then it runs no commands
    And it prints the error:
      """
      the main branch cannot be renamed
      """
    And I am still on the "main" branch
