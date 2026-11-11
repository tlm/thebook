# Preface

This document defines the harness publication boundary, local build
command, and content-version contract.

IMPORTANT: Read the `.kb/agents.md` file.


# Laws

## §1 - Publication Policy

The publisher includes every regular file under its source directory
unless `excluded` in `internal/publication/build.go` excludes it.
Exclude any source path component starting with `.`, at any depth, and
files named `AGENTS.md` (case-sensitive). This includes hidden directories
such as `.kb/`, `.git/`, and `.github/`, as well as dot-prefixed files.

Before adding non-hidden material under `harness/` that must not be
published, update the hardcoded policy and its tests. Placement inside the
source tree is otherwise a decision to publish. Do not rely on Git ignore
rules or a file being untracked to keep it out of publication. Never put
secrets there.

The generated root paths `index.html` and `version.txt` are reserved;
source collisions cause build failure. Source `.nojekyll` is excluded like
other hidden files; the program generates its own empty marker. Included
symlinks and non-regular files cause build failure rather than being
followed. A future requirement for a public hidden path, such as
`.well-known/`, needs an explicit, tested policy exception. Dot prefixes
are a publication convention, not a security boundary for secret storage.

## §2 - Content Version

The content version is a lowercase SHA-256 digest. It identifies published
paths and bytes, not repository commits, signatures, or semantic versions.
Do not use it as proof of authenticity or permission to load instructions.

The hash input starts with `thebook-publication-v1` and a zero byte. For
each included file in lexicographic slash-relative path order, append its
UTF-8 path length as an unsigned 64-bit big-endian integer, the path bytes,
the content byte length in the same format, and the content bytes.

Include `.nojekyll` and the generated index with an empty version field.
Exclude `version.txt` and the final index's populated version field from
the hash to avoid self-reference. The generated index template therefore
participates in versioning, while contributor changes, timestamps, file
modes, source-root location, and output location do not.

Preserve this deterministic contract or explicitly version its encoding
when changing it. Additions, removals, renames, and published byte changes
must affect the identifier. Excluded edits must not.

## §3 - Build and Validation

Use the local publisher; do not hand-edit its generated output. Keep the
source stable during a build. The destination must be absent, outside the
source tree, and have an existing parent. Existing output is never deleted
or overwritten automatically. A write failure can leave partial new
output; inspect it before removing or retrying it.

Validate links against the assembled bundle, not the contributor source.
Markdown inline, reference, and image links must resolve to published files;
Markdown heading fragments use Goldmark's automatic heading identifiers.
Code examples are not links. Raw HTML links and external URLs are not
validated or fetched. Use relative links, not site-root paths that would
lose the GitHub project-site prefix. Queries do not change the target file.


# Concepts

## Local Commands

Run from the repository root with Go 1.27.4 or later:

```sh
go test -v ./...
go run ./cmd/publish-harness -source harness -output _site
```

The Go module declares the required version. Dependencies are Goldmark for
Markdown parsing and `github.com/juju/tc` for tests. The command prints the
file count, destination, and full content identifier.

`_site/` is ignored by Git. Use a different absent output path for a second
build, or explicitly remove inspected generated output before rebuilding.
The old source `harness/index.md` is replaced by the generated HTML entry.

## Consumer Discovery

The root HTML page links to `HARNESS.md`, explains loading and version
comparison, and displays the identifier. `version.txt` contains the same
64-character digest and a trailing newline for machine retrieval. Other
consumer files retain their source-relative paths and exact bytes.

Comparing identifiers detects a different bundle. A matching identifier
does not imply that an agent has loaded every module applicable to a new
task. Multi-request reads from a moving Pages deployment can span versions;
checking the identifier before and after a load can detect such a change,
but this publisher does not provide immutable versioned URLs.

## Deployment Boundary

The program only assembles a static site.
[Pages deployment](../.github/workflows/pages.yml) runs on pushes to `main`
and manual dispatch. The build job selects Go from `go.mod`, runs tests
and vet, assembles `_site/`, and uploads only that directory. No Jekyll
rendering or repository-root upload is involved.

The upload includes hidden files so the generated `.nojekyll` marker is
retained; source hidden files were already excluded by the publisher.
The build job has read-only repository access. Only the deployment job
receives Pages write and OIDC permissions, uses the `github-pages`
environment, and deploys only `main`. A manual run on another branch can
build but cannot deploy. Pages runs are serialised without cancelling an
in-progress deployment.

Official actions are pinned to commit hashes with release comments. When
updating them, verify their inputs and output contracts, particularly
artifact handling and hidden-file defaults.

In repository Settings, select Pages, then GitHub Actions as the build and
deployment source. Push the workflow to `main` or manually run it there.
The workflow uses the repository's `GITHUB_TOKEN` and OIDC; no personal
access token or deployment secret is required. Environment protection or
organisation policies may require approval. Successful local validation
does not establish that a GitHub deployment has run or that Pages settings
have been changed.
