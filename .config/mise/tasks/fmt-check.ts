#!/usr/bin/env bun
// #MISE description="Check Go code formatting (for CI)"

import { $ } from "bun";
import { existsSync, copyFileSync } from "fs";

async function syncChangelog() {
  const src = "CHANGELOG.md";
  const dest = "cmd/morphir/cmd/CHANGELOG.md";

  if (existsSync(src)) {
    copyFileSync(src, dest);
  }
}

async function main() {
  // Sync changelog first
  await syncChangelog();

  console.log("Checking code formatting...");

  // Ignore dependency/build caches while including new, unignored Go sources.
  const files = (await $`git ls-files -z --cached --others --exclude-standard -- '*.go'`.text())
    .split("\0")
    .filter((file) => file && existsSync(file));
  const result = files.length ? await $`gofmt -s -l ${files}`.text() : "";
  const unformatted = result.trim();

  if (unformatted) {
    console.error("The following files are not formatted:");
    console.error(unformatted);
    console.error("");
    console.error("Run 'mise run fmt' to fix formatting issues.");
    process.exit(1);
  }

  console.log("All files are properly formatted");
}

main().catch((err) => {
  console.error("Format check failed:", err.message);
  process.exit(1);
});
