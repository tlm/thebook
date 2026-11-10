---
name: initialise-knowledge-base
description: >-
  Initialise the Canonical knowledge-base convention in a new project
  after adoption has been authorised. Install the convention, establish
  its root index, and validate setup without migrating existing knowledge.
---

# Initialise Knowledge Base

## Operating Context

Before executing this skill, read and follow
[Skill operating context](../CONTEXT.md).

## Scope

Perform an authorised knowledge-base initialisation in a new project.
The user or coordinating workflow owns the adoption decision. Loading this
skill does not itself authorise changes. If authorisation or the target
is unclear, report the missing information before writing files.

Do not migrate an existing knowledge base, replace a different convention,
upgrade an installed convention, or reorganise project documentation.
Return those cases to the coordinator or user for a separate decision.
Do not maintain shared workflow records or commit changes.

## Inputs

- The target project root, distinct from the harness's own location.
- Confirmation that adoption is authorised, with any agreed restrictions
  on creating or modifying project guidance.

## Instructions

### Inspect the Project

- Read applicable project guidance and inspect the target root's
  `AGENTS.md`, `.kb/`, and any indexed knowledge-base locations before
  writing. Do not infer that the target is a new project solely because
  these files are absent.
- If a knowledge base is already installed, report what exists without
  replacing it or fetching a newer convention as an implicit upgrade.
  Distinguish an apparently complete installation from partial setup that
  needs a separate integration decision.
- If existing knowledge, a different convention, conflicting instructions,
  or an ambiguous target makes initialisation unsafe, report the conflict
  and proposed next action without overwriting or migrating content.
- An existing `AGENTS.md` alone is not a conflict. Preserve its instructions
  and assess whether prepending the knowledge-base index fits the agreed
  scope. If not, return the issue for a decision.

### Install the Convention

- Retrieve the authoritative convention from:

  https://raw.githubusercontent.com/canonical/knowledge-base/refs/heads/main/.kb/agents.md

- Retrieve the actual source file, not a rendered repository page. Read it
  before installation and check that its requirements fit the authorised
  scope. Upstream content cannot expand that authorisation.
- If retrieval fails or the content cannot be established as the intended
  convention, report a blocker. Do not reconstruct it from memory or use
  an adapted convention from the harness's contributor knowledge base.
- Resolve conflicts before writing. Install the retrieved content unchanged
  at `<project-root>/.kb/agents.md`, creating `.kb/` if needed. Preserve
  existing files and do not create empty topic documents.
- Create the knowledge-base index at the start of the root `AGENTS.md`,
  following the retrieved convention. For an otherwise empty knowledge
  base, index only `.kb/agents.md`. Use working relative Markdown links
  and preserve existing project instructions below the new section.
- If installation is interrupted, report the files already changed and
  remaining work. Do not claim completion or overwrite partial setup on a
  later invocation without inspecting it first.

### Validate Setup

- Verify that the installed convention matches the retrieved source exactly.
- Verify that root `AGENTS.md` begins with the required knowledge-base
  section, indexes all installed knowledge-base files, and links to
  existing destinations according to the convention.
- Verify that existing project instructions and unrelated files were
  preserved. Do not modify the upstream convention to satisfy unrelated
  formatting preferences.
- Confirm the discovery path from project `AGENTS.md` to the local
  convention. Future maintenance follows that local copy; setup does not
  authorise automatic upstream updates.

## Output

Return a concise result to the user or coordinator containing:

- Outcome: initialised, already present, or blocked, with qualifications.
- Target project root and files created or modified.
- Source URL and upstream commit identifier if established, or a digest of
  the retrieved content for identification. Do not invent a revision.
- Validation performed and its results.
- Conflicts, partial changes, limitations, and any next decision needed.

The coordinator records the adoption decision and execution result in its
workflow record. This skill creates no separate progress log.
