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

### Research Execution

**Name:** `research-execution`

**Purpose:** investigate prepared briefs and collect evidence-backed
findings that inform the feature.

**Instructions:**

- Read the prepared briefs and existing research-execution records before
  assigning work. Identify new or revised briefs that need investigation;
  do not repeat completed or outstanding assignments without a reason.
- Start with ready briefs; framing does not need to finish first. Return
  questions that cannot be investigated within their scope to framing.
- Select a skill from the list below that fits each brief. Read its
  instructions before use. If no skill fits, report the gap rather than
  silently substituting an unrelated skill.
- Allocate a findings file using the shared working-file naming convention
  before delegation. Ensure its parent directory exists, but do not
  pre-create the findings file; the researcher creates it when writing
  results. Give each subagent its brief, the selected skill's location,
  the output path, and enough context to work independently.
  Require it to load the skill and its operating context; do not assume
  it inherits the coordinating agent's conversation or guidance.
- Delegate independent briefs in parallel when supported. Keep assignments
  bounded to one primary question each. If delegation is unavailable,
  execute them directly using the same skill and recording conventions.
- Decide whether each assignment needs scratch storage before starting it.
  When needed and permitted, allocate a separate scratch directory for the
  assignment and supply its path explicitly. The researcher may create
  that directory if missing. Otherwise, state that no scratch directory
  is supplied; the researcher must not use temporary storage and must
  report a blocker if it becomes necessary.
- Keep durable findings in the assigned output file, never only in scratch
  storage. Scratch contents have no preservation guarantee.
- Maintain assignment status in the workflow record's `research-execution`
  section, linking each consumed brief and expected findings file. Record
  the skill used and subagent identifier when available. Subagents write
  their assigned findings, not the shared workflow record.
- On receiving a result, confirm the findings exist and address the brief,
  with supporting references and material limitations. Record completion,
  blockers, or failure accurately; an assigned output path is not proof
  that research completed.
- Preserve earlier findings when further investigation is needed. Allocate
  a new attempt for a repeated assignment, or return a changed question to
  framing for a revised brief. Record the reason and relevant links.
- Summarise material implications and unresolved questions in the execution
  section, linking to findings rather than copying them. Keep research
  evidence distinct from agreed feature requirements.
- If there are no ready briefs, record that outcome rather than inventing
  research work.

**Available skills:**

- [Research question](../skills/research-question/SKILL.md): investigate one
  bounded question from a prepared brief and produce concise,
  evidence-backed findings. Use for the research briefs defined above.

Listing a skill does not activate it. Load the selected skill when its
assignment is ready to execute.

**Outcome:** findings linked to their briefs, with assignment status,
material implications, and remaining gaps recorded. Execution may be
revisited as further briefs become ready; research findings do not by
themselves establish agreed feature requirements.

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
