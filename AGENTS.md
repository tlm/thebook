# Knowledge Base

IMPORTANT: Read the `.kb/agents.md` file.

- [.kb/agents.md](.kb/agents.md) - Structure and rules of the knowledge base.
  You MUST read this file.
- [.kb/harness-publication.md](.kb/harness-publication.md) - Publication
  exclusions, build commands, link validation, and content versioning.
- [harness/.kb/harness-structure.md](harness/.kb/harness-structure.md) -
  Consumer and contributor boundaries, entry-point roles, and publication.
- [harness/documents/.kb/document-standards.md](harness/documents/.kb/document-standards.md) -
  Purpose, responsibilities, and maintenance of reusable document standards.
- [harness/workflows/.kb/workflow-design.md](harness/workflows/.kb/workflow-design.md) -
  Workflow-authoring properties: resumption, context loading, ownership,
  reconciliation, delegation, and proportionate validation.

## Contributor Guidance

References to `.kb/agents.md` in contributor knowledge documents refer to
that file at the repository root, including from nested `.kb/` directories.

This entry point governs development of this repository. It is not the
consumer harness entry point.

Before starting contributor work, read and follow the local
[Agent Harness](harness/HARNESS.md), using its loading convention to load
default and applicable conditional guidance.

Read relevant contributor knowledge before changing the area it describes,
and maintain that knowledge when durable understanding changes.

Use the local harness source, not the published website. Loading the
harness does not by itself select a workflow or authorise additional
actions.

If harness guidance points back to project instructions already loaded,
reuse that context rather than repeatedly traversing the same entry points.
