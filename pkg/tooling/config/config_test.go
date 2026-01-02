package config

import (
	"testing"
)

func TestDefaultReturnsExpectedDefaults(t *testing.T) {
	cfg := Default()

	// IR section defaults
	if got := cfg.IR().FormatVersion(); got != 3 {
		t.Errorf("IR.FormatVersion: want 3, got %d", got)
	}
	if got := cfg.IR().StrictMode(); got != false {
		t.Errorf("IR.StrictMode: want false, got %t", got)
	}

	// Workspace section defaults
	if got := cfg.Workspace().OutputDir(); got != ".morphir" {
		t.Errorf("Workspace.OutputDir: want .morphir, got %q", got)
	}
	if got := cfg.Workspace().Root(); got != "" {
		t.Errorf("Workspace.Root: want empty, got %q", got)
	}

	// Codegen section defaults
	if got := cfg.Codegen().OutputFormat(); got != "pretty" {
		t.Errorf("Codegen.OutputFormat: want pretty, got %q", got)
	}
	if got := cfg.Codegen().Targets(); got != nil {
		t.Errorf("Codegen.Targets: want nil, got %v", got)
	}

	// Cache section defaults
	if got := cfg.Cache().Enabled(); got != true {
		t.Errorf("Cache.Enabled: want true, got %t", got)
	}
	if got := cfg.Cache().MaxSize(); got != 0 {
		t.Errorf("Cache.MaxSize: want 0, got %d", got)
	}

	// Logging section defaults
	if got := cfg.Logging().Level(); got != "info" {
		t.Errorf("Logging.Level: want info, got %q", got)
	}
	if got := cfg.Logging().Format(); got != "text" {
		t.Errorf("Logging.Format: want text, got %q", got)
	}

	// UI section defaults
	if got := cfg.UI().Color(); got != true {
		t.Errorf("UI.Color: want true, got %t", got)
	}
	if got := cfg.UI().Interactive(); got != true {
		t.Errorf("UI.Interactive: want true, got %t", got)
	}
	if got := cfg.UI().Theme(); got != "default" {
		t.Errorf("UI.Theme: want default, got %q", got)
	}
}

func TestCodegenTargetsDefensiveCopy(t *testing.T) {
	// Create a config with targets
	cfg := Config{
		codegen: CodegenSection{
			targets: []string{"go", "scala"},
		},
	}

	// Get targets and modify the returned slice
	targets := cfg.Codegen().Targets()
	if targets == nil {
		t.Fatal("expected non-nil targets")
	}
	targets[0] = "mutated"

	// Original should be unchanged
	originalTargets := cfg.Codegen().Targets()
	if originalTargets[0] != "go" {
		t.Errorf("defensive copy failed: original was mutated to %q", originalTargets[0])
	}
}

func TestLoadResultSourcesDefensiveCopy(t *testing.T) {
	// Create a LoadResult with sources
	result := LoadResult{
		sources: []SourceInfo{
			{name: "project", path: "/path/to/morphir.toml", priority: 4},
		},
	}

	// Get sources and modify the returned slice
	sources := result.Sources()
	if sources == nil {
		t.Fatal("expected non-nil sources")
	}
	sources[0] = SourceInfo{name: "mutated"}

	// Original should be unchanged
	originalSources := result.Sources()
	if originalSources[0].Name() != "project" {
		t.Errorf("defensive copy failed: original was mutated to %q", originalSources[0].Name())
	}
}

func TestSourceInfoAccessors(t *testing.T) {
	info := SourceInfo{
		name:     "project",
		path:     "/home/user/project/morphir.toml",
		priority: 4,
	}

	if got := info.Name(); got != "project" {
		t.Errorf("Name: want project, got %q", got)
	}
	if got := info.Path(); got != "/home/user/project/morphir.toml" {
		t.Errorf("Path: want /home/user/project/morphir.toml, got %q", got)
	}
	if got := info.Priority(); got != 4 {
		t.Errorf("Priority: want 4, got %d", got)
	}
}

func TestLoadReturnsDefaults(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: unexpected error: %v", err)
	}

	// Should return defaults when no config files exist
	if got := cfg.IR().FormatVersion(); got != 3 {
		t.Errorf("IR.FormatVersion: want 3, got %d", got)
	}
}

func TestLoadWithDetailsReturnsConfig(t *testing.T) {
	result, err := LoadWithDetails()
	if err != nil {
		t.Fatalf("LoadWithDetails: unexpected error: %v", err)
	}

	cfg := result.Config()
	if got := cfg.IR().FormatVersion(); got != 3 {
		t.Errorf("IR.FormatVersion: want 3, got %d", got)
	}
}
