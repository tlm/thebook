# Preface

This document guides contributors creating or modifying reusable workflow
definitions. It describes design properties to consider, not additional
instructions consumers must load when executing a workflow.

IMPORTANT: Read the `.kb/agents.md` file.


# Laws

## §1 - Contracts and Scope

Follow [Workflow Operating Context](../CONTEXT.md) for the shared execution
and storage contract. Keep workflow-specific instructions in the workflow
definition and execution state in the work-item record. Do not duplicate
shared rules or introduce conflicting storage and naming conventions.

Define each step's stable name, purpose, instructions, and outcome. State
actual prerequisites and authority boundaries explicitly; list order alone
must not imply a rigid sequence. Support revisiting affected activities
without assuming that exploration or local agreement authorises execution.

## §2 - Resumption and Selective Loading

Design a path for a fresh agent to find the goal, current user direction,
active artefact versions, outstanding work, and next action without relying
on conversation history or reading every historical file.

Specify what to read first and when to load further detail. Keep current
positions concise, link to supporting artefacts, and read prior rounds or
superseded versions when a decision, conflict, or changed input requires
it. Avoid unconditional instructions to reload all history on resumption.
Do not make unsupported claims that a workflow resumes efficiently merely
because it stores sufficient history.

## §3 - Current State and History

Keep the current operational position distinct from preserved review and
discussion history. A current position identifies the live state; history
explains decisions, revisions, and provenance. Neither should duplicate
entire findings or conversation transcripts.

Make current versions and superseded material distinguishable. When
revising working artefacts, preserve earlier versions using the shared
working-file naming convention and identify what each revision supersedes.
Update current references when outcomes change. Do not introduce a second
authoritative status document merely to provide a handover summary.

## §4 - Ownership and Reconciliation

Each step owns its state and links to the inputs it consumes and outputs
it produces. The coordinator maintains shared records. Cross-step requests
must not create competing assignment or completion records.

When evidence, proposals, or user direction change, identify affected
versions, assumptions, and conclusions. Record whether follow-up resolved
the original concern; requesting work is not completion. Revisit affected
parts while continuing independent work. Do not silently carry an earlier
approval across newly introduced behaviour or treat stale assessments as
current.

## §5 - Bounded Delegation

Consider delegating independent research, strategy assessments, interface
reviews, comparisons, and document composition when their inputs and
expected outputs can be bounded. Delegation is not a requirement to split
every activity or a substitute for managing the coordinator's context.

Supply selected current inputs, a focused question, relevant guidance,
expected output, and action boundaries. Do not assume inherited context or
file-writing capability. Define delivery and persistence responsibilities
before assigning work, and verify that returned results were retained.

Keep user direction, prioritisation, reconciliation, and agreement with
the coordinator. A delegated recommendation does not establish a product
decision. Return concise conclusions and limitations with links to detail;
do not impose arbitrary brevity that removes evidence needed for the task.

## §6 - Proportionate Process and Validation

Every required record, review, and output should serve a clear continuation,
decision, evidence, or delivery need. Avoid mandatory empty artefacts,
duplicate summaries, and bookkeeping that outweighs the work itself.

Evaluate context cost as well as recoverability. Test important workflow
properties through representative executions, and distinguish intended
design from demonstrated behaviour. Record limitations rather than adding
rules solely because one execution failed to follow existing guidance.


# Concepts

## Resumption Evaluation

A cold-resumption trial gives a fresh agent the harness entry point, a
substantial work-item location, and a bounded continuation request. Check
whether it identifies current proposals, settled and superseded decisions,
outstanding assignments, unreconciled consequences, and a useful next
interaction without making new decisions on the user's behalf.

Assess what it had to load as well as the correctness of its answer.
Reconstructing the entire history demonstrates recoverability, not efficient
resumption. Self-reported reading is not an independent execution trace.
A small successful trial is evidence, not proof of general reliability.

## Current Validation Status

The feature-specification definition requires current positions, numbered
reviews and discussion rounds, and versioned artefacts. These instructions
establish intended record-keeping, not demonstrated execution behaviour.
No cold-resumption trial of a substantial feature work item is recorded
here. Efficient resumption therefore remains unverified. The properties
above are authoring expectations, not a claim that every existing workflow
already satisfies them.
