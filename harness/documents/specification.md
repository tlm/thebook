# Specification Guide

## Purpose and Audience

Produce a concise feature specification that reviewers can approve and
implementers can build from. Readers are expected to know the product and
its surrounding context. The specification is not a whitepaper, a research
report, or an implementation plan.

The specification defines the feature's public-facing behaviour and
contract. It must stand alone when exported from the work item: a reader
must not need the workflow record or working files to understand it.

## Document Format

Write the specification in Markdown that converts cleanly to common
document editors, so the user can move it into their own document system
once satisfied. Aim for a close structural match with the user's template,
not pixel-perfect formatting.

- Use headings, paragraphs, lists, simple tables, emphasis, links, and
  fenced code blocks. Avoid raw HTML, footnotes, nested tables, task lists,
  and embedded base64 images.
- Carry the title in the metadata table. Use level-one headings for the
  top-level sections below, matching the template's heading levels. This
  is a deliberate document-specific exception to a single title heading.
- Use level-two and level-three headings for subtopics. Do not make
  headings bold or skip levels for visual effect.
- Specify the language on fenced code blocks when known.
- Do not link to work-item records or working files; those links break on
  export. Link only to resources the reader can reach, such as published
  documentation, repositories, issues, or related specifications.
- Prefer numbered sequences or tables for flows. When a diagram is needed,
  supply its source in a fenced block, such as Mermaid, and tell the user
  it must be rendered before export. Do not assume the target editor
  renders diagram source.
- If the user supplies their own template or an earlier specification,
  follow its section names, metadata fields, and vocabulary in preference
  to the defaults here.

## Document Structure

Use the following sections in this order. Omit optional sections when they
add nothing; do not fill them with boilerplate.

### Metadata Table

Required. Begin the document with a metadata table:

```markdown
| Index | [index] |  |  |
| :---- | :---- | :---- | :---- |
| Title | [title] |  |  |
| **Type** | **Author(s)** | **Status** | **Created** |
| [type] | [author] | [status] | [date] |
|  | **Reviewer(s)** | **Status** | **Date** |
|  | [reviewer] | [review status] | [date] |
```

- Replace bracketed placeholders with values the user supplies. Write
  `TBD` for values not yet known; do not invent them.
- **Index** is a document identifier from the user's numbering scheme.
  Do not allocate one.
- **Type** classifies the specification, for example `Implementation` or
  `Standard`. Use the user's vocabulary.
- **Status** records the document's lifecycle stage. Common stages are
  `Braindump`, `Drafting`, `Pending Review`, and `Approved`. Use the user's
  vocabulary when it differs. Only the user changes status.
- **Author(s)** and **Reviewer(s)** are people named by the user. Add one
  row per reviewer. Do not list the agent as an author or reviewer.
- Use ISO 8601 dates, such as `2026-11-13`, unless the user prefers
  another format.

### Abstract

Required. Summarise the problem and the proposed solution in a couple of
sentences.

### Rationale

Optional. Explain why the specification should be accepted and why the
existing state, if any, is inadequate. Include use cases where they help
define the goal. Omit this section rather than presenting weak reasoning
when the motivation is obvious.

### Specification

Required. Describe the issue and the proposed behaviour in enough detail
to review and implement the feature's contract.

- Open with a short overview of how the section is organised and the scope
  of this iteration. State first-pass limitations, deferred support, and
  hard design constraints where readers would otherwise assume them.
- Organise subtopics by consumer, journey, or area of the contract, not by
  workflow step or research assignment.
- Ground each subtopic in worked examples: configuration fragments,
  commands with output, scripts, or API exchanges. Introduce each example
  briefly and call out what is new in it, such as a new field or command.
- Describe behaviour as named cases with observable outcomes. For each
  case, state the condition, what the consumer observes, such as output,
  exit status, or state, and who decides how to handle it. Cover important
  failure and recovery cases, not only the happy path.
- Include public interface contracts: inputs, defaults, outputs, stable
  machine-consumable formats, compatibility with existing contracts,
  operating limits, and security expectations.
