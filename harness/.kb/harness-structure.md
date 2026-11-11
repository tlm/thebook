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

Keep bounded skill instructions in `<skill-name>/SKILL.md`, with the
required discovery frontmatter. Keep individual workflow definitions in
named Markdown files with their own discovery frontmatter.

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

Publication must select consumer artefacts and exclude contributor
`AGENTS.md` files and `.kb/` directories. Validate that published guidance
links do not depend on excluded material. Do not assume that every file
under `harness/` is publishable.


# Concepts

## Source Layout

The reusable product lives under `harness/`. Repository-level material,
including the root contributor entry point, knowledge-base convention,
TODO list, and research, is outside that consumer source tree.

Contributor instructions and knowledge may live alongside consumer files
for locality, but remain distinct by role. The document-standards knowledge
in `harness/documents/.kb/` follows the repository-root convention; it does
not maintain a separate convention copy.

## Publication Status

A selective publisher and its deployment action have not yet been
implemented. The publication boundary above is a requirement for that
work, not a claim that current Pages deployment enforces it. Source moves
and renamed entry points require publication configuration to be updated.
