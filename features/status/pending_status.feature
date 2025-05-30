@smoke
Feature: output the current Git Town command

  Background:
    Given a Git repo with origin

  Scenario: Git Town command ran successfully
    Given I ran "git-town sync"
    When I run "git-town status --pending"
    Then Git Town prints:
      """
      """

  Scenario: Git Town command in progress
    Given the branches
      | NAME    | TYPE    | PARENT | LOCATIONS     |
      | feature | feature | main   | local, origin |
    And the commits
      | BRANCH  | LOCATION | MESSAGE                    | FILE NAME        | FILE CONTENT    |
      | main    | local    | conflicting main commit    | conflicting_file | main content    |
      | feature | local    | conflicting feature commit | conflicting_file | feature content |
    And the current branch is "feature"
    And I run "git-town sync"
    When I run "git-town status --pending"
    Then Git Town prints:
      """
      sync
      """

  Scenario: no runstate exists
    When I run "git-town status --pending"
    Then Git Town prints:
      """
      """

  Scenario: outside a Git repo
    Given I am outside a Git repo
    When I run "git-town status --pending"
    Then Git Town prints:
      """
      """
