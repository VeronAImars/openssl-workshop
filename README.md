# openssl-workshop

Lessons generated from OpenSSL's own documentation for the version you pick, graded on the real
artifacts you make, re-tested the night the docs change, published as a site.

Title: **OpenSSL Workshop**. Maintainer: veronaimars. Apache-2.0. A workshop built as a tribute to OpenSSL,
on OpenSSL's own documentation, at the OpenSSL project's encouragement.

## What it does

- You pick an OpenSSL release. Every lesson step cites the manual section for that release.
- You run the real `openssl` in a sandbox and make real keys, requests, certificates and chains.
- The grader checks the artifact you made, not the text you typed, and tells you which section
  explains the result.
- When OpenSSL changes a page a lesson relies on, that lesson is re-run the same night and fixed
  before you see it again.
- The site shows every command with the output captured by the run that graded it.

Lessons are files anyone can read and propose changes to. The model drafts and explains; the
checks grade; a person confirms.

## Why

The earlier public teaching module taught an end-of-life release, a chain check that passed with
a missing intermediate, and a key match that only worked for RSA, and nobody noticed for a year.
Material that cites a versioned source and is re-tested against the real binary cannot go stale
quietly.

## Status

Scaffold. Nothing runs yet. Delivery state: `docs/PHASES.md`, proven by `hack/acceptance/`.
Direction: `ROADMAP.md`. Design: `docs/DESIGN.md`. Decisions: `docs/decisions/`.

## Layout

| Path | What |
|---|---|
| `cmd/workshop/` | the CLI (`lesson check`, `lesson draft`, `site build`, later `ask`) |
| `internal/` | lesson runner, generator, site renderer (created in M0) |
| `lessons/` | lessons as files; see `lessons/README.md` |
| `site/` | templates for the static site; output is published, never committed |
| `docs/` | design, phases, decisions |
| `hack/` | `check.sh` (scrub) and the acceptance gates |

The corpus, the release images and the checks come from
[cryptographers-core](https://github.com/veronaimars/cryptographers-core), imported as a Go module
and pinned to a release.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md): the gate, the DCO sign-off (`git commit -s`), the commit
grammar, and what happens after you open an issue or a pull request. Governance is in
[GOVERNANCE.md](GOVERNANCE.md) and [MAINTAINERS.md](MAINTAINERS.md); security reports go through
[SECURITY.md](SECURITY.md). Issues labelled `good first issue` need no prior context.

## Licence

Apache-2.0 (`LICENSE`). Third-party material, if any, is listed in `NOTICE`.
