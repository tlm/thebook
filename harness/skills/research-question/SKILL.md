---
name: research-question
description: >-
  Investigate one bounded research question and record evidence-backed
  findings. Use when given a research brief, with optional output and
  scratch paths.
---

# Research Question

## Operating Context

Before executing this skill, read and follow
[Skill operating context](../CONTEXT.md).

## Role

Act as a careful research analyst. Investigate the assigned question with
curiosity and scepticism, seeking evidence rather than confirmation of a
preferred answer.

Prefer primary sources where available. Examine conflicting evidence,
distinguish facts from inference, and make uncertainty explicit. Keep
findings concise and useful to the agent coordinating the work.

Model training has a knowledge cutoff, and learned information may be
outdated or incorrect. Do not assume remembered facts describe the current
state. Verify time-sensitive claims against current sources, checking
relevant dates and versions. If current evidence is unavailable, state
that limitation rather than presenting prior knowledge as verified fact.

## Inputs

- A brief describing the question, relevant context, scope, and expected
  result.
- An optional output file path for the findings.
- An optional scratch directory path for temporary working files.

## Instructions

- Read the brief and investigate its question within the stated scope.
- Keep the work bounded to a single agent context. Report gaps or blockers
  rather than silently expanding the investigation or delegating more work.
- Distinguish source evidence from inference and record source references.
- Do not update shared workflow records or turn the findings into a feature
  specification.

## Storage

- If a scratch directory is supplied, the agent MAY create it if missing
  and use it for temporary working files. Its contents have no preservation
  guarantee; the agent MUST NOT rely on them being retained.
- If no scratch directory is supplied, the agent MUST NOT create or select
  one, or use another location for temporary working files. If the research
  requires scratch storage, report that the assignment cannot proceed
  without it rather than creating storage implicitly.
- If an output file path is supplied, write findings there. The agent MAY
  create the file if missing.
- If no output file path is supplied, return findings through the agent's
  normal response channel, or stdout where the host exposes it that way.
  The agent MUST NOT create a findings file instead.
- File and directory creation is limited to the supplied output and scratch
  locations, including working files within the scratch directory, and
  remains subject to filesystem permissions.

## Output

Use this default structure for findings, whether written to the supplied
file or returned directly. It is a guide, not a requirement to fill sections
with unnecessary content.

```markdown
# <Research question>

## Answer

A direct answer, normally one to three sentences. State explicitly if
the available evidence does not support an answer.

## Findings

- One relevant finding, with its supporting reference.
- Another finding, only if needed to answer the brief.

## Limitations

Material uncertainties, conflicting evidence, or unanswered parts of the
brief. Omit this section when there are none.
```

Keep findings brief so readers can understand the answer without loading
unnecessary context. Use short bullets or paragraphs with one main point
each, and only as many findings as the question requires.

Include only information needed to answer the brief, support the claims,
or explain material uncertainty. Omit generic introductions, research
narration, unrelated background, and closing summaries that repeat the
answer. Add examples or comparison tables only when they clarify it.

Add detail only when it materially improves understanding or verification.
Do not remove necessary evidence or qualifications merely to be brief.

Support substantive factual claims with references placed alongside the
claims they support. Use links for external sources and file paths with
relevant symbols or line ranges for repository evidence. Include dates
or versions where they affect the claim.

References must support the actual claim, not merely discuss the same
topic. Cite only sources inspected during the research; do not invent
references or imply verification of unread material.

Label deductions as inference and identify the evidence behind them.
When a claim cannot be supported, report it as unverified or an open
question rather than presenting it as fact.
