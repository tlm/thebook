# Preface

This document describes the purpose of `documents/` and how to maintain
its reusable document standards.

IMPORTANT: Read the `.kb/agents.md` file.


# Laws

## §1 - Content Boundaries

Keep reusable document standards in `documents/`. Standards define the
intended audience, structure, editorial expectations, and completion
checks for documents produced by agents.

Keep workflow definitions, skill instructions, and generated work-item
documents in their respective locations, not in this directory. Keep
contributor knowledge about maintaining standards in `documents/.kb/`.

## §2 - Responsibility

The invoking workflow or skill explicitly links to the relevant standard.
A document standard does not execute a workflow or select output storage.
The invoking workflow owns output paths, versioning, progress, and
completion status for its documents.

## §3 - Maintenance

Keep standards independent of a particular workflow where practical.
Preserve consistent required structure while allowing topic-specific
composition. When changing a standard's contract, check its callers for
conflicting instructions and update them where necessary.

Link to authoritative standards rather than duplicating their rules in
knowledge-base documents.


# Concepts

## Markdown Writing

[Markdown Writing](../markdown.md) provides shared source-formatting,
linking, and validation guidance. The root harness routes to it when
writing or modifying Markdown. Document-specific standards retain their
own structure and conventions.

## Research Briefing

The [Research Briefing Guide](../research-briefing.md) defines the standard
for human-first research consolidation. It helps humans understand the
research and provide further direction without reading execution logs.
Agents use the briefing for orientation and follow its links for deeper
research and review context.

The guide is authoritative for the briefing's structure, editorial
standards, traceability, and completion review. The
[Research Consolidation step](../../workflows/feature-specification.md#research-consolidation)
in the feature-specification workflow invokes it and manages the produced
document's lifecycle.
