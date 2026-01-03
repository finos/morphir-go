@examples @workspace
Feature: Workspace loading with real example projects
  Example projects in the examples/ directory serve as both documentation
  and integration tests. Each example includes a test.yaml file that
  declares the expected behavior.

  Rule: Simple single-project workspaces

    Scenario Outline: Load example workspace and verify expectations
      Given the example project "<example>"
      When I load the example workspace
      Then all workspace expectations should pass

      Examples:
        | example              |
        | simple-project       |
        | morphir-elm-compat   |

  Rule: Multi-member workspaces (monorepos)

    Scenario: Load monorepo workspace with multiple members
      Given the example project "monorepo-workspace"
      When I load the example workspace
      Then all workspace expectations should pass
