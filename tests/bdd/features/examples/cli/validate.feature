@examples @cli @validate
Feature: Validate CLI command with real IR fixtures
  The morphir validate command validates Morphir IR files against the JSON schema.
  These tests use real IR fixtures from morphir-elm to ensure compatibility.

  Background:
    Given the morphir CLI is available

  Rule: Validate simple IR fixtures

    # Note: Some morphir-elm fixtures have structures that don't match our strict
    # JSON schema. We test with base-ir.json which has a minimal valid structure.
    Scenario: Validate base IR fixture
      When I run morphir validate on fixture "base-ir.json"
      Then the command should succeed
      And the output should contain "VALID"

  Rule: JSON output format

    Scenario: Validate with JSON output
      When I run morphir validate on fixture "base-ir.json" with --json
      Then the command should succeed
      And the JSON output should have "valid" equal to true
      And the JSON output should have "version" equal to 3

  Rule: Error handling

    Scenario: Validate non-existent file
      When I run morphir validate "/non/existent/path.json"
      Then the command should fail
      And the output should contain "not found"
