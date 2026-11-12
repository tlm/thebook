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
[Workflow operating context](CONTEXT.md).

## Goal

Develop a shared understanding of a feature within a product and capture
how it will work for the product's users. The specification defines the
feature's public-facing contract, not the internal implementation.
Supporting artefacts may explore system-level technical strategies;
code-level design, implementation planning, and production implementation
remain outside this workflow.

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

The usual progression is research context, then
[UX Exploration](#ux-exploration), then
[Technical Strategy Exploration](#technical-strategy-exploration).
These activities are revisitable, not gates requiring every detail to be
settled. Both exploration steps can request research, and technical choices
can introduce or change UX options. Route consequences to the responsible
step with links to affected versions and assumptions. Reconsider affected
parts rather than restarting the entire workflow, and do not treat earlier
agreement as approval of newly introduced behaviour.

[Interface Contract Definition](#interface-contract-definition) runs
alongside these activities as public boundaries emerge. It records their
semantics, compatibility, security expectations, and operating limits;
it is not an exhaustive prerequisite to starting research or exploration.

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
  pre-create the findings file. Check whether the selected subagent can
  write to that location before choosing delivery instructions.
- If the subagent can write the findings file, supply its output path.
  Otherwise, omit the skill's output path and request findings through the
  subagent's response channel. The coordinator then persists the returned
  findings to the allocated file, preserving references, qualifications,
  and attribution without silently rewriting the research. Do not repeat
  an investigation solely because the researcher could not write a file.
- Give each subagent its brief, the selected skill's location, the chosen
  delivery instructions, and enough context to work independently. Require
  it to load the skill and its operating context; do not assume it inherits
  the coordinating agent's conversation or guidance.
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
  the skill used, delivery method, and subagent identifier when available.
  Record who produced the findings and who persisted them. Only the
  coordinator maintains the shared workflow record.
- On receiving a result, verify that findings are persisted in the
  allocated file, whether written by the researcher or the coordinator.
  Confirm they address the brief, with supporting references and material
  limitations. Record receipt and persistence separately if saving fails;
  retain the result for persistence rather than repeating the research.
  Record completion, blockers, or failure accurately; an assigned output
  path is not proof that research completed.
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

- Start with the current position and linked current briefs and findings.
  Read prior review entries when needed to resolve a decision, conflict,
  changed assumption, or provenance question. Identify new or revised
  inputs, outstanding concerns, and affected earlier conclusions.
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
blocked or unnecessary. Consolidation provides context for
[UX Exploration](#ux-exploration), not specification-ready behaviour.

### Interface Contract Definition

**Name:** `interface-contract-definition`

**Purpose:** identify the feature's public interfaces and progressively
establish contracts consumers can use without relying on undocumented
implementation decisions.

**Instructions:**

- Read the current problem context, available UX proposals, technical
  strategies, relevant research, and prior contract records. Begin as
  interfaces emerge, including during research; consolidation is not a
  prerequisite. Grow and revisit the inventory rather than demanding an
  exhaustive list before other work can proceed.
- Initially identify consumers, public boundaries, existing contracts, and
  familiar product patterns. Develop semantics, examples, security
  behaviour, failures, and limits through
  [UX Exploration](#ux-exploration) and
  [Technical Strategy Exploration](#technical-strategy-exploration).
  Before specification agreement, assess coverage, consistency, and
  consequential unknowns. Placement in this list does not require complete
  contracts before exploration begins.
- Identify interfaces introduced or extended by the feature, and existing
  contracts it uses or must not break. For each, name the human or system
  consumer, purpose, public boundary, and affected journeys. Identify
  authoritative existing contracts and distinguish unchanged dependencies
  from behaviour this feature needs to define or clarify.
- Specify the boundary the consumer actually uses. A public CLI needs its
  command contract, not its internal API, unless that API is independently
  exposed to an identified consumer. Include configuration, protocols,
  events, and persisted formats where they form consumer-facing contracts;
  do not inventory every internal call or implementation detail.
- Examine existing product-family patterns before designing extensions or
  new interfaces. Reuse suitable terminology, syntax, configuration shapes,
  defaults, output conventions, and interaction patterns so consumers can
  transfer prior learning. Check behavioural compatibility separately from
  look-and-feel consistency. Similar presentation must not conceal different
  semantics; explain and discuss justified departures with the user.
- Develop accepted inputs, defaults, validation, operation meanings,
  observable effects, outputs, state transitions, and lifecycle semantics
  as relevant. Distinguish human-readable presentation from stable
  machine-consumable output. Address repetition, concurrency, interruption,
  or partial completion where consequential to consumer expectations.
- Use concrete normal, failure, and recovery examples with
  [UX Exploration](#ux-exploration). Examples illustrate the contract but
  do not replace explicit semantics. Identify gaps along complete journeys,
  including interactions between interfaces, not just isolated operations.
- Define operating limits and consequential classes of supported failure:
  observable responses, side effects, retry safety, and recovery paths.
  Distinguish supported guarantees, explicit non-guarantees, and unresolved
  questions. Do not turn an unanswered question into an unsupported case
  merely because its behaviour has not been designed.
- Treat security as part of each applicable interface contract. Identify
  actors, authority, affected resources, sensitive data, and relevant trust
  boundaries. Develop legitimate-user denial, expiry, revocation, and
  recovery behaviour alongside relevant malicious-input, unauthorised-use,
  impersonation, and compromised-participant scenarios. Define required
  protective behaviour and its limits without prescribing unnecessary
  internal mechanisms or claiming protections are already established.
- Ask [Technical Strategy Exploration](#technical-strategy-exploration)
  to assess whether proposed guarantees and protections can be delivered.
  Route evidence gaps through [Research Framing](#research-framing).
  Return proposed user-facing changes to [UX Exploration](#ux-exploration)
  for discussion. Contract definition records the resulting semantics;
  it does not silently choose product behaviour or maintain other steps'
  assignments and discussion histories.
- Reconcile changed evidence, strategies, and user direction with dependent
  contracts and examples. Identify affected versions, flag stale
  assumptions, and record revisions and reasons. Earlier agreement does
  not approve newly introduced behaviour. Continue unaffected work while
  consequential questions are investigated.
- Maintain the authoritative interface and contract inventory in the
  work-item's `feature-specification.md` workflow record, under
  `interface-contract-definition` / `Current position`. Update it as
  interfaces are discovered, changed, or excluded. Do not keep it only in
  conversation or scattered proposal files. `workflow.log` remains the
  workflow index, not the inventory or review history.
- For each inventory entry, record:
  - The interface, its consumers and relevant journeys, and whether it is
    new, extended, or an existing contract dependency.
  - Why its contract is relevant: the consumer need, journey, compatibility
    obligation, or security boundary that warrants inclusion.
  - Its origin: who identified it and the user direction, exploration,
    research finding, or existing contract that prompted it. Link the
    relevant round, review, or artefact. Use an agent role or assignment
    identifier when known; do not invent attribution or require model
    metadata. State when the origin cannot be established.
  - Scope status: proposed for inclusion, confirmed in scope, excluded, or
    superseded, with the reason and supporting decision reference.
  - Contract status, a link to its current definition or authoritative
    reference, supporting UX and technical artefacts, dependencies, and
    open questions or coverage gaps.
- Distinguish identifying an interface from agreeing its scope or behaviour.
  Agents may add proposed entries to expose gaps; that does not establish
  user approval. Do not count silence as agreement. Preserve excluded or
  superseded entries and their decision history rather than silently
  deleting them.
- Keep numbered `Review 001`, `Review 002`, and subsequent review subheadings
  under `interface-contract-definition`, continuing across executions.
  Keep rationale concise in the current inventory and detailed decision
  history in these reviews.
- Append concise review entries identifying exact inputs and contract
  versions reviewed, coverage gaps, user direction, assessments, changes,
  and cross-step follow-up. Preserve previous conclusions and identify
  which concerns were resolved or carried forward. Keep detailed contracts
  in linked working files using `interface-contract-definition` as the
  filename's step component. Preserve earlier versions and identify what
  revisions supersede; attempt numbers are independent of review numbers.
- Assess readiness against consumer use and coverage of important journeys.
  Unresolved questions block agreement when they affect promised behaviour,
  safety or security, an important journey, or existing-contract
  compatibility. Other uncertainties may remain explicit if they do not
  undermine the offered contract. Record the reason for that assessment;
  do not declare the whole feature ready while consequential interfaces
  remain undefined.

**Outcome:** a versioned inventory and linked public contracts covering
consumer semantics, product consistency, compatibility, security, and
operating limits, with proposed and agreed behaviour distinguished and
remaining gaps explicit. These contracts support specification drafting;
they do not imply implementation or technical verification is complete.

### UX Exploration

**Name:** `ux-exploration`

**Purpose:** collaboratively develop candidate user experiences through
research-informed variations, concrete examples, and repeated discussion
with the user.

**Instructions:**

- Start with the current UX position and linked current proposals, problem
  statement, user direction, and consolidated research. Read prior rounds
  when needed to resolve a decision, conflict, changed assumption, or
  provenance question. Use research to inform the experience, not to limit
  it to existing products or mechanisms. If research was unnecessary, use
  the available context.
- Explore the desired experience before optimising for implementation
  convenience. Respect agreed product constraints and established safety
  or hard external constraints. Distinguish those constraints from current
  mechanism limitations, implementation difficulty, and unverified
  assumptions; the latter do not by themselves justify compromising UX.
- Treat multiple approaches as valid candidates. Present meaningfully
  different experiences or variations with their user-facing benefits and
  trade-offs. Do not make the first or easiest-to-build option the default
  winner, manufacture alternatives, or require a fixed number of options.
  Let the user combine elements, reject options, or redirect exploration.
- Express proposals through concrete examples that convey the look and
  feel of using the feature. Use numbered actions and system responses,
  flows, CLI sessions with output and prompts, configuration fragments,
  screen sketches, or API exchanges as appropriate to the product.
  Abstract pattern descriptions alone are insufficient.
- Show what the user needs to know, what they do, what the product shows or
  does, and what happens next. Expose setup, terminology, choices,
  permissions, ongoing maintenance, and important failure and recovery
  interactions where relevant, rather than showing only a happy path.
- Feed discovered public boundaries and concrete examples into
  [Interface Contract Definition](#interface-contract-definition). Use
  existing product patterns where appropriate, and reconcile contract
  questions with the affected journeys rather than leaving semantics to
  implementation.
- Label candidate commands, configuration formats, interactions, and
  capabilities as illustrative proposals, not existing functionality.
  Distinguish agent suggestions, user preferences, explicitly agreed
  behaviour, and assumptions requiring validation. Continued discussion or
  silence is not approval.
- Iterate with the user through multiple rounds as needed. Present concrete
  proposals, gather feedback, revise examples, and explain meaningful
  changes. Preserve alternatives until deliberately set aside and record
  why. Do not turn discussion into a one-shot approval gate or require
  every detail to be agreed before independent exploration continues.
- Treat agreement on one aspect as progress within exploration, not
  completion of the step. After incorporating and recording feedback,
  select the next logically related unresolved topic using the current
  position and user goals. When context is sufficient, continue in the
  same response with a concrete proposal or meaningful variations and
  invite feedback. Merely naming the next topic or offering to continue
  is insufficient; do not require the user to request each transition.
- Keep transitions focused on one useful discussion topic. Briefly connect
  it to settled direction without reopening decisions unnecessarily or
  presenting every remaining question. Advancing discussion does not
  authorise choosing an option, declaring the specification complete, or
  beginning implementation.
- Respect requests to pause, change direction, or perform a bounded action
  such as saving or committing work. Ask a focused question when material
  ambiguity prevents a useful proposal. If evidence blocks one topic,
  consider independent topics before pausing for research; do not invent
  assumptions merely to keep the conversation moving.
- When no useful unresolved UX topic remains, summarise the agreed
  candidate experience, remaining assumptions, and feasibility questions,
  and establish the next action with the user. Do not invent more topics
  or imply that technical validation or specification completion follows
  from local agreement.
- Return questions needing evidence to
  [Research Framing](#research-framing), allowing the research loop to
  inform revisions. Record requests and their links without maintaining
  research assignment status here. Continue independent UX work while
  research is outstanding and reconcile returned findings with affected
  proposals and earlier user direction.
- Bring technical assumptions and implementation difficulties to
  [Technical Strategy Exploration](#technical-strategy-exploration).
  Do not claim feasibility without evidence or silently change the desired
  experience to fit current mechanisms. Review technical consequences with
  the user here, including newly enabled interactions as well as proposed
  compromises. Reconcile them with affected UX versions and earlier user
  direction; mark proposals needing reconsideration until reviewed.
  This step does not perform implementation planning or establish a
  completed specification.
- Maintain the workflow record's `ux-exploration` section with a
  `Current position` subheading and numbered `Round 001`, `Round 002`, and
  subsequent round subheadings. Continue numbering across executions.
  Keep the current position linked to the latest proposals and examples,
  user goals and preferences, agreed direction, viable alternatives, open
  choices, research requests, feasibility assumptions, and next action.
  Distinguish settled direction, the topic currently being explored, and
  the next likely topic with a brief reason it follows. Update that next
  topic as user direction or evidence changes; it is not a fixed sequence.
- Append a concise record for each round: the exact proposal versions and
  research discussed, user feedback and direction, changes and reasons,
  decisions, unresolved questions, and follow-up. Reference earlier rounds
  when revising direction. Preserve outcomes, not conversation transcripts,
  so another agent can resume without repeating settled discussion.
- Store proposals and examples in working files using the shared naming
  convention with `ux-exploration` as the step component. Write revisions
  to new attempts, link them from the round record, and identify which
  earlier material they supersede. File attempt numbers are independent
  of round numbers; one round may discuss several artefacts.

**Outcome:** user-shaped candidate experiences expressed through concrete
examples, with meaningful variations, user direction, open choices, and
feasibility assumptions recorded. The current position and round history
support continued exploration and later validation without implying that
the proposals are implemented, technically validated, or a final contract.

### Technical Strategy Exploration

**Name:** `technical-strategy-exploration`

**Purpose:** explore and compare system-level strategies for delivering
candidate user experiences, making their mechanisms, trade-offs, security
obligations, and unresolved feasibility questions explicit.

**Instructions:**

- Start with the current strategy position and linked current strategies,
  UX proposals, user direction, and reviewed research. Read prior rounds
  when needed to resolve a decision, conflict, changed assumption, or
  provenance question. Identify the exact versions and assumptions being
  considered. Start when enough UX context exists for useful exploration;
  do not require every interaction to be settled first.
- Develop meaningfully different strategies where useful, without requiring
  a fixed number or manufacturing alternatives. Compare them against the
  same UX goals, scenarios, and constraints. Explain which interactions
  each supports, constrains, or newly enables.
- Describe how important operations work end to end using system-level
  diagrams, numbered sequences, or concise comparisons. Identify the
  components and external systems involved, their responsibilities, and
  existing mechanisms versus proposed additions. Do not describe an
  unimplemented mechanism as an available capability.
- Identify data ownership, trust boundaries, identity, permissions,
  lifecycle, compatibility, and operational concerns where relevant.
  Explain failure, recovery, and fallback behaviour. State the guarantees
  and engineering obligations a strategy would require; naming isolation
  or authorisation does not establish that those problems are solved.
- Keep exploration at the system-strategy level. Do not map changes to
  functions, packages, or files, produce implementation task breakdowns,
  or write production code. Targeted source inspection may answer a
  decisive capability question, but must not become a broad codebase dive.
  Establish scope and authorisation before executing probes or using
  infrastructure; merely proposing a strategy does not authorise either.
- Distinguish established evidence, candidate designs, assumptions, and
  unresolved risks. Match feasibility claims to what the evidence actually
  supports. Do not present source inspection as runtime verification or
  claim that an unresolved security obligation has been satisfied.
- Explore credible ways to preserve the desired experience before proposing
  compromises. Distinguish current implementation difficulty from genuine
  constraints and unknowns. Discuss benefits, costs, and limitations with
  the user rather than selecting the easiest implementation by default.
- Iterate through concrete strategies and variations with the user. Keep
  agent proposals, user preferences, and explicitly agreed direction
  distinct. Agreement on one aspect is not approval of all its consequences
  or permission to implement. Preserve alternatives until deliberately set
  aside and record why.
- After incorporating and recording feedback, continue with the next
  related unresolved topic when context permits, presenting a concrete
  strategy or meaningful variations in the same response. Do not merely
  offer to continue. Keep transitions focused, respect pauses and bounded
  requests, and ask when material ambiguity prevents useful exploration.
  When no useful topic remains, summarise the current strategies and
  outstanding obligations and establish the next action with the user.
- Send evidence gaps to [Research Framing](#research-framing), identifying
  the bounded question, affected strategy, and why the answer matters.
  Use the existing research loop rather than an untracked parallel
  investigation. Link requests and resulting work; research steps retain
  ownership of briefs, assignments, and evidence review.
- Return user-facing consequences to [UX Exploration](#ux-exploration),
  including new opportunities, approval flows, lifecycle interactions,
  and recovery behaviour as relevant. Technical agreement does not approve
  those interactions on the user's behalf. UX Exploration owns their
  examples and user-facing direction.
- Assess the guarantees and limits proposed in
  [Interface Contract Definition](#interface-contract-definition). Return
  unsupported assumptions and newly discovered public boundaries there;
  technical choices must not silently change a consumer's contract.
- Reconcile returned evidence and changed UX or technical choices with
  affected strategy versions and assumptions. Record which conclusions
  remain supported, change, or need reconsideration. Identify when an
  assessment applies to an older UX version rather than silently treating
  it as current. Continue unaffected work while follow-up is outstanding.
- Maintain the workflow record's `technical-strategy-exploration` section
  with a `Current position` subheading and numbered `Round 001`,
  `Round 002`, and subsequent round subheadings. Continue numbering across
  executions. Link the current UX versions, research, and strategy files;
  distinguish settled direction, viable alternatives, open choices,
  technical and security obligations, research requests, affected UX
  proposals, the current topic, next likely topic, and next action.
- Append concise round entries recording exact inputs and strategy versions
  discussed, user feedback and direction, assessments, changes and reasons,
  cross-step consequences, and follow-up. Reference earlier rounds when
  revising conclusions. Do not maintain other steps' state here or copy
  conversation transcripts into the record.
- Store strategy artefacts using the shared working-file naming convention
  with `technical-strategy-exploration` as the step component. Preserve
  earlier versions, link revisions from their round entries, and identify
  what they supersede. Attempt numbers are independent of round numbers.

**Outcome:** one or more user-discussed, evidence-informed technical
strategies, linked to the candidate UX, with trade-offs, obligations,
assumptions, and outstanding questions explicit. The record supports
continued research and UX refinement; it is not a detailed implementation
plan or a declaration that unresolved feasibility or security concerns
have been resolved.

## Output

A concise specification describing:

- The problem and intended users.
- The proposed solution and its scope.
- User journeys and examples showing the public-facing behaviour.
- Requirements on users and guarantees the product provides.
- Public interface contracts, including relevant existing-contract
  compatibility, product-family consistency, operating limits, supported
  failure and recovery behaviour, and security expectations.
- Any unresolved questions that prevent agreement on that behaviour.

Use only the structure and detail the feature needs. Keep supporting
system-level technical strategies in separate, linked artefacts rather
than embedding internal architecture in the specification. Implementation
plans and code-level design remain outside this workflow. User-visible
technical details belong in the specification when part of the contract.

Review the specification with the user. Distinguish agreed behaviour from
open questions, and do not proceed to implementation as part of this
workflow.
