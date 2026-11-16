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
runtime permission lives in the
[Skill operating context](../CONTEXT.md#interaction-formats).
