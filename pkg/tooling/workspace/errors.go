package workspace

import (
	"errors"
	"fmt"
	"strings"
)

// Sentinel errors for workspace operations.
var (
	// ErrNotDirectory is returned when a path is expected to be a directory
	// but is not.
	ErrNotDirectory = errors.New("path is not a directory")
)

// DiscoverError represents an error that occurred during workspace discovery.
type DiscoverError struct {
	StartDir string // The directory where discovery started
	Err      error  // The underlying error
}

// Error returns the error message.
func (e *DiscoverError) Error() string {
	return fmt.Sprintf("failed to discover workspace from %q: %v", e.StartDir, e.Err)
}

// Unwrap returns the underlying error.
func (e *DiscoverError) Unwrap() error {
	return e.Err
}

// NotFoundError is returned when no workspace is found after searching
// from a starting directory to the filesystem root.
type NotFoundError struct {
	StartDir     string   // The directory where search started
	SearchedDirs []string // All directories that were searched
}

// Error returns the error message.
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("no morphir workspace found (searched %d directories from %q)",
		len(e.SearchedDirs), e.StartDir)
}

// Is allows errors.Is to match NotFoundError.
func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

// Detail returns a detailed message about the search that was performed.
func (e *NotFoundError) Detail() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("No Morphir workspace found.\n\nSearched %d directories:\n",
		len(e.SearchedDirs)))
	for _, dir := range e.SearchedDirs {
		b.WriteString(fmt.Sprintf("  - %s\n", dir))
	}
	b.WriteString("\nTo create a workspace, run: morphir init\n")
	return b.String()
}
