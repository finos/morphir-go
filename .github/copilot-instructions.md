# Copilot Instructions for Morphir Go

This is a **Go port of the Morphir tooling ecosystem**. Morphir is a technology-agnostic intermediate representation (IR) for business logic and data models, enabling code generation, documentation, and analysis across multiple target platforms.

## Tech Stack

- **Language**: Go 1.24+
- **CLI Framework**: Cobra + Bubbletea (TUI)
- **Build Tool**: Just (Justfile)
- **Testing**: Go testing with BDD via godog/cucumber
- **Linting**: golangci-lint

## Project Structure

```
cmd/morphir/     - CLI application entry point
pkg/models/      - Morphir IR model types (core data structures)
pkg/tooling/     - Utilities and tools
pkg/sdk/         - SDK for building Morphir applications
pkg/pipeline/    - Processing pipelines for IR transformations
tests/bdd/       - BDD feature tests with Gherkin specifications
scripts/         - Build and CI scripts (bash + PowerShell)
docs/adr/        - Architecture Decision Records
```

Each package is a separate Go module, managed via `go.work` for development.

## Coding Guidelines

### Functional Programming First

This codebase follows functional programming principles:

- **Immutable data structures**: Prefer value types, avoid mutating state
- **Pure functions**: No side effects, deterministic outputs
- **Functional composition**: Build complex behavior from simple functions
- **Return values and errors**: Never mutate inputs, use error returns over panics

### Code Style

- Follow Go conventions and idioms
- Keep functions small and focused
- Write self-documenting code with clear names
- Organize code by feature/domain, not technical layer

### Testing Requirements

- **TDD**: Write tests before implementation
- **BDD**: Use Gherkin specifications for feature tests
- Tests must be fast, independent, and repeatable

### CLI Development

- **stdout**: Command output (data, results)
- **stderr**: Logging, diagnostics, progress, errors
- Support `--json` flag for machine-readable output on all non-interactive commands

## Build Commands

```bash
just build       # Build the CLI
just test        # Run all tests
just fmt         # Format code
just lint        # Run linters
just ci-check    # Run all CI checks
```

## Reference Implementations

When implementing features, consult these Morphir implementations for consistency:

- **finos/morphir-elm**: Primary reference (most mature)
- **finos/morphir-dotnet**: Contains IR spec and JSON schemas
- **finos/morphir-scala**: Scala patterns
- **finos/morphir-jvm**: JVM implementation

## Commit Guidelines

- Use [Conventional Commits](https://conventionalcommits.org/) format
- Types: `feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `chore:`, `ci:`
- Do NOT add AI assistants as commit co-authors (breaks EasyCLA)

## Architecture Decisions

ADRs are in `docs/adr/`. For discriminated unions/sum types in Go, see `ADR-0001`.
