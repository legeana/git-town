Feature: git town: show an error message when minimum Git version is not satisfied

  Scenario: using an unsupported Git Version
    Given my computer has Git "2.6.2" installed
    When I run "git-town config"
    Then it prints the error:
      """
      Git Town requires Git 2.7.0 or higher
      """
