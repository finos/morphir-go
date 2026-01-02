package config

// Option is a functional option for configuring the Load function.
type Option func(*loadOptions)

// loadOptions holds the configuration for loading.
type loadOptions struct {
	// workDir is the directory to start searching for configuration.
	// If empty, uses the current working directory.
	workDir string

	// skipSystem disables loading system-wide configuration.
	skipSystem bool

	// skipGlobal disables loading global user configuration.
	skipGlobal bool

	// skipProject disables loading project configuration.
	skipProject bool

	// skipUser disables loading user-local configuration.
	skipUser bool

	// skipEnv disables loading environment variables.
	skipEnv bool

	// configPath specifies an explicit configuration file to load.
	// When set, the normal discovery process is bypassed.
	configPath string

	// envPrefix is the prefix for environment variables (default: "MORPHIR").
	envPrefix string
}

// defaultLoadOptions returns the default load options.
func defaultLoadOptions() loadOptions {
	return loadOptions{
		workDir:     "",
		skipSystem:  false,
		skipGlobal:  false,
		skipProject: false,
		skipUser:    false,
		skipEnv:     false,
		configPath:  "",
		envPrefix:   "MORPHIR",
	}
}

// WithWorkDir sets the working directory for configuration discovery.
func WithWorkDir(dir string) Option {
	return func(o *loadOptions) {
		o.workDir = dir
	}
}

// WithConfigPath sets an explicit configuration file path.
// When set, the normal discovery process is bypassed and only this file is loaded.
func WithConfigPath(path string) Option {
	return func(o *loadOptions) {
		o.configPath = path
	}
}

// WithEnvPrefix sets the prefix for environment variables.
// The default prefix is "MORPHIR".
func WithEnvPrefix(prefix string) Option {
	return func(o *loadOptions) {
		o.envPrefix = prefix
	}
}

// WithoutSystem disables loading system-wide configuration (/etc/morphir/morphir.toml).
func WithoutSystem() Option {
	return func(o *loadOptions) {
		o.skipSystem = true
	}
}

// WithoutGlobal disables loading global user configuration (~/.config/morphir/morphir.toml).
func WithoutGlobal() Option {
	return func(o *loadOptions) {
		o.skipGlobal = true
	}
}

// WithoutProject disables loading project configuration (morphir.toml).
func WithoutProject() Option {
	return func(o *loadOptions) {
		o.skipProject = true
	}
}

// WithoutUser disables loading user-local configuration (.morphir/morphir.user.toml).
func WithoutUser() Option {
	return func(o *loadOptions) {
		o.skipUser = true
	}
}

// WithoutEnv disables loading configuration from environment variables.
func WithoutEnv() Option {
	return func(o *loadOptions) {
		o.skipEnv = true
	}
}

// Load loads configuration from all available sources.
// Sources are loaded in priority order, with higher priority sources
// overriding lower priority ones.
//
// The load order (from lowest to highest priority) is:
//  1. Built-in defaults
//  2. System configuration (/etc/morphir/morphir.toml)
//  3. Global user configuration (~/.config/morphir/morphir.toml)
//  4. Project configuration (morphir.toml or .morphir/morphir.toml)
//  5. User-local configuration (.morphir/morphir.user.toml)
//  6. Environment variables (MORPHIR_*)
//
// Options can be used to customize the loading behavior.
func Load(opts ...Option) (Config, error) {
	options := defaultLoadOptions()
	for _, opt := range opts {
		opt(&options)
	}

	// Start with defaults
	cfg := Default()

	// TODO: Implement actual loading from sources
	// This will be implemented in subsequent tasks:
	// - morphir-go-p4z: TOML source
	// - morphir-go-4bh: Environment variable source
	// - morphir-go-sa3: XDG path resolution
	// - morphir-go-843: Config merger
	// - morphir-go-u82: Loader orchestration

	_ = options // suppress unused warning for now

	return cfg, nil
}

// LoadWithDetails loads configuration and returns detailed information
// about the sources that were loaded.
func LoadWithDetails(opts ...Option) (LoadResult, error) {
	cfg, err := Load(opts...)
	if err != nil {
		return LoadResult{}, err
	}

	// TODO: Populate sources from actual loading process
	return LoadResult{
		config:  cfg,
		sources: nil,
	}, nil
}
