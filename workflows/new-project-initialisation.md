---
name: new-project-initialisation
description: >-
  Collaboratively establish a new project and its working conventions.
  Use when the user asks to initialise a project, not merely when project
  guidance or a knowledge base is absent. Includes optional Canonical
  knowledge-base setup; does not provide a complete project scaffold.
---

# New Project Initialisation

## Operating Context

Before executing this workflow, read and follow
[Workflow operating context](AGENTS.md).

## Goal

Establish the context and agreed working conventions for a new project,
including an optional knowledge base. Preserve existing files and user
decisions. This workflow does not select a technology stack, generate a
complete scaffold, specify features, or authorise implementation.

## Approach

- Keep setup proportionate to the user's request. Use supplied information
  and ask only about decisions needed to proceed safely.
- Reuse the selected work item and existing workflow record under the
  shared operating context. Keep the work-item location distinct from the
  target project root, even when the work item is stored inside it.
- Preserve decisions and results across executions. Revisit a recorded
  choice only when the user requests it or changed context warrants it;
  explain the reason rather than repeatedly offering declined setup.

## Steps

### Establish Project Context

**Name:** `establish-project-context`

**Purpose:** identify the project, target location, and authorised setup
scope before changing project files.

**Instructions:**

- Read existing work-item context and user instructions. Establish the
  project's purpose, target root, and requested setup scope. Do not conduct
  a full product-specification interview or invent technology choices.
- Confirm that the task concerns new-project initialisation. Missing
  `AGENTS.md` or `.kb/` files alone do not establish that a project is new.
- Inspect existing files and applicable project guidance within permitted
  access. Record what must be preserved and identify conflicting knowledge
  conventions or instructions. Do not initialise an unrelated directory.
- Ask for a decision where the target or scope is ambiguous. Do not migrate
  an existing project under this workflow; report the mismatch and agree
  a separate next action.
- Record the project root, purpose, agreed scope, relevant existing content,
  and unresolved questions in the `establish-project-context` section.
  Link supporting material rather than duplicating it.

**Outcome:** an established target and setup scope, or a recorded blocker
requiring clarification. This context is a prerequisite for
[Knowledge Base Setup](#knowledge-base-setup).

### Knowledge Base Setup

**Name:** `knowledge-base-setup`

**Purpose:** offer optional knowledge-base adoption and perform authorised
initialisation through the dedicated skill.

**Instructions:**

- Read the established project context and prior adoption decisions before
  offering setup. If adoption is already explicitly requested or accepted,
  do not require duplicate approval. Respect recorded decline or deferral
  unless the user changes direction or relevant context changes.
- Explain that the Canonical knowledge-base convention gives agents a
  structured place to maintain current project knowledge. Setup adds
  `.kb/agents.md` and a knowledge-base index at the start of root
  `AGENTS.md`, preserving existing instructions. Workflow progress remains
  in the work item, separately from maintained project knowledge.
- Ask whether to adopt the convention when no decision is established.
  Keep the explanation brief and adoption optional. Do not install files
  merely because this workflow was selected.
- Record acceptance, decline, or deferral and the agreed restrictions in
  the `knowledge-base-setup` section. Decline or deferral does not block
  unrelated setup, and does not authorise partial installation.
- After acceptance, read and execute or delegate
  [Initialise Knowledge Base](../skills/initialise-knowledge-base/SKILL.md).
  Supply the target project root, adoption authorisation, and agreed
  restrictions. A delegated agent must load the skill and its operating
  context; do not assume it inherits this conversation or guidance.
- The coordinator maintains the workflow record. Record the assignment,
  subagent identifier when available, returned outcome, files changed,
  source identity, and validation results. Confirm installed files and
  index links rather than treating an assignment as proof of completion.
- Return conflicts and requests to expand scope to the user. Do not
  silently replace conventions, migrate documentation, or upgrade an
  installed knowledge base. Record partial changes and outstanding work
  accurately; a blocked installation is not completed adoption.
- On resumption, inspect recorded results and actual installation before
  retrying. Do not overwrite existing files or reinstall a convention
  simply because this step is revisited.

**Outcome:** a recorded adoption decision and, when accepted, a validated
installation or explicit blocker. No knowledge-base files are created for
an unaccepted offer.

### Review Initialisation

**Name:** `review-initialisation`

**Purpose:** reconcile setup results and establish what happens next.

**Instructions:**

- Review [Establish Project Context](#establish-project-context) and
  [Knowledge Base Setup](#knowledge-base-setup) records against the actual
  project state. Distinguish completed and validated setup from decisions
  declined or deferred, partial changes, blockers, and work not attempted.
- Record this assessment and the next action in the `review-initialisation`
  section. Do not mark an accepted installation complete while necessary
  validation or installation work remains unresolved.
- Summarise the result for the user, including remaining project setup
  outside this workflow's scope. Establish the next action without
  implying that a complete scaffold, feature specification, or implemented
  product now exists. Do not commit changes unless requested.

**Outcome:** a clear account of project initialisation status and the next
agreed action, with outstanding work preserved for resumption.

## Output

A durable workflow record identifying the project, setup scope, adoption
decision, changes, validation results, and outstanding work. If knowledge
base adoption was accepted and successful, the project has the local
convention and root discovery index; no empty topic documents are required.
