# Skill Operating Context

Read this guidance before executing a skill. Use the harness routing
convention: read Default entries whenever this entry point is activated,
and follow the routing instructions in linked entry points. Report missing
or unreadable required guidance rather than silently proceeding without it.

## Guidance

- [Security](../security/AGENTS.md)
  - **Applies to:** all skill execution.
  - **Read policy:** Default.
- [Language](../language.md)
  - **Applies to:** written language and spelling.
  - **Read policy:** Default.
- [Judgement](../judgement.md)
  - **Applies to:** all skill execution.
  - **Read policy:** Default.

Do not assume a subagent inherits the coordinating agent's conversation or
loaded guidance. Load the required guidance explicitly before starting the
assignment.

These references identify required guidance; they do not grant permissions
or replace host instruction precedence, sandboxing, or approval controls.
