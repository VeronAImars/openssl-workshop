# DR-002 — One Go binary; lessons as files; the core does the cryptography

**Date:** 2026-11-07. **Status:** decided. **Maintainer:** veronaimars.

## Question

What is the workshop written in, where do lessons live, and what runs the cryptography?

## Context

- Learners should not need a language runtime to run a lesson.
- Teaching material must be editable by people who do not write Go.
- The corpus, the release images and the checks already exist once, in the core.

## Decision

- One static Go binary, `workshop`, no cgo. It imports `cryptographers-core` as a Go module pinned
  to a release and never re-implements a corpus or a check.
- Lessons are YAML files validated by a schema in the binary. Every step cites a manual page and
  section; every check names a core check or inlines one in the core's schema.
- The static site is rendered from the lessons and the CI-captured output; it is published to
  GitHub Pages and never committed.
- Milestones are files (`docs/PHASES.md`, `hack/acceptance/`), never GitHub milestones.

## Consequences

- Running a lesson locally needs a container runtime for the core's images; the browser sandbox
  (M5) removes that need for learners.
- The model, when added (M3, M4), drafts and explains. It is never in the grading path.
