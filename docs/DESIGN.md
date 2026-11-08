# OpenSSL Workshop — design

> Status: outline. Filled after the M0 spike (`docs/findings-m0.md`). Subsystem designs go in
> `docs/design/`, one file each, reviewed by pull request before the code lands.

## 1. Context

<!-- what exists, what is wrong with it, with a URL per claim -->

## 2. The core constraint

<!-- the one thing the design must respect; everything else bends to it -->

## 3. Architecture

<!-- components and the data between them; one diagram at most -->

## 4. Components

| Component | Responsibility | Package or path |
|---|---|---|

## 5. Repositories

Three public repositories, one program (DR-003): `cryptographers-core` (corpus, diff, images,
checks, Go API), `openssl-workshop` (this one: lessons, generator, site, drift re-run) and
`cryptographers-terminal` (the inspection workbench). This repository imports the core as a Go
module and pins a release; it never holds a copy of the manuals.

## 6. Build roadmap as slices

<!-- the phases in docs/PHASES.md, one line each, why in this order -->

## 7. Risks and mitigations

| Risk | Verified how | Mitigation |
|---|---|---|

## 8. Verification

<!-- how we know it works: the proving commands, the fixtures -->

## 9. Decisions

<!-- one line per DR in docs/decisions/, by id -->
