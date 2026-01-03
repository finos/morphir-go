package workspace

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// ConfigFilter specifies which config formats are allowed.
type ConfigFilter int

const (
	// ConfigFilterAll allows all config formats (toml and json).
	ConfigFilterAll ConfigFilter = iota
	// ConfigFilterTOML allows only TOML config format.
	ConfigFilterTOML
	// ConfigFilterJSON allows only JSON config format.
	ConfigFilterJSON
)

// DiscoverMembers finds all workspace members matching the given glob patterns.
// It returns absolute paths to directories containing valid project configurations.
//
// Parameters:
//   - root: absolute path to workspace root
//   - patterns: glob patterns for member discovery (e.g., "packages/*", "libs/**")
//   - excludes: patterns to exclude from results (e.g., "**/testdata")
//
// Returns paths in sorted order for deterministic behavior.
func DiscoverMembers(root string, patterns []string, excludes []string) ([]string, error) {
	if len(patterns) == 0 {
		return []string{}, nil
	}

	if err := validateRootDirectory(root); err != nil {
		return nil, err
	}

	memberSet := collectMembers(root, patterns, excludes)
	return sortedKeys(memberSet), nil
}

// validateRootDirectory ensures the root path exists and is a directory.
func validateRootDirectory(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return &DiscoverError{StartDir: root, Err: ErrNotDirectory}
	}
	return nil
}

// collectMembers gathers all matching member directories into a set.
func collectMembers(root string, patterns, excludes []string) map[string]struct{} {
	memberSet := make(map[string]struct{})

	for _, pattern := range patterns {
		// Parse the pattern to extract config filter and base pattern
		basePattern, filter := parsePatternFilter(pattern)
		matches := expandPattern(root, basePattern)
		for _, match := range matches {
			if isValidMemberWithFilter(root, match, excludes, filter) {
				memberSet[match] = struct{}{}
			}
		}
	}

	return memberSet
}

// parsePatternFilter extracts the extension filter from a pattern.
// Patterns ending with .toml, .json, or .{toml,json} specify config format restrictions.
// Returns the base pattern (with extension stripped) and the filter to apply.
func parsePatternFilter(pattern string) (string, ConfigFilter) {
	// Check for brace expansion patterns like *.{toml,json}
	if strings.HasSuffix(pattern, ".{toml,json}") || strings.HasSuffix(pattern, ".{json,toml}") {
		return strings.TrimSuffix(strings.TrimSuffix(pattern, ".{toml,json}"), ".{json,toml}"), ConfigFilterAll
	}

	// Check for explicit .toml extension
	if strings.HasSuffix(pattern, ".toml") {
		return strings.TrimSuffix(pattern, ".toml"), ConfigFilterTOML
	}

	// Check for explicit .json extension
	if strings.HasSuffix(pattern, ".json") {
		return strings.TrimSuffix(pattern, ".json"), ConfigFilterJSON
	}

	// No extension filter - accept all formats
	return pattern, ConfigFilterAll
}

// expandPattern resolves a glob pattern to matching paths.
func expandPattern(root, pattern string) []string {
	absPattern := filepath.Join(root, pattern)
	matches, err := doublestar.FilepathGlob(absPattern)
	if err != nil {
		return nil
	}
	return matches
}

// isValidMember checks if a path is a valid workspace member.
func isValidMember(root, path string, excludes []string) bool {
	return isValidMemberWithFilter(root, path, excludes, ConfigFilterAll)
}

// isValidMemberWithFilter checks if a path is a valid workspace member with a config filter.
func isValidMemberWithFilter(root, path string, excludes []string, filter ConfigFilter) bool {
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return false
	}
	if isExcluded(root, path, excludes) {
		return false
	}
	return HasProjectConfigWithFilter(path, filter)
}

// sortedKeys returns the keys of a map as a sorted slice.
func sortedKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// isExcluded checks if a path matches any exclude pattern.
func isExcluded(root, path string, excludes []string) bool {
	if len(excludes) == 0 {
		return false
	}

	// Get relative path for pattern matching
	relPath, err := filepath.Rel(root, path)
	if err != nil {
		return false
	}

	for _, pattern := range excludes {
		// Try matching against the relative path
		matched, err := doublestar.Match(pattern, relPath)
		if err == nil && matched {
			return true
		}

		// Also try matching against just the base name for simple patterns
		matched, err = doublestar.Match(pattern, filepath.Base(path))
		if err == nil && matched {
			return true
		}
	}

	return false
}

// HasProjectConfig checks if a directory contains a valid project configuration.
// A valid project config is one of:
//   - morphir.toml with [project] section
//   - .morphir/morphir.toml with [project] section
//   - morphir.json (legacy format)
func HasProjectConfig(dir string) bool {
	return HasProjectConfigWithFilter(dir, ConfigFilterAll)
}

// HasProjectConfigWithFilter checks if a directory contains a valid project configuration
// matching the specified filter.
func HasProjectConfigWithFilter(dir string, filter ConfigFilter) bool {
	_, _, found := FindProjectConfigWithFilter(dir, filter)
	return found
}

// FindProjectConfig locates the project configuration file in a directory.
// It returns the path, format, and whether it was found.
//
// Detection priority:
//  1. morphir.toml with [project] section
//  2. .morphir/morphir.toml with [project] section
//  3. morphir.json (legacy format)
func FindProjectConfig(dir string) (path string, format string, found bool) {
	return FindProjectConfigWithFilter(dir, ConfigFilterAll)
}

// FindProjectConfigWithFilter locates the project configuration file in a directory,
// respecting the specified format filter.
// It returns the path, format, and whether it was found.
//
// Detection priority:
//  1. morphir.toml with [project] section (if filter allows TOML)
//  2. .morphir/morphir.toml with [project] section (if filter allows TOML)
//  3. morphir.json (legacy format) (if filter allows JSON)
func FindProjectConfigWithFilter(dir string, filter ConfigFilter) (path string, format string, found bool) {
	// Check TOML configs if filter allows
	if filter == ConfigFilterAll || filter == ConfigFilterTOML {
		// Check morphir.toml first
		tomlPath := filepath.Join(dir, "morphir.toml")
		if hasProjectSection(tomlPath) {
			return tomlPath, "toml", true
		}

		// Check .morphir/morphir.toml
		hiddenTomlPath := filepath.Join(dir, ".morphir", "morphir.toml")
		if hasProjectSection(hiddenTomlPath) {
			return hiddenTomlPath, "toml", true
		}
	}

	// Check JSON config if filter allows
	if filter == ConfigFilterAll || filter == ConfigFilterJSON {
		jsonPath := filepath.Join(dir, "morphir.json")
		if isRegularFile(jsonPath) {
			return jsonPath, "json", true
		}
	}

	return "", "", false
}

// hasProjectSection checks if a TOML file exists and contains a [project] section.
func hasProjectSection(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	// Simple check for [project] section
	// This is a basic check - full parsing would be done when loading
	content := string(data)
	return strings.Contains(content, "[project]")
}

// isRegularFile checks if a path exists and is a regular file.
func isRegularFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
