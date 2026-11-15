# Preface

This document explains the harness source layout, entry-point roles, and
separation between contributor instructions and consumer artefacts.

IMPORTANT: Read the `.kb/agents.md` file.


# Laws

## §1 - Audience Boundaries

Use `AGENTS.md` files for instructions to agents developing this repository.
Use `.kb/` directories for maintained contributor knowledge. Neither is a
consumer harness artefact.

Keep contributor guidance out of the consumer loading graph. Contributors
explicitly load the local `harness/HARNESS.md` and follow its default and
applicable conditional guidance. Do not maintain a separate contributor
list of consumer rules that can drift from the harness's own routing.

This dependency is one-way: consumer loading remains independent of this
repository's contributor instructions. If project discovery points back
to contributor guidance already loaded, reuse that context rather than
repeatedly traversing the cycle. Loading the harness does not itself
select a workflow or authorise additional actions.

## §2 - Consumer Entry Points

Use `HARNESS.md` as the root consumer entry point. It defines the loading
convention and routes to guidance and workflows.

Use `GUIDANCE.md` for subject-specific consumer entry points, such as Go,
project guidance, security, and source control.

Use `CONTEXT.md` for shared operating instructions that consumers read
before executing a class of artefacts, currently skills and workflows.
Route their guidance through `HARNESS.md` and its loading convention
rather than keeping separate guidance lists. Partial copies of the
harness defaults drift, and they never reach subagents with conditional
guidance.

Keep bounded skill instructions in `<skill-name>/SKILL.md`, with the
required discovery frontmatter. Keep individual workflow definitions in
named Markdown files with their own discovery frontmatter.

`HARNESS.md` index entries are the pointers agents read when selecting a
workflow or skill, so each must carry its triggers. Keep each frontmatter
`description` the same wording as its index entry, apart from initial
capitalisation. Change both together.

Index skills intended to be invoked directly in a `HARNESS.md` Skills
section, following the workflow index format and stating when to read a
skill. List skills used only within a workflow in the step that uses
them, not in the index. Add the section with its first entry rather than
as an empty placeholder.

Entry-point links instruct consumers what to read; Markdown provides no
automatic imports or instruction authority. Preserve applicability and
read-policy semantics when changing routing.

## §3 - Linking and Changes

Use relative links within the consumer tree so its internal layout can be
preserved in publication. When moving or renaming a file or heading, check
and update affected links and validate destinations and anchors.

Do not rename references to a consuming project's own `AGENTS.md` merely
to match harness entry-point naming. Those references direct agents to
project instructions, not to this repository's contributor entry points.

## §4 - Publication Boundary

Publication must exclude contributor `AGENTS.md` files and every source
file or directory whose name starts with `.`, including `.kb/`.
The publisher uses a hardcoded exclusion policy; all other regular source
files are published by default. Update the policy and tests before adding
other non-consumer material under `harness/`. Validate that published
guidance links do not depend on excluded material.


# Concepts

## Source Layout

The reusable product lives under `harness/`. Repository-level material,
including the root contributor entry point, knowledge-base convention,
TODO list, and research, is outside that consumer source tree.

Contributor instructions and knowledge may live alongside consumer files
for locality, but remain distinct by role. The document-standards knowledge
in `harness/documents/.kb/` follows the repository-root convention; it does
not maintain a separate convention copy.

## Loading Paths

A coordinating agent reaches guidance through `HARNESS.md`. A subagent
reaches it through the operating context of its skill or workflow, which
points back to `HARNESS.md`. Agents that have already loaded the harness
reuse it, so coordinators pay no extra cost. A subagent loads the harness
defaults, roughly 2,500 words, and any conditional guidance its task
triggers.

Delegated work without a skill relies on the delegation rule in the
workflow operating context. Without that rule, such a subagent would load
no guidance at all.

## Publication Status

The Go command `cmd/publish-harness` assembles the consumer site and
generates its root discovery page and content version. See
[Harness publication](../../.kb/harness-publication.md) for the exclusion
policy, hash contract, local commands, and validation limits.

The workflow `.github/workflows/pages.yml` tests and assembles the harness,
then deploys the generated site to Pages from `main`. It requires GitHub
Actions to be selected as the repository's Pages source; branch publication
does not run the publisher or enforce its exclusions. Repository settings
and actual deployment success must be verified separately.
