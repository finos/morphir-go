// Package report provides validation report generation in various formats.
package report

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/finos/morphir-go/pkg/tooling/validation"
)

// Format represents the output format for reports.
type Format string

const (
	// FormatMarkdown generates a Markdown report.
	FormatMarkdown Format = "markdown"
)

// Generator generates validation reports.
type Generator interface {
	// Generate creates a report from the validation result.
	Generate(result *validation.Result) string
}

// NewGenerator creates a new report generator for the specified format.
func NewGenerator(format Format) Generator {
	switch format {
	case FormatMarkdown:
		return &MarkdownGenerator{}
	default:
		return &MarkdownGenerator{}
	}
}

// ParsedError represents a parsed validation error with context.
type ParsedError struct {
	Path        string
	Message     string
	Expected    string
	Actual      string
	Suggestion  string
	NestedPaths []string
}

// parseValidationError extracts structured information from a validation error string.
func parseValidationError(errStr string) []ParsedError {
	var errors []ParsedError

	// Pattern to match path and message: "at '/path': message"
	pathPattern := regexp.MustCompile(`at '([^']+)':\s*(.+)`)

	lines := strings.Split(errStr, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "-") {
			continue
		}

		matches := pathPattern.FindStringSubmatch(line)
		if len(matches) >= 3 {
			path := matches[1]
			message := matches[2]

			parsed := ParsedError{
				Path:    path,
				Message: message,
			}

			// Extract expected/actual from common patterns
			if strings.Contains(message, "value must be") {
				parsed.Expected = extractAfter(message, "value must be ")
			}
			if strings.Contains(message, "got ") {
				parts := strings.SplitN(message, "got ", 2)
				if len(parts) == 2 {
					parsed.Actual = strings.TrimSuffix(strings.Split(parts[1], ",")[0], ")")
				}
			}
			if strings.Contains(message, "want ") {
				parts := strings.SplitN(message, "want ", 2)
				if len(parts) == 2 {
					parsed.Expected = strings.TrimSpace(parts[1])
				}
			}

			// Generate suggestions based on error type
			parsed.Suggestion = generateSuggestion(parsed)

			errors = append(errors, parsed)
		}
	}

	// If no structured errors found, create a generic one
	if len(errors) == 0 && errStr != "" {
		errors = append(errors, ParsedError{
			Path:    "/",
			Message: errStr,
		})
	}

	return errors
}

func extractAfter(s, prefix string) string {
	idx := strings.Index(s, prefix)
	if idx == -1 {
		return ""
	}
	result := s[idx+len(prefix):]
	// Trim quotes and trailing punctuation
	result = strings.Trim(result, "\"' ")
	return result
}

func generateSuggestion(err ParsedError) string {
	msg := strings.ToLower(err.Message)

	switch {
	case strings.Contains(msg, "pattern"):
		return "The value doesn't match the required naming pattern. Check that identifiers use lowercase letters and numbers only."

	case strings.Contains(msg, "oneof failed"):
		return "The value doesn't match any of the expected type variants. This usually means the IR structure differs from the schema definition."

	case strings.Contains(msg, "allof failed"):
		return "The value doesn't satisfy all required constraints. Check that all required fields are present and correctly formatted."

	case strings.Contains(msg, "missing properties"):
		return "Required fields are missing from this object. Add the missing properties listed in the error."

	case strings.Contains(msg, "minitems") || strings.Contains(msg, "maxitems"):
		return "The array has the wrong number of elements. Check the IR structure against the schema specification."

	case strings.Contains(msg, "got array, want object"):
		return "Expected an object but found an array. This may indicate a version mismatch between the IR and schema."

	case strings.Contains(msg, "got object, want array"):
		return "Expected an array but found an object. This may indicate a version mismatch between the IR and schema."

	case strings.Contains(msg, "got null"):
		return "A null value was found where a value was required. Ensure all required fields have values."

	default:
		return ""
	}
}

// explainJSONPath provides a human-readable explanation of a JSON path.
func explainJSONPath(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")

	var explanations []string
	for i, part := range parts {
		switch part {
		case "distribution":
			explanations = append(explanations, "the distribution")
		case "modules":
			explanations = append(explanations, "modules list")
		case "types":
			explanations = append(explanations, "type definitions")
		case "values":
			explanations = append(explanations, "value definitions")
		case "value":
			explanations = append(explanations, "value specification")
		case "body":
			explanations = append(explanations, "function body")
		case "doc":
			explanations = append(explanations, "documentation field")
		default:
			// Check if it's a numeric index
			if isNumeric(part) {
				if i > 0 && i < len(parts) {
					explanations = append(explanations, fmt.Sprintf("element %s", part))
				}
			}
		}
	}

	if len(explanations) == 0 {
		return path
	}

	return strings.Join(explanations, " → ")
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}

// countUniqueIssues counts unique issue types from parsed errors.
func countUniqueIssues(errors []ParsedError) map[string]int {
	counts := make(map[string]int)
	for _, e := range errors {
		msg := e.Message
		// Normalize the message to group similar issues
		switch {
		case strings.Contains(msg, "pattern"):
			counts["Pattern mismatch"]++
		case strings.Contains(msg, "oneOf failed"):
			counts["Type variant mismatch"]++
		case strings.Contains(msg, "allOf failed"):
			counts["Constraint validation failed"]++
		case strings.Contains(msg, "missing properties"):
			counts["Missing required fields"]++
		case strings.Contains(msg, "minItems") || strings.Contains(msg, "maxItems"):
			counts["Array length mismatch"]++
		default:
			counts["Other validation errors"]++
		}
	}
	return counts
}
