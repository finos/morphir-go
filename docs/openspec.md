# OpenSpec workflow

This repository uses the `html-design` bundle from
[Mission Control](https://github.com/moonrockz/mission-control/tree/main/apps/openspec-ext).
New changes produce a proposal, capability specs, a standalone `design.html`
with inline CSS and SVG, and implementation tasks. Existing changes retain their
recorded schema. Follow this repository's contribution and approval rules.

## Bootstrap

Tool versions are pinned in `.config/mise/config.toml`; `.mise.toml` defines
the environment and OpenSpec tasks.

Install [mise](https://mise.jdx.dev/), then run `mise install`. OpenSpec is pinned
to 1.14.3. The shared checker is the released `mission-control-tools` 0.1.0 binary
from [Mission Control](https://github.com/moonrockz/mission-control/tree/main/apps/tools).
mise installs it from the `github:moonrockz/mission-control` entry (release tag
`tools-v0.1.0`) and checks each download against its GitHub asset digest and
artifact attestation. A prebuilt binary is available for Linux x64/ARM64, macOS
Apple Silicon/Intel and Windows x64. Windows on Arm runs the x64 build under
emulation. No other toolchain is necessary.

Then verify the checked-in bundle:

```sh
mise run openspec:verify
```

This runs `mission-control-tools openspec check --bundle` (the ownership and
bundle check), then schema and change validation. An empty spec/change tree
reports no items to validate. CI runs the same checks.

## Propose and implement a change

```sh
mise exec -- openspec new change <change-name>
mise exec -- openspec instructions design --change <change-name>
```

No schema override is needed. The design instructions target `design.html`.
Use the generated core skills to propose, explore, apply, update, sync specs and
archive changes. Skills are checked in for portable/Codex agents, Claude,
Copilot, Cursor, Grok, Gemini and OpenCode. `mise run openspec:setup` regenerates
them using an isolated core profile without changing personal OpenSpec settings.

Keep project-specific context and rules in `openspec/config.yaml` and AGENTS.md.
Existing skills outside the generated OpenSpec set are preserved.

## Update the shared bundle

`openspec/extensions.lock.json` (format 2) records the tool that wrote it
(`moonrockz/mission-control-tools` 0.1.0) and SHA-256 hashes for the seven
managed files. Git attributes preserve LF line endings so Windows checkouts keep
those hashes. Keep the schema, templates, notices and license identical to
Mission Control. Do not patch them locally. `mission-control-tools openspec check`
validates the installed snapshot against the lock; `--bundle` also requires that
the lock records the installed tool version and that each managed file is the
file in the binary.

To update the bundle, change `version` in the `github:moonrockz/mission-control`
entry in `.config/mise/config.toml`, then run:

```sh
mise install
mise exec -- mission-control-tools openspec install
mise run openspec:verify
```

Review the diff before you commit. The installer refuses edits to managed files
before it writes anything, and it writes the lock last. `openspec:check` fails
when the mise version and the lock version are different.

The schema's Apache-2.0 license and upstream OpenSpec MIT notice are retained in
`openspec/schemas/html-design/LICENSE` and `NOTICE`.
