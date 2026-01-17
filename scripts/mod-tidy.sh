#!/usr/bin/env bash
# Run go mod tidy for all modules in the monorepo

set -e

echo "Running go mod tidy for all modules..."

cd "$(dirname "$0")/.." || exit 1

modules=(
    "cmd/morphir"
    "pkg/bindings/golang"
    "pkg/bindings/morphir-elm"
    "pkg/bindings/typemap"
    "pkg/bindings/wit"
    "pkg/config"
    "pkg/docling-doc"
    "pkg/logging"
    "pkg/models"
    "pkg/nbformat"
    "pkg/pipeline"
    "pkg/sdk"
    "pkg/task"
    "pkg/toolchain"
    "pkg/tooling"
    "pkg/vfs"
    "tests/bdd"
)

for module in "${modules[@]}"; do
    echo "Running go mod tidy in $module..."
    (cd "$module" && go mod tidy)
done

echo "All modules tidied successfully!"
