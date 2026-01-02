package configloader

import (
	"github.com/finos/morphir-go/pkg/tooling/config"
)

// ToConfig converts a merged map[string]any to a config.Config struct.
// This function extracts values from the map using the Get* helpers
// and constructs an immutable Config instance.
func ToConfig(m map[string]any) config.Config {
	// Use the config package's internal construction
	// Since Config has unexported fields, we need to use a builder pattern
	// or the config package needs to expose a way to construct from raw values.
	//
	// For now, we return the default config as a placeholder.
	// The actual implementation will need coordination with the config package.
	return config.Default()
}

// ConfigFromMap creates a Config from a map[string]any.
// This is an internal function that uses reflection or direct field access.
//
// The map structure should match:
//
//	{
//	  "morphir": { "version": string },
//	  "workspace": { "root": string, "output_dir": string },
//	  "ir": { "format_version": int, "strict_mode": bool },
//	  "codegen": { "targets": []string, "template_dir": string, "output_format": string },
//	  "cache": { "enabled": bool, "dir": string, "max_size": int },
//	  "logging": { "level": string, "format": string, "file": string },
//	  "ui": { "color": bool, "interactive": bool, "theme": string },
//	}
func ConfigFromMap(m map[string]any) config.Config {
	// For the BDD tests, we need to be able to verify values from the raw map.
	// The config.Config struct has unexported fields, so we can't directly
	// construct it. Instead, the BDD tests will work with the raw map.
	//
	// When the public config.Load() is wired up (morphir-go-sul),
	// it will need to either:
	// 1. Expose a config.FromMap() constructor in the config package
	// 2. Use a builder pattern
	// 3. Return the raw map and let callers use GetString/GetInt etc.
	//
	// For now, return defaults.
	return config.Default()
}
