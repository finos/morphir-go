#!/usr/bin/env bash
# Verify all modules build successfully

set -e

echo "Verifying all modules build..."

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
)

for module in "${modules[@]}"; do
    echo "Building $module..."
    go build "./$module"
done

echo "All modules build successfully!"
