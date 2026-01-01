[![FINOS - Incubating](https://cdn.jsdelivr.net/gh/finos/contrib-toolbox@master/images/badge-incubating.svg)](https://community.finos.org/docs/governance/Software-Projects/stages/incubating)

# Morphir Go

A Go implementation of the Morphir tooling ecosystem. Morphir is a technology-agnostic intermediate representation (IR) for business logic and data models, enabling code generation, documentation, and analysis across multiple target platforms.

This project provides a CLI application built with [Cobra](https://github.com/spf13/cobra) and [Bubbletea](https://github.com/charmbracelet/bubbletea), along with library modules for working with Morphir IR.

## Prerequisites

- **Go 1.25.5** or later ([download](https://golang.org/dl/))
- **just** - A command runner for build orchestration ([install](https://github.com/casey/just))
- **PowerShell** (Windows only) - For running build scripts on Windows ([install](https://learn.microsoft.com/en-us/powershell/scripting/install/installing-powershell))

## Installation

### Building from Source

```sh
# Clone the repository
git clone https://github.com/finos/morphir-go.git
cd morphir-go

# Build the CLI application
just build

# The binary will be in bin/morphir
```

### Installing to System

```sh
# Build and install to $GOPATH/bin or $GOBIN
just install
```

## Usage

### Running the CLI

```sh
# Launch the interactive TUI
morphir

# Get help
morphir help

# Get help for a specific command
morphir help workspace

# Initialize a workspace (stubbed - coming soon)
morphir workspace init [path]
```

## Module Structure

This is a Go monorepo with multiple modules:

- **`cmd/morphir/`** - CLI application (Cobra + Bubbletea)
- **`pkg/models/`** - Morphir IR model types and data structures
- **`pkg/tooling/`** - Utilities and tools for working with Morphir IR
- **`pkg/sdk/`** - SDK for building applications that work with Morphir IR
- **`pkg/pipeline/`** - Processing pipelines for Morphir IR transformations

Each package is a separate Go module, managed via `go.work` for seamless development.

## Development Workflow

### Build Orchestration with Just

We use [`just`](https://github.com/casey/just) for build orchestration. Common commands:

```sh
# List all available commands
just

# Build the CLI application
just build

# Run tests across all modules
just test

# Format all Go code
just fmt

# Run linters (requires golangci-lint)
just lint

# Download dependencies for all modules
just deps

# Run go mod tidy for all modules
just mod-tidy

# Clean build artifacts
just clean

# Verify all modules build successfully
just verify

# Run CI checks (format, build, test, lint)
just ci-check
```

### Local Development and Testing

For local development, we recommend using `morphir-dev` to distinguish your development version from any installed `morphir` CLI:

```sh
# Build the development version
just build-dev

# The binary will be in bin/morphir-dev

# Run the development version directly
just run-dev

# Or run it manually
./bin/morphir-dev

# Test specific commands
./bin/morphir-dev help
./bin/morphir-dev workspace init

# Launch the TUI
./bin/morphir-dev

# Install morphir-dev to your system (makes it available in PATH)
just install-dev

# After installation, you can use morphir-dev from anywhere
morphir-dev help
```

**Why use `morphir-dev`?**
- Keeps your development version separate from any installed `morphir` CLI
- Allows you to test changes without affecting your installed version
- Makes it clear which version you're running during development
- Enables side-by-side comparison with the installed version
- Can be installed system-wide for easy access during development

### Development Setup

1. **Clone the repository**
   ```sh
   git clone https://github.com/finos/morphir-go.git
   cd morphir-go
   ```

2. **Install dependencies**
   ```sh
   just deps
   ```

3. **Build the project**
   ```sh
   # For development, use build-dev
   just build-dev
   
   # Or for standard build
   just build
   ```

4. **Run tests**
   ```sh
   just test
   ```

5. **Test your changes**
   ```sh
   # Build and run the development version
   just run-dev
   
   # Or test specific commands
   ./bin/morphir-dev help
   ```

### Go Workspace

This project uses Go workspaces (`go.work`) to manage the multi-module monorepo. The workspace file includes all modules, allowing seamless cross-module development without requiring local replacements.

### Scripts Directory

The `scripts/` directory contains reusable scripts used in build, CI, and development workflows. These scripts are referenced from the `Justfile` and can also be used directly:

- `scripts/mod-tidy.sh` / `scripts/mod-tidy.ps1` - Runs `go mod tidy` for all modules
- `scripts/install-dev.sh` / `scripts/install-dev.ps1` - Installs `morphir-dev` to Go bin directory
- `scripts/verify.sh` / `scripts/verify.ps1` - Verifies all modules build successfully
- `scripts/ci-check.sh` / `scripts/ci-check.ps1` - Runs all CI checks (format, build, test, lint)

**Cross-Platform Support:**
- All scripts have both bash (`.sh`) and PowerShell (`.ps1`) versions
- The `Justfile` automatically detects the platform and uses the appropriate script
- On Windows, PowerShell scripts are used when PowerShell is available
- On Unix-like systems (Linux, macOS), bash scripts are used

Scripts are used in the `Justfile` for complex operations and can be invoked directly when needed.

## Development Principles

This project follows functional programming principles and practices. For detailed guidance on:

- Functional programming patterns
- Test-driven development (TDD)
- Behavior-driven development (BDD)
- Code organization principles
- Morphir design principles

See **[AGENTS.md](AGENTS.md)** for comprehensive development guidelines.

## Roadmap

- [x] Initial monorepo structure
- [x] CLI application with Cobra and Bubbletea
- [x] Basic command structure (help, workspace)
- [ ] Workspace initialization implementation
- [ ] Morphir IR model support
- [ ] Tooling utilities
- [ ] SDK implementation
- [ ] Pipeline processing

## Contributing
For any questions, bugs or feature requests please open an [issue](https://github.com/finos/morphir-go/issues).

To submit a contribution:
1. Fork it (<https://github.com/finos/morphir-go/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Read our [contribution guidelines](.github/CONTRIBUTING.md) and [Community Code of Conduct](https://www.finos.org/code-of-conduct)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

_NOTE:_ Commits and pull requests to FINOS repositories will only be accepted from those contributors with an active, executed Individual Contributor License Agreement (ICLA) with FINOS OR who are covered under an existing and active Corporate Contribution License Agreement (CCLA) executed with FINOS. Commits from individuals not covered under an ICLA or CCLA will be flagged and blocked by the FINOS Clabot tool (or [EasyCLA](https://community.finos.org/docs/governance/Software-Projects/easycla)). Please note that some CCLAs require individuals/employees to be explicitly named on the CCLA.

*Need an ICLA? Unsure if you are covered under an existing CCLA? Email [help@finos.org](mailto:help@finos.org)*

## License

Copyright 2022 FINOS

Distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

SPDX-License-Identifier: [Apache-2.0](https://spdx.org/licenses/Apache-2.0)
