@examples @cli @validate
Feature: Validate CLI command with real IR fixtures
  The morphir validate command validates Morphir IR files against the JSON schema.
  These tests use real IR fixtures from morphir-elm to ensure compatibility.

  Background:
    Given the morphir CLI is available

  Rule: Valid IR files pass validation

    Scenario Outline: Validate <fixture> passes
      When I run morphir validate on fixture "<fixture>"
      Then the command should succeed
      And the output should contain "VALID"

      Examples:
        | fixture       |
        | base-ir.json  |

  Rule: Invalid IR files report schema violations

    # These fixtures have structures that don't match our strict JSON schema.
    # The validator correctly identifies them as invalid.

    Scenario Outline: Validate <fixture> reports schema errors
      When I run morphir validate on fixture "<fixture>"
      Then the command should fail
      And the output should contain "INVALID"

      Examples:
        | fixture                    | description                      |
        | listType-ir.json           | Complex value definitions        |
        | multilevelModules-ir.json  | Module naming pattern mismatch   |
        | simpleTypeTree-ir.json     | Type definition structure        |

  Rule: JSON output format

    Scenario: Validate with JSON output for valid file
      When I run morphir validate on fixture "base-ir.json" with --json
      Then the command should succeed
      And the JSON output should have "valid" equal to true
      And the JSON output should have "version" equal to 3

    Scenario: Validate with JSON output for invalid file
      When I run morphir validate on fixture "listType-ir.json" with --json
      Then the command should fail
      And the JSON output should have "valid" equal to false
      And the JSON output should have "version" equal to 3

  Rule: Error handling

    Scenario: Validate non-existent file
      When I run morphir validate "/non/existent/path.json"
      Then the command should fail
      And the output should contain "not found"
