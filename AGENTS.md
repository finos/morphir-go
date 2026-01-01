# AGENTS.md - Agentic Hints for Morphir Go

This document provides guidance for AI assistants and developers working on the Morphir Go project.

## Project Overview

This is a **Go port of the Morphir tooling ecosystem**. Morphir is a technology-agnostic intermediate representation (IR) for business logic and data models, enabling code generation, documentation, and analysis across multiple target platforms.

### Reference Implementations

When implementing features, refer to these existing Morphir implementations for consistency:

- **finos/morphir** - Core Morphir project and IR specification
- **finos/morphir-elm** - Reference implementation in Elm (most mature)
- **finos/morphir-jvm** - JVM-based implementation
- **finos/morphir-scala** - Scala implementation
- **finos/morphir-dotnet** - .NET implementation (contains IR spec and JSON schemas in documentation)
- **finos/morphir-rust** - Early-stage Rust tooling

### Morphir IR Specification

The Morphir IR specification and JSON schemas are available in the morphir-dotnet documentation. Always maintain alignment with the official IR specification when implementing features.

## Core Morphir Design Principles

### Functional Programming and Functional Domain Practices

**Functional programming is fundamental to this codebase.** All code should follow functional programming principles:

- **Immutable data structures** - Prefer immutable types and avoid mutating state
- **Pure functions** - Functions should have no side effects when possible
- **Separation of concerns** - Clearly define I/O boundaries
- **Functional composition** - Build complex behavior from simple, composable functions
- **Domain-driven design alignment** - Model the domain using functional patterns

### Code Organization Principles

When writing code:

1. **Prefer pure functions over impure ones**
   - Pure functions are easier to test, reason about, and compose
   - Isolate side effects to I/O boundaries

2. **Return values and errors instead of mutating state**
   - Functions should return new values rather than modifying inputs
   - Use error returns instead of panics where possible

3. **Use immutable data structures**
   - Prefer structs with value semantics
   - Avoid global mutable state
   - Use functional update patterns (return new instances)

4. **Separate I/O from business logic**
   - Keep business logic pure and testable
   - Isolate file system, network, and user interaction to boundaries

5. **Functional composition over imperative flow**
   - Compose small functions into larger behaviors
   - Use higher-order functions where appropriate

## Development Practices

### Test-Driven Development (TDD)

**Write tests before implementation.** Follow the TDD cycle:

1. Write a failing test
2. Write minimal code to make it pass
3. Refactor while keeping tests green

Tests should be:
- Fast
- Independent
- Repeatable
- Self-validating
- Timely

### Behavior-Driven Development (BDD)

**Specify behavior before implementation.** Use BDD for feature specifications:

- Write feature specifications in clear, domain language
- Define scenarios with Given-When-Then structure
- Ensure tests reflect business requirements

### Functional Domain Modeling

**Model the domain using functional patterns:**

- Use algebraic data types where appropriate
- Model domain concepts as immutable types
- Separate domain logic from infrastructure concerns
- Use functional composition to build domain workflows

### Clean, Well-Organized Code

- Write self-documenting code with clear names
- Keep functions small and focused
- Follow Go conventions and idioms
- Organize code by feature/domain, not by technical layer

## CLI Development Guidelines

### Output Format and Streams

**Separation of Output Streams:**
- **stdout** - Use for actual command output (data, results, structured output)
- **stderr** - Use for logging, diagnostics, progress messages, and error messages

This separation allows users to pipe command output while still seeing diagnostic information, and enables proper shell redirection patterns.

```go
// Good: Output to stdout, diagnostics to stderr
func runCommand(cmd *cobra.Command, args []string) error {
    fmt.Fprintf(os.Stderr, "Processing...\n") // Diagnostic message
    result := processData(args)
    fmt.Fprintf(os.Stdout, "%s\n", result) // Actual output
    return nil
}

// Avoid: Mixing output streams
func runCommand(cmd *cobra.Command, args []string) error {
    fmt.Println("Processing...") // Goes to stdout - wrong!
    fmt.Println(result) // Actual output
    return nil
}
```

### JSON Output Support

**All non-interactive commands should support a `--json` flag** to output results in JSON format. This enables:
- Machine-readable output for scripting and automation
- Integration with other tools and pipelines
- Consistent structured output across commands

**Implementation Pattern:**
```go
var jsonOutput bool

func init() {
    validateCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output results as JSON")
}

func runValidate(cmd *cobra.Command, args []string) error {
    result := validateIR(args)
    
    if jsonOutput {
        // Output JSON to stdout
        encoder := json.NewEncoder(os.Stdout)
        encoder.SetIndent("", "  ")
        return encoder.Encode(result)
    }
    
    // Output human-readable format to stdout
    fmt.Fprintf(os.Stdout, "%s\n", formatHumanReadable(result))
    return nil
}
```

