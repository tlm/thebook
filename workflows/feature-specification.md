---
name: feature-specification
description: >-
  Collaboratively explore a feature within a product and define its
  user-facing behaviour and contract. Use when developing a feature idea
  into a lightweight specification, before technical implementation planning.
---

# Feature Specification

## Operating Context

Before executing this workflow, read and follow
[Workflow operating context](AGENTS.md).

## Goal

Develop a shared understanding of a feature within a product and capture
how it will work for the product's users. The specification defines the
feature's public-facing contract, not the internal implementation.

## Approach

- Start with the problem, who experiences it, and what they need to achieve.
- Explore ideas with the user. Research existing products, capabilities, or
  alternatives when it helps resolve a question; do not research by default.
- Work through concrete examples of how someone will use the product.
  Describe interactions, UX elements, requirements on users, and observable
  outcomes, including important failure cases.
- Identify scope, constraints, and unresolved questions. Ask rather than
  silently turning assumptions into requirements.
- Refine a lightweight specification with the user as decisions become
  clear. Keep research out unless it is needed to understand the contract.

## Steps

### Research Framing

**Name:** `research-framing`

**Purpose:** identify what needs to be learned to inform the feature and
turn the initial problem statement into bounded research briefs.

**Instructions:**

- Extract the problem, intended outcome, and research leads from the user's
  prompt. Identify additional areas that could materially inform the
  feature, distinguishing user-provided leads from agent suggestions.
- Keep each brief focused on one primary question that can be investigated
  independently within a single subagent context. Its result should inform
  a specific aspect of the feature, not survey an entire domain.
- For each brief, state the question, its relevance to the feature, what is
  in and out of scope, and the expected result. Include enough context for
  another agent to investigate it independently.
- Break broad areas into multiple narrow assessments intended for separate
  subagents using appropriate research skills. Avoid unnecessary overlap;
  prefer the smallest set of assessments that covers the relevant questions,
  rather than creating assignments merely to increase parallelism.
- Suggest additions or ask questions when useful. Resolve ambiguities that
  would materially change an assignment, but do not require user approval
  of the briefs. Ready briefs may pass to research execution while
  questions about other briefs remain open.
- When revisiting framing, read existing briefs first and add or refine
  only what is needed.
- Record prepared briefs in working files using the shared naming
  convention. Link them from the workflow record's `research-framing`
  section. Write additions or revisions to new working files, preserving
  earlier versions. Identify which earlier material a revision replaces.
- Keep delegation and research completion out of the framing section;
  those belong to research execution. If no research is needed, record
  that outcome rather than inventing assignments.

**Outcome:** new or refined research briefs based on the information
currently available, with unresolved questions recorded, or a record that
no research is currently needed. Framing may be revisited as findings or
other context become available; it does not need to finish before research
execution can proceed.

## Output

A concise specification describing:

- The problem and intended users.
- The proposed solution and its scope.
- User journeys and examples showing the public-facing behaviour.
- Requirements on users and guarantees the product provides.
- Any unresolved questions that prevent agreement on that behaviour.

Use only the structure and detail the feature needs. Do not include internal
architecture, implementation plans, or code-level design. User-visible
technical details belong here only when they are part of the contract.

Review the specification with the user. Distinguish agreed behaviour from
open questions, and do not proceed to implementation as part of this
workflow.
