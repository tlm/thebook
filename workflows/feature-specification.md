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

[Research Framing](#research-framing),
[Research Execution](#research-execution), and
[Research Synthesis](#research-synthesis) form an iterative loop, not
sequential phases. Execute a brief only when it is ready, and synthesise
findings when they are available; neither requires all other work in a
step to finish. Synthesis may request corrections or further research
while other assignments continue. Revisit the relevant steps as questions,
evidence, and the feature's needs change.

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
  of the briefs. Ready briefs may pass to
  [Research Execution](#research-execution) while
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
- Start with ready briefs;
  [Research Framing](#research-framing) does not need to finish first.
  Return questions that cannot be investigated within their scope to that
  step.
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
  [Research Framing](#research-framing) for a revised brief. Record the
  reason and relevant links.
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

### Research Synthesis

**Name:** `research-synthesis`

**Purpose:** assess whether research is fit to inform the feature,
integrate usable findings, and identify necessary follow-up.

**Instructions:**

- Read the current position and prior review entries before
  starting. Identify new or revised briefs and findings, outstanding
  concerns, and earlier conclusions affected by the available material.
- Review findings against their briefs and the feature decisions they are
  intended to inform. Assess whether the question was answered, evidence
  supports the claims, and scope and depth are sufficient for that use.
  Check material assumptions, source suitability, relevant dates and
  versions, uncertainty, and conflicting evidence.
- Review proportionately. Inspect underlying sources when consequential
  or questionable claims need checking; do not treat reading a report as
  independent verification or repeat every investigation by default.
- Integrate findings across assignments. Identify agreement, contradiction,
  gaps, and implications for intended users and outcomes. Distinguish
  evidence from inference, proposed choices, and agreed requirements.
- Return inadequate answers to
  [Research Execution](#research-execution) for correction or further
  investigation against the existing brief. Return new, changed, or overly
  broad questions to [Research Framing](#research-framing) for new or
  revised briefs. Record the reason and relevant links; the receiving step
  owns its resulting
  briefs, assignments, and status.
- Bring product choices to the user rather than treating research as a
  decision on their behalf. Record limitations that cannot reasonably be
  resolved and explain how they constrain use of the findings.
- Maintain the workflow record's `research-synthesis` section with a
  `Current position` subheading and numbered `Review 001`, `Review 002`,
  and subsequent review subheadings. Continue numbering across executions.
  Update the current position with usable conclusions and supporting
  findings links, material limitations, and outstanding concerns and
  follow-up.
  Append review entries; preserve earlier assessments and their context.
- In each review entry, record:
  - Exact brief and findings-file versions reviewed, with links, and any
    underlying sources inspected directly. Identify relevant claims or
    sections where needed; do not log every reading action.
  - Prior conclusions and review entries used as context. Attribute
    earlier reviews by step and number, such as
    `research-synthesis / Review 001`.
  - Which findings are usable, usable with limitations, or insufficient for
    their intended use, with concise reasons.
  - Which findings support each synthesis conclusion. Distinguish material
    reviewed from material actually used to support conclusions.
  - Follow-up requested, unresolved concerns, and links to any resulting
    briefs or assignments once available. A request is not a resolution.
  - Earlier concerns or conclusions resolved, qualified, contradicted,
    superseded, or carried forward, with references to their prior entries.
  - Changes made to the current position, with supporting findings links,
    or a statement that no change was needed and why.
- Reconcile new evidence with affected earlier conclusions, not just newly
  arrived files. Check whether follow-up resolves the original concern.
  Reuse prior assessments when unaffected; reconsider them when briefs,
  evidence, intended use, or source limitations materially change.
- Keep ongoing synthesis conclusions in the current position and attribute
  changes through the numbered reviews. Do not compile a standalone
  research document during each review;
  [Research Consolidation](#research-consolidation) owns that document.
- Do not rewrite research reports or briefs in synthesis. Route corrections
  through the responsible step, preserving earlier artefacts. Do not
  maintain other steps' execution state in the synthesis section.
- Work with partial results when useful. Stop requesting research when the
  evidence is sufficient for the current feature questions and material
  limitations are understood, not when no further questions can be imagined.

**Outcome:** a current research position identifying what can inform the
feature now, material limitations, choices requiring user discussion, and
research follow-up needed. The current position and review history allow
later executions to reconcile new evidence without losing earlier
assessments or mistaking outstanding work for resolved concerns.

### Research Consolidation

**Name:** `research-consolidation`

**Purpose:** compile reviewed research into a coherent, standalone
research document for subsequent feature exploration.

**Instructions:**

- Begin when [Research Synthesis](#research-synthesis) establishes that
  research is sufficient for the
  current feature questions and no outstanding research work is necessary
  before proceeding. Record remaining limitations explicitly; do not
  require every conceivable question to be answered.
- Read the current research position, relevant review entries, supporting
  findings, and any existing consolidation. Do not treat unreviewed results
  as established findings or silently resolve research concerns here.
- Read and follow the
  [Research Briefing Guide](../documents/research-briefing.md).
  Use its document structure, editorial standard, and completion review to
  produce a human-first briefing that supports discussion and direction.
  Agents should be able to get their bearings and follow links for deeper
  context. Research does not establish agreed feature requirements.
- Link substantive findings to their supporting research and sources.
  Explain the research coherently without requiring readers to reconstruct
  the review history. Keep assignment status and review-process narration
  in the workflow record rather than the research document.
- Write the document using the shared working-file naming convention, with
  `research-consolidation` as the step component. Preserve earlier versions
  when revising it and identify which document the new version supersedes.
- In the workflow record's `research-consolidation` section, record the
  scope, synthesis review entries and findings used, why research is ready
  to consolidate, and a link to the document produced. Identify the current
  version and retain links to earlier versions. Record the completion
  review outcome and any unresolved issues; do not mark consolidation
  complete until the briefing meets the guide.
- If compilation reveals a material gap or contradiction, return it to
  [Research Synthesis](#research-synthesis) for review and follow-up rather
  than silently filling it.
  Record whether the issue blocks consolidation or limits the document's
  use. If later evidence or scope changes undermine a published conclusion,
  mark the affected document as needing revision in the consolidation
  record, return to the research loop, and consolidate again when ready.
- If no research was needed, record that outcome without creating an empty
  research document.

**Outcome:** a source-linked research document representing the reviewed
research for the current feature scope, with limitations explicit and its
inputs and current version recorded, or a record that consolidation was
blocked or unnecessary.

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