**Guidelines:**
- JSON output should be written to **stdout** (not stderr)
- Logging and diagnostics should still go to **stderr** even when `--json` is used
- JSON output should be well-structured and follow consistent schemas
- When `--json` is used, avoid mixing JSON with human-readable text
- Use proper JSON encoding with indentation for readability (when appropriate)

**Interactive Commands:**
- Commands that launch interactive UIs (like the root TUI) do not need `--json` support
- Commands that can be both interactive and non-interactive should support `--json` for non-interactive mode

## When Contributing

### Code Style

1. **Follow functional programming patterns**
   - Avoid mutable state
   - Prefer pure functions
   - Use functional composition

2. **Write tests first (TDD)**
   - Start with failing tests
   - Implement to make tests pass
   - Refactor with confidence

3. **Use BDD for feature specifications**
   - Define behavior clearly
   - Write scenarios that reflect requirements

4. **Reference other Morphir implementations**
   - Check how similar features are implemented in other languages
   - Maintain consistency with Morphir IR specification
   - Learn from reference implementations (especially morphir-elm)

5. **Maintain alignment with Morphir IR specification**
   - Ensure compatibility with the official IR
   - Validate against JSON schemas when available
   - Test interoperability with other Morphir tools

6. **Follow CLI development guidelines**
   - Separate stdout (output) from stderr (logging/diagnostics)
   - Add `--json` flag support to all non-interactive commands
   - Ensure JSON output is well-structured and consistent

### Example: Functional Pattern

```go
// Good: Pure function, immutable data
func ProcessModel(model Model) (ProcessedModel, error) {
    // Process without mutating input
    processed := transform(model)
    return processed, nil
}

// Avoid: Mutating input
func ProcessModel(model *Model) error {
    // Mutating model - not functional
    model.Field = newValue
    return nil
}
```

### Example: Functional Update Pattern

```go
// Good: Return new instance
func UpdateState(state State, value int) State {
    return State{
        Count: state.Count + value,
        // Copy other fields
    }
}

// Avoid: Mutating state
func UpdateState(state *State, value int) {
    state.Count += value
}
```

## Project Structure

- `cmd/morphir/` - CLI application (Cobra + Bubbletea)
- `pkg/models/` - Morphir IR model types
- `pkg/tooling/` - Utilities and tools
- `pkg/sdk/` - SDK for building Morphir applications
- `pkg/pipeline/` - Processing pipelines for IR transformations

Each package is a separate Go module, managed via `go.work` for development.

## Build and Development

- Use `just` for build orchestration (see `Justfile`)
- Run `just build` to build the CLI
- Run `just test` to run all tests
- Run `just fmt` to format code
- Run `just lint` to run linters

### Scripts Directory

The `scripts/` directory contains reusable shell scripts used in build, CI, and development workflows. These scripts are referenced from the `Justfile` for longer task definitions and can also be used directly.

**Available Scripts:**
- `scripts/mod-tidy.sh` / `scripts/mod-tidy.ps1` - Runs `go mod tidy` for all modules in the monorepo
- `scripts/install-dev.sh` / `scripts/install-dev.ps1` - Installs the `morphir-dev` binary to the Go bin directory
- `scripts/verify.sh` / `scripts/verify.ps1` - Verifies all modules build successfully
- `scripts/ci-check.sh` / `scripts/ci-check.ps1` - Runs all CI checks (format, build, test, lint)

**Cross-Platform Support:**
- All scripts have both bash (`.sh`) and PowerShell (`.ps1`) versions for cross-platform support
- The `Justfile` uses `scripts/detect-os.sh` for proper OS detection (windows, linux, darwin)
- OS detection is done via helper recipes (`_os`, `_bin-ext`, `_script-ext`, `_powershell`)
- Windows: Uses PowerShell scripts (`.ps1`) and adds `.exe` extension to binaries
- Unix-like (Linux, macOS): Uses bash scripts (`.sh`) and no extension for binaries
- The Justfile automatically selects the correct scripts and binary extensions based on detected OS

**Guidelines for Scripts:**
- Scripts should be executable (`chmod +x` for `.sh` files)
- Bash scripts: Use `#!/usr/bin/env bash` shebang and `set -e` to exit on errors
- PowerShell scripts: Use `$ErrorActionPreference = "Stop"` for error handling
- Scripts should be idempotent when possible
- Keep scripts focused on a single task
- Use scripts in `Justfile` for complex or multi-step operations
- Scripts can be used directly or via `just` commands
- When adding new scripts, create both `.sh` and `.ps1` versions for cross-platform support

**Adding New Scripts:**
- Place new scripts in the `scripts/` directory
- Create both `.sh` (bash) and `.ps1` (PowerShell) versions
- Make bash scripts executable: `chmod +x scripts/your-script.sh`
- Reference them in the `Justfile` with platform detection
- Document their purpose in comments at the top of the script

## Questions?

When in doubt:
1. Check reference implementations (especially morphir-elm)
2. Consult Morphir IR specification
3. Follow functional programming principles
4. Write tests first
5. Keep code simple and composable
