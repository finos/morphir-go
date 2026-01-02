package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Morphir configuration",
	Long:  `Commands for viewing and managing Morphir configuration settings.`,
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the resolved configuration",
	Long: `Display the fully resolved configuration after merging all sources.

Configuration is loaded from multiple sources in order of priority:
1. Built-in defaults (lowest priority)
2. System configuration (/etc/morphir/morphir.toml)
3. Global user configuration (~/.config/morphir/morphir.toml)
4. Project configuration (morphir.toml in workspace root)
5. User override (.morphir/morphir.user.toml)
6. Environment variables (MORPHIR_* prefix, highest priority)`,
	RunE: runConfigShow,
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show configuration file locations",
	Long:  `Display the paths where configuration files are searched and their status.`,
	RunE:  runConfigPath,
}

// runConfigShow displays the resolved configuration
func runConfigShow(cmd *cobra.Command, args []string) error {
	cfg, err := GetConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Display morphir section
	fmt.Println("[morphir]")
	fmt.Printf("  version = %q\n", cfg.Morphir().Version())

	// Display workspace section
	fmt.Println("\n[workspace]")
	fmt.Printf("  root = %q\n", cfg.Workspace().Root())
	fmt.Printf("  output_dir = %q\n", cfg.Workspace().OutputDir())

	// Display ir section
	fmt.Println("\n[ir]")
	fmt.Printf("  format_version = %d\n", cfg.IR().FormatVersion())
	fmt.Printf("  strict_mode = %v\n", cfg.IR().StrictMode())

	// Display codegen section
	fmt.Println("\n[codegen]")
	fmt.Printf("  targets = %v\n", cfg.Codegen().Targets())
	fmt.Printf("  template_dir = %q\n", cfg.Codegen().TemplateDir())
	fmt.Printf("  output_format = %q\n", cfg.Codegen().OutputFormat())

	// Display cache section
	fmt.Println("\n[cache]")
	fmt.Printf("  enabled = %v\n", cfg.Cache().Enabled())
	fmt.Printf("  dir = %q\n", cfg.Cache().Dir())
	fmt.Printf("  max_size = %d\n", cfg.Cache().MaxSize())

	// Display logging section
	fmt.Println("\n[logging]")
	fmt.Printf("  level = %q\n", cfg.Logging().Level())
	fmt.Printf("  format = %q\n", cfg.Logging().Format())
	fmt.Printf("  file = %q\n", cfg.Logging().File())

	// Display ui section
	fmt.Println("\n[ui]")
	fmt.Printf("  color = %v\n", cfg.UI().Color())
	fmt.Printf("  interactive = %v\n", cfg.UI().Interactive())
	fmt.Printf("  theme = %q\n", cfg.UI().Theme())

	return nil
}

// runConfigPath displays configuration file locations
func runConfigPath(cmd *cobra.Command, args []string) error {
	result, err := GetConfigResult()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	fmt.Println("Configuration sources (in priority order):")
	fmt.Println()

	sources := result.Sources()
	for _, src := range sources {
		status := "not found"
		if src.Loaded() {
			status = "loaded"
		} else if src.Error() != nil {
			status = fmt.Sprintf("error: %v", src.Error())
		}

		fmt.Printf("  [%s] %s\n", statusIcon(src.Loaded()), src.Name())
		fmt.Printf("      Path: %s\n", src.Path())
		fmt.Printf("      Status: %s\n", status)
		fmt.Printf("      Priority: %d\n\n", src.Priority())
	}

	return nil
}

// statusIcon returns a checkmark or X based on loaded status
func statusIcon(loaded bool) string {
	if loaded {
		return "✓"
	}
	return "✗"
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configPathCmd)
}
