#!/usr/bin/env bash
# CI check script - runs all checks that should pass in CI

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

cd "$REPO_ROOT" || exit 1

echo "Running CI checks..."

# Format check
echo "Checking code formatting..."
UNFORMATTED=$(gofmt -s -l .)
if [ -n "$UNFORMATTED" ]; then
    echo "✗ The following files are not formatted:"
    echo "$UNFORMATTED"
    echo ""
    echo "Run 'just fmt' to fix formatting issues."
    exit 1
else
    echo "✓ Code is properly formatted"
fi

# Build verification
echo "Verifying all modules build..."
"$SCRIPT_DIR/verify.sh"

# Run tests
echo "Running tests..."
for dir in cmd/morphir pkg/models pkg/tooling pkg/sdk pkg/pipeline; do
    if [ -d "$dir" ]; then
        echo "Testing $dir..."
        (cd "$dir" && go test ./...)
    fi
done

# Lint check (if available)
if command -v golangci-lint > /dev/null; then
    echo "Running linters..."
    for dir in cmd/morphir pkg/models pkg/tooling pkg/sdk pkg/pipeline; do
        echo "Linting $dir..."
        (cd "$dir" && golangci-lint run --timeout=5m)
    done
    echo "✓ Linting passed"
else
    echo "⚠ golangci-lint not found, skipping lint check"
fi

echo "✓ All CI checks passed!"
