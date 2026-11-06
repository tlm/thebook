# Workflow Operating Context

Read this guidance before executing a workflow.

## Guidance

Read Default entries whenever this entry point is activated. Report missing
or unreadable required guidance rather than silently proceeding without it.

- [Judgement](../judgement.md)
  - **Applies to:** all workflow execution.
  - **Read policy:** Default.

## Work Items

A work item represents a goal and holds the durable context and artefacts
accumulated while pursuing it. It is independent of any particular
workflow, agent, or conversation. Multiple workflows and repeated
executions can contribute to the same work item.

## Storage

Each work item has one authoritative storage location. Use dedicated
work-item tooling when available and appropriate; otherwise, represent
the work item as a directory.

Do not move or duplicate an existing work item merely because another
storage mechanism becomes available.

## Selecting or Creating a Work Item

- Reuse the current work item or one explicitly identified by the user,
  including a supplied directory. Inspect existing information before
  writing; do not reinitialise an existing work item.
- When filesystem storage is needed and no work item has been selected or
  supplied, the agent may create `work-item-<goal-slug>` in the current
  working directory, within its existing filesystem permissions.
- The goal slug must be a short, lowercase, hyphen-separated description of
  the goal, not the workflow, agent, or execution. For example:
  `work-item-team-invitations`.
- Before creating it, check for a clearly matching work item in the current
  working directory. Reuse a clear match; ask the user when ambiguous.
- Do not overwrite an existing directory or automatically create a suffixed
  duplicate to avoid a collision.
- Retain the selected location across workflows and executions, even if
  the current working directory changes.

## Filesystem Structure

### Goal

When creating a work-item directory, create `goal.md` at its root. Read
this file first when determining whether an existing work item is relevant.

Keep `goal.md` small: only a broad-strokes description of the goal,
normally one to three sentences. Do not include research, requirements,
plans, progress, decisions, or execution history. Those belong in workflow
records or their linked artefacts.

Agents contributing to the work item may refine the description as
understanding improves, but must keep it brief and preserve the identity
of the goal. Updating the description does not rename the directory.

### Workflow Index

Create `workflow.log` beside `goal.md` if it does not exist. Despite its
extension, this is a Markdown index, not a chronological event log.
Preserve existing entries.

Name each workflow record `<workflow-name>.md`, using the `name` from the
workflow definition's frontmatter. Store it at the work-item root.
Repeated executions reuse the same file.

Each index entry identifies a workflow by its name and links to its record
relative to the work-item directory. For example:

```markdown
# Workflows

- [feature-specification](feature-specification.md)
```

The linked file is the workflow's work-item-local record, not its reusable
harness definition. Before creating a record, check the index for an entry
with the same workflow name. Reuse that record across repeated executions;
do not register a new entry for each execution.

If no entry exists, create a record and add its link to the index. The
workflow defines the record's contents, including its execution context
and progress, and may link to further artefacts. If an existing entry's
record is missing, report it rather than silently replacing lost context.

## Workflow Steps

A step is a named activity within a workflow with a specific purpose and
an expected outcome. It may involve collaboration, delegated assignments,
and multiple context windows.

Steps are not inherently ordered. The coordinating agent selects and
revisits them as the work requires, using available context and outputs
from other steps. Any required ordering or prerequisites must be stated
explicitly by the workflow.

Define each step in the workflow document with these required fields:

- **Name:** a stable, lowercase, hyphen-separated identifier, unique within
  the workflow.
- **Purpose:** what the step is intended to achieve.
- **Instructions:** how to carry out the work.
- **Outcome:** what the step produces or establishes during an invocation.

A step may contain multiple assignments. Use the step's stable name as the
`step` component of working filenames. For example, `research` is a step
and `research-invitation-expiry-001` identifies an assignment within it.

Keep step definitions in the reusable workflow document. Record their
execution state and progress in the work-item-local workflow record.

## Workflow Working Files

A workflow may create `<workflow-name>/` at the work-item root for working
material, using its frontmatter `name`. This directory persists across
context windows and repeated executions; it is not disposable storage.

Name working files `<step>-<topic>-<attempt>.md`:

- `step` identifies the workflow step, such as `research`.
- `topic` is a short, lowercase, hyphen-separated description.
- `attempt` is a sequence starting at `001` for each step and topic.

For example: `research-invitation-expiry-001.md`.

The coordinating agent allocates filenames before delegation, checking
existing files and assignments recorded in the workflow record. Use the
next unused attempt number for a repeated assignment; do not overwrite
earlier findings.

Each subagent writes to its assigned working file. The coordinating agent
maintains the shared workflow record and links to the working files.

## Recording Progress

The coordinating agent records meaningful progress in `<workflow-name>.md`
as it occurs, rather than reconstructing it at the end.

Organise the workflow record into a separate section for each step, using
the step's stable name as the heading. Reuse that section when revisiting
the step.

Each step owns its section. Record its actions, inputs, outputs, and
progress there. A step may read other steps' sections and consume their
outputs, but must not maintain its state in their sections.

Link to the prepared material the step consumes and the working files it
produces. Record enough information to distinguish material already acted
on from newly available inputs when the step resumes.

A producing step records what it prepared, not whether another step has
processed it. The consuming step records that processing in its own
section.

Within each step's section, record:

- **Step started:** which workflow step is underway.
- **Work identified:** the areas or assignments selected.
- **Work delegated:** the assignment name, subagent identifier if available,
  skill used, and a relative link to the expected findings file.
- **Result received:** whether the assignment completed, was blocked, or
  failed, with a brief outcome. Confirm whether findings were produced;
  an expected destination alone does not establish that results exist.
- **Next action:** what should happen next, including any user input needed.

Keep entries factual and concise: actions, outcomes, and outstanding work.
Do not record internal reasoning, every tool call, or copies of findings.
Link to working files for detail. Preserve enough context for a fresh agent
to distinguish completed work from outstanding assignments.

## Durable Context

Store information needed to continue the work in the work item, rather
than relying on conversation history. A fresh agent must be able to
understand the goal and find the relevant context. Preserve existing
information when resuming work.
