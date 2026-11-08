# Lessons

One directory per lesson: `lessons/<module>/<nn>-<slug>/lesson.yaml`, plus any fixture files.

A lesson declares the versions it runs on, the manual sections it cites, the goal in one
sentence, and its checks. A check runs a real OpenSSL command inside the pinned container and
compares the result; it reads the artifacts the learner made, never the text they typed.

Schema (draft, fixed in M0.1):

```yaml
id: <module>/<nn>-<slug>
versions: ["3.5", "4.1"]
cites:
  - page: openssl-verify        # man1 page name
    section: OPTIONS
    option: -untrusted           # optional
goal: one sentence
setup:                           # commands run before the learner's step, in the sandbox
  - ...
checks:
  - run: openssl verify -CAfile root.pem -untrusted int.pem leaf.pem
    expect: "leaf.pem: OK"
  - run: openssl verify -CAfile root.pem leaf.pem
    expect_fail: true
```

`expect` is a substring of stdout; `expect_fail` asserts a non-zero exit. Richer artifact
assertions (subject, key type, extensions) arrive with M1.