- Describe mechanisms only where consumers observe or depend on them, such
  as an environment variable, socket path, or file format. Keep internal
  architecture, code structure, and implementation tasks out.
- Distinguish guarantees, explicit non-guarantees, and deferred behaviour.
- When an approach has not been settled, present the viable candidates as
  options so the draft can serve as a discussion piece. See
  [Unsettled Approaches](#unsettled-approaches).
- End with an `Open Issues` subsection when questions remain. State each
  issue and why it matters, including each choice between options. Open
  issues must be resolved before approval. Omit the subsection once none
  remain.

#### Unsettled Approaches

- Place options within the subtopic they affect, under headings such as
  `Option A: Invitation links` and `Option B: Invitation codes`. State
  behaviour common to all options once, outside the options.
- Scope each option to the part of the contract that differs. Do not
  duplicate the whole specification for each candidate.
- Give each option the same depth: worked examples, named behaviour cases,
  and contract details sufficient for reviewers to compare them fairly.
- Follow the options with a short comparison of their consumer-facing
  benefits, costs, and limitations. Use a table when the points are
  genuinely comparable.
- State any recommendation and whose it is. A recommendation is not a
  decision.
- Link each set of options to the open issue that records the decision
  needed.
- Keep option labels stable across revisions. Do not reuse a label for a
  different approach.
- Once the user decides, describe the chosen approach as the proposal and
  move set-aside options to Further Information with the reason.

Keep the specification as brief as possible while conveying the contract.
There is no minimum length.

### Further Information

Optional. Explain significant design decisions, alternatives considered,
and related work, such as how other systems handle the issue. Record
options that were set aside and why. Link to related specifications,
repositories, issues, and external information.

Use subsections such as `References` and `Future Considerations` where
they help. Future considerations describe likely later work that this
specification deliberately positions for, without committing to it.

### Spec History and Changelog

Required. Record the document's progress in a table:

```markdown
| Author(s) | Status | Date | Comment |
| :---- | :---- | :---- | :---- |
| [author] | [status] | [date] | [comment] |
```

- Record status changes, review meetings, attendees, and conclusions
  supplied by the user. Do not invent meetings, attendees, or approvals.
- Keep agent drafting revisions in the workflow record, not the changelog,
  unless the user asks otherwise.

## Editorial Standard

- Write in direct, neutral, professional Australian English, unless the
  user's organisation requires another variety.
- Assume the reader knows the product. Explain feature-specific terms once
  and use terminology consistently.
- Keep paragraphs short. Prefer concrete examples and named behaviour cases
  over abstract description.
- Present agreed behaviour as the proposal under review. Do not label
  agreed examples as illustrative.
- Label unsettled approaches as options and keep them distinct from agreed
  behaviour. Raise other unagreed proposals as open issues rather than
  presenting them as the proposal.
- Do not include organisation names, internal links, people, or contact
  details unless the user supplies them for this specification.
- Omit research narration, workflow process, and review history.

## Completion Review

Before presenting a draft for user review, check:

- Does the abstract accurately summarise the body?
- Can a reader who knows the product understand the specification without
  opening the work item?
- Does every behaviour stated in the body trace to agreed direction, with
  no new behaviour introduced during composition?
- Are important failure and recovery cases covered with observable
  outcomes?
- Are public contracts, compatibility, limits, and security expectations
  explicit where consumers depend on them?
- Are guarantees, non-guarantees, deferred work, and open issues clearly
  distinguished?
- Are options clearly labelled, compared fairly at similar depth, and each
  linked to an open issue for the decision?
- Is internal implementation detail excluded unless consumers depend on it?
- Are metadata values supplied by the user or marked `TBD`?
- Will the Markdown convert cleanly: no raw HTML, work-item links, or
  unrendered diagrams presented as final?
- Has unnecessary detail, repetition, and process narration been removed?

Revise composition problems before presenting the draft. Return gaps in
agreed behaviour to the responsible workflow step rather than filling them
during composition.
