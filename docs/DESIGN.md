# openssl-workshop — design

> **OpenSSL Workshop**: lessons generated from OpenSSL's own manuals for the release you pick, graded on the
> real artifacts you make, re-run the night the manuals change, published as a site. One Go binary `workshop`;
> lessons are files. The corpus, the release images, the checks and the host probe come from
> [`cryptographers-core`](https://github.com/VeronAImars/cryptographers-core); its `docs/DESIGN.md` is the
> reference for every shared piece and is cited here by section.
>
> **Status: v1 (2026-11-07), pre-spike.** ⬜ marks a claim the M0 spike confirms.

## 1. Context

The maintainer's 2025 module for cybersecurity students (archived) taught OpenSSL 1.1.1 (end of life
2023-09-11), a chain check that passes with a missing intermediate, and an RSA-only key match, behind a fake
terminal. Its link is still circulating. The OpenSSL project asked the maintainer to pursue teaching material
built on the current manuals. Verified facts about the manuals and their rate of change: core design §1.

The gap this closes is the one between what is taught and what operators meet: a learner is told "run
`openssl`" on a Mac whose `/usr/bin/openssl` is LibreSSL while Python, Node and curl each link something else
(core design §2). Teaching one command in one version, from a fixed binary, and then showing the learner the
stack they actually have, is the whole idea.

## 2. The core constraint

**A lesson is only as correct as the release it names, and only as honest as the binary that graded it.**

- Every step cites a page and section in the corpus for the releases the lesson declares. A step with no
  citation is not merged (`CONTRIBUTING.md`).
- Every check grades an artifact the learner made by running the real `openssl` in the pinned image for each
  declared release (core §5.3). The grader never reads what the learner typed.
- The page shows the output captured by the run that graded it, never a pasted transcript.
- When a cited section changes upstream, the lesson is re-run that night; a failing lesson is withdrawn from the
  site until fixed (M2).
- The model drafts and explains; the checks grade; a person confirms before anything is published (M3).

## 3. Architecture

```text
lessons/<module>/<nn>-<slug>/lesson.yaml ──▶ workshop lesson check ──▶ core pkg/check (sandbox, pinned image) ──▶ findings
                                           └──▶ workshop site build ──(findings + captured output)──▶ site/ ──▶ GitHub Pages
core index.json ──▶ workshop lesson draft <page> ──(model drafts, cites)──▶ pull request ──(person confirms)──▶ lessons/
core nightly changed-section list ──▶ workshop drift ──▶ re-run lessons that cite it ──▶ one issue per failure
core host probe ──▶ the "stack you actually have" module (lessons that read the learner's own machine)
workshop ask / shell completion ──▶ core index for the release `openssl version` reports ──▶ cited answer
```

### Components

| Component | Responsibility | Path |
|---|---|---|
| CLI | `lesson check\|run\|draft`, `site build`, `drift`, `ask`, `completion` | `cmd/workshop/` |
| Lesson | schema for `lesson.yaml`; map steps to core checks; declared versions | `internal/lesson/`, `lessons/` |
| Runner | run a lesson's checks through the core for each version; collect findings | `internal/runner/` |
| Site | render lessons with captured output to static pages | `internal/site/`, `site/` |
| Generator | draft a lesson from a page's DESCRIPTION, OPTIONS, EXAMPLES; every step cited | `internal/draft/` |
| Drift | take the core's changed-section list; find citing lessons; re-run; open issues | `internal/drift/`, `.github/workflows/nightly.yml` |
| Tutor | release-aware completion for `openssl` itself; `ask` answering from the index with the section | `internal/tutor/` |

## 4. Platform realities and the modules that teach them

The workshop grades in the core's pinned images, so a lesson's result never depends on the learner's
machine (core design §2). The learner's machine is still the point: core design §4 measured five TLS stacks
on one Mac, and that gap between what is taught and what runs is what the second module teaches from.

Lessons are grouped in modules. The first two:

