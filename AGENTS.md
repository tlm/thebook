# Introduction

This file is the reusable entry point to the full harness for agents
working on software projects.

Use the routing convention below to discover and read the guidance
required for the current task.

## Loading Convention

Use the same routing convention at every entry point in this harness.
Each guidance entry provides a document link, its applicability, and a read
policy. These policies are instructions for the agent, not standard
`AGENTS.md` metadata or automatic imports. They govern which guidance to
read, not instruction precedence.

- **Default:** read whenever the containing entry point is activated,
  before starting the task or work governed by that entry point.
- **Conditional:** read before work matching the entry's applicability.
  Conditional guidance is mandatory when applicable, not optional advice.
- Determine applicability from the requested work and relevant project
  context. A language's presence in the repository does not by itself make
  its guidance applicable to every task.
- When a linked document is another entry point, evaluate its guidance
  entries using this same convention. Defaults apply within that active
  entry point; they do not require loading every module in the harness.
- Re-evaluate conditional entries as the task develops or its scope changes.

- Report missing or unreadable required guidance rather than silently
  treating it as inapplicable.

## Guidance

| Guidance | Applies to | Read policy |
| --- | --- | --- |
| [Project](project/AGENTS.md) | Work within a consuming project | Default |
| [Security](security/AGENTS.md) | All agent work | Default |
| [Language](language.md) | Written language and spelling | Default |
| [Judgement](judgement.md) | All agent work | Default |
| [Go](go/AGENTS.md) | Go code, tests, APIs, design or docs | Conditional |
| [Markdown](documents/markdown.md) | Writing or modifying Markdown | Conditional |

## Workflows

Workflows are reusable orchestration guides for substantial areas of
development work. They define how work progresses towards an outcome,
including collaboration, decision points, and validation. Workflows may
span multiple context windows and may be combined to achieve a larger
goal.

Each workflow is a Markdown file in `workflows/`, with YAML frontmatter
containing a `name` and `description`, followed by its instructions.

Read a workflow only when the user asks to execute it, or when its
description matches the current task and you judge it suitable. An index
entry alone does not activate a workflow. Once selected, follow its
instructions within the task's scope and the applicable harness guidance.

### [New project initialisation](workflows/new-project-initialisation.md)

- **Applies to:** collaboratively initialise a new project and its working
  conventions, including optional knowledge-base setup.
- **Read policy:** Conditional.

### [Feature specification](workflows/feature-specification.md)

- **Applies to:** collaboratively explore a feature within a product and
  define its user-facing behaviour and contract.
- **Read policy:** Conditional.
