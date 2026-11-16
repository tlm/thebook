# Preface

This document guides contributors creating or modifying reusable skill
definitions in `harness/skills/`. It records authoring conventions, not
instructions consumers load when executing a skill.

IMPORTANT: Read the `.kb/agents.md` file.


# Laws

## §1 - Interaction Formats

A skill may define a conversational format for interacting with the user,
including emoji, such as a question layout. Present such a template as
conversational output, for example in a fenced block, so it is not read
as documentation style.

The Markdown guidance's prohibition on emoji applies to documentation
prose and persisted records, including documents a skill produces. The
The runtime permission lives in the
[Skill operating context](../CONTEXT.md#interaction-formats).

## §2 - Composition

A skill meant to be reused by other skills or workflow steps declares
the settings its callers may supply, each with a default for direct
invocation. Caller settings take precedence. The base skill must not name
or special-case its callers; see harness-structure.md§5. Callers state
their settings where they invoke the skill rather than restating its
procedure.

## §3 - Adapted Material

When a skill adapts third-party material, credit its origin in a short
attribution section in `SKILL.md`. Keep the full licence notice in a
sibling `LICENSE` file stating that it covers that directory only, so the
notice travels with the material without adding to the agent's context.
