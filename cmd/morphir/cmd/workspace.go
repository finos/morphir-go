package cmd

import (
	"errors"
	"fmt"

	"github.com/finos/morphir-go/pkg/tooling/workspace"
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage Morphir workspaces",
	Long:  `Commands for managing Morphir workspaces, including initialization and configuration.`,
}

var (
	initHidden bool
	initName   string
)

var workspaceInitCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new Morphir workspace",
	Long: `Initialize a new Morphir workspace in the specified directory.
If no path is provided, the current directory will be used.

This creates:
  - morphir.toml (or .morphir/morphir.toml with --hidden)
  - .morphir/ directory with subdirectories
  - .morphir/.gitignore for common exclusions`,
	Args: cobra.MaximumNArgs(1),
	RunE: runWorkspaceInit,
}

// runWorkspaceInit executes the workspace init command
func runWorkspaceInit(cmd *cobra.Command, args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Determine config style
	style := workspace.ConfigStyleRoot
	if initHidden {
		style = workspace.ConfigStyleHidden
	}

	// Initialize workspace
	result, err := workspace.Init(workspace.InitOptions{
		Path:  path,
		Name:  initName,
		Style: style,
	})
	if err != nil {
		// Check for already exists error
		var alreadyExists *workspace.AlreadyExistsError
		if errors.As(err, &alreadyExists) {
			return fmt.Errorf("workspace already exists at %s", alreadyExists.ExistingRoot)
		}
		return fmt.Errorf("failed to initialize workspace: %w", err)
	}

	// Report success
	ws := result.Workspace()
	fmt.Printf("Initialized Morphir workspace at %s\n", ws.Root())
	fmt.Printf("  Config: %s\n", ws.ConfigPath())

	if len(result.CreatedDirs()) > 0 {
		fmt.Println("\nCreated directories:")
		for _, dir := range result.CreatedDirs() {
			fmt.Printf("  - %s\n", dir)
		}
	}

	if len(result.CreatedFiles()) > 0 {
		fmt.Println("\nCreated files:")
		for _, file := range result.CreatedFiles() {
			fmt.Printf("  - %s\n", file)
		}
	}

	return nil
}

func init() {
	workspaceCmd.AddCommand(workspaceInitCmd)

	// Flags for workspace init
	workspaceInitCmd.Flags().BoolVar(&initHidden, "hidden", false,
		"Place morphir.toml inside .morphir/ directory")
	workspaceInitCmd.Flags().StringVar(&initName, "name", "",
		"Project name (defaults to directory name)")
}