- **`pki-basics`**: key (`genpkey`, `pkey`), request (`req`), chain (`x509`, `verify -untrusted`), revocation
  (`ca`, `crl`, `verify -crl_check`), server (`s_client`). Each lesson's artifact and checks are in its issue.
- **`your-stack`** (answers the conflation between what is taught and what runs): the learner runs the core's
  host probe on their own machine and the lesson reads the result. Which `openssl` is on `PATH`; what their
  Python, Node and curl actually link; what `OPENSSL_CONF` would change; why `/usr/bin/openssl` on a Mac is not
  OpenSSL. The checks grade the probe output, so the lesson is true for every learner's machine.

## 5. Schemas

### 5.1 Lesson

```yaml
id: pki-basics/03-chain
title: Verify a leaf through an intermediate
versions: ["3.5", "4.1"]
cites:
  - {page: openssl-verify, section: OPTIONS, option: -CAfile}
  - {page: openssl-verify, section: OPTIONS, option: -untrusted}
goal: Verify leaf.pem against root.pem, supplying the intermediate.
setup:                 # run in the sandbox before the learner's step; builds the fixture
  - openssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:P-256 -out root.key
  - ...
steps:
  - say: The verifier builds a path from the leaf to a trusted root. It needs every link.
    cite: {page: openssl-verify, section: DESCRIPTION}
checks:                # the core's check schema, or a reference to a core rule
  - rule: correctness/chain
```

`checks` either names a core rule (`rule:`) or inlines commands in the core's schema (core §5.3). A lesson with
both `setup` and `checks` is complete; `steps` are prose for the page and the tutor.

### 5.2 Page data

`workshop site build` writes one JSON per lesson per version beside the HTML, from the core's findings (core
§5.4) plus the lesson's prose. The page is a function of that file; nothing on it is hand-written after grading.

## 6. Repositories

Core design §6. This repository imports the core as a Go module and pins a release; it never holds a copy of
the manuals.

## 7. Build roadmap as slices

- **M0** one lesson graded on two releases through the core; its page on GitHub Pages; the earlier module's link
  redirected (301).
- **M1** the `pki-basics` module, five lessons.
- **M2** drift: nightly re-run of lessons that cite a changed section; one issue per failure.
- **M3** the generator: `lesson draft <page>` opens a pull request a person confirms.
- **M4** the tutor: release-aware completion for `openssl` itself, then `ask`.
- **M5** the browser sandbox (OpenSSL compiled to WebAssembly); **M6** practice targets.
- `your-stack` module: after M0 and the core's host probe land; its own issue set.

## 8. Risks and mitigations

| Risk | Verified how | Mitigation |
|---|---|---|
| A lesson passes on one release and fails on another | ✅ the chain-check case (core §8) | every lesson lists versions; CI runs each; the page shows each |
| Learners have no container runtime | ⚠️ inferred for students on managed laptops | the site shows captured output for reading; M5 removes the need |
| The old link's host is unknown, so the redirect cannot be set | ✅ archived repo has `netlify.toml`, no homepage | track in issue #2; one-file `_redirects` signpost if the host is Netlify, DNS if a custom domain |
| Generated lessons are fluent and wrong | ⬜ M3 measures the confirm-rate | a draft is a pull request; the checks must pass before a person reads it; nothing publishes on its own |

## 9. Verification

```sh
workshop lesson check lessons/pki-basics/03-chain --version 3.5 --version 4.1   # exit 0; remove -untrusted → non-zero, cites the section
workshop site build && test -f site/out/pki-basics/03-chain/index.html
curl -sI <old module URL> | grep -i '^location: https://<pages url>/'          # 301 in place
hack/acceptance/m0.sh
```

## 10. Decisions

- DR-001 independent, clean-room, open source; sole maintainer.
- DR-002 one Go binary; lessons as files; the core does the cryptography.
- DR-003 three repositories, one program; this one carries OpenSSL's name as a dedicated workshop.
