# Preface

This file defines how to read and maintain knowledge about the document
standards in `documents/`. Paths such as `.kb/` are relative to that
directory.

IMPORTANT: Read the `.kb/agents.md` file.


# Laws

## §1 - Scope and Organisation

Keep maintained knowledge about document standards in `documents/.kb/`.
Do not apply these conventions to the document standards themselves or to
the documents consumers produce from them.

Use dashed lowercase Markdown filenames that are unique among knowledge
base documents in the repository.

## §2 - Headers and References

Use only these top-level headings, in this order:

- **Preface:** required; explain the document's scope and relevance. End
  with: "IMPORTANT: Read the `.kb/agents.md` file."
- **Laws:** optional; terse constraints in `## §<N> - <Title>` subsections.
- **Concepts:** optional; explanations, design details, and references.

Omit empty or trivial optional sections. Leave two empty lines before
top-level headings and one before subheadings.

Keep law numbers stable. Never change a law's subject; if obsolete, retain
its number with the heading `## §<N> - Revoked`.

## §3 - Maintenance

Record current, durable understanding rather than task logs, plans, or
conversation transcripts. Update relevant knowledge when it changes.
Link to authoritative document standards instead of duplicating them.

Do not edit this convention unless explicitly requested. List additional
knowledge-base files under Concepts in this file when explicitly
authorised to update it. Report missing guidance rather than silently
assuming it does not apply.


# Concepts

## Knowledge Index

- [Document standards](document-standards.md): purpose, boundaries, usage,
  and maintenance of the standards in `documents/`.
