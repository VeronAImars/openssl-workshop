# Roadmap

Grouped by minor version. What a version contains is the promise; when it ships is not. Delivery
state lives in `docs/PHASES.md` and is proven by `hack/acceptance/phases.json`; this file never
claims more than the ledger does.

One Go binary, `workshop`. Lessons are files. The corpus, the pinned release images and the checks
come from the core; this repository decides what to teach, generates it from the manuals, grades
it and publishes it.

## v0.1 — the first lesson, graded and published

- M0 walking skeleton: one lesson graded on two releases in CI, its page on the static site, the
  earlier public module link redirected there
- M1 the PKI basics module: key, request, chain, verify, revocation, with current commands; then the
  your-stack module, which teaches from the learner's own host probe
- M2 drift: when the core's nightly index changes a cited section, the lessons that cite it are
  re-run, one issue per failure

## v0.2 — generated lessons and the shell tutor

- M3 the generator: `workshop lesson draft <command>` drafts a lesson from the manual section and
  its EXAMPLES block; a person confirms before it is published
- M4 the shell tutor: completion for `openssl` itself, one line per option from the manual for the
  installed release, with options that no longer exist flagged; and `workshop ask`, which answers from
  that release's index and prints the section it used; it never grades

## v0.3 — more rooms

- M5 the browser sandbox: OpenSSL compiled to WebAssembly, same lessons, no install
- M6 practice targets: deliberately misconfigured TLS endpoints for the server-configuration lessons

## Not yet assigned

- a MUD front end
- translations
