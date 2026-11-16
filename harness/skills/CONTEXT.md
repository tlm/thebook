# Skill Operating Context

Read this guidance before executing a skill.

## Guidance

Read the [Agent Harness](../HARNESS.md) and follow its loading
convention, unless it is already loaded. Report missing or unreadable
required guidance rather than silently proceeding without it.

Do not assume a subagent inherits the coordinating agent's conversation or
loaded guidance. Load the required guidance explicitly before starting the
assignment.

These references identify required guidance; they do not grant permissions
or replace host instruction precedence, sandboxing, or approval controls.

## Interaction Formats

A skill may define a format for its conversation with the user, such as
a question layout that uses emoji. Use that format only in conversation.
Documents and records the skill produces follow the harness guidance for
documentation.
