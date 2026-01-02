# CI check script - runs all checks that should pass in CI

$ErrorActionPreference = "Stop"

$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
$repoRoot = Split-Path -Parent $scriptPath
Set-Location $repoRoot

Write-Host "Running CI checks..."

# Format check
Write-Host "Checking code formatting..."
$unformatted = gofmt -s -l .
if ($unformatted) {
    Write-Host "✗ The following files are not formatted:" -ForegroundColor Red
    Write-Host $unformatted
    Write-Host ""
    Write-Host "Run 'just fmt' to fix formatting issues." -ForegroundColor Red
    exit 1
} else {
    Write-Host "✓ Code is properly formatted" -ForegroundColor Green
}

# Build verification
Write-Host "Verifying all modules build..."
& (Join-Path $scriptPath "verify.ps1")

# Run tests
Write-Host "Running tests..."
$modules = @("cmd/morphir", "pkg/models", "pkg/tooling", "pkg/sdk", "pkg/pipeline")
foreach ($dir in $modules) {
    if (Test-Path $dir) {
        Write-Host "Testing $dir..."
        Push-Location $dir
        go test ./...
        Pop-Location
    }
}

# Lint check (if available)
if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
    Write-Host "Running linters..."
    foreach ($dir in $modules) {
        Write-Host "Linting $dir..."
        Push-Location $dir
        golangci-lint run --timeout=5m
        Pop-Location
    }
    Write-Host "✓ Linting passed" -ForegroundColor Green
} else {
    Write-Host "⚠ golangci-lint not found, skipping lint check" -ForegroundColor Yellow
}

Write-Host "✓ All CI checks passed!" -ForegroundColor Green
