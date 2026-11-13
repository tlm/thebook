# Language and Style

## Spelling

- Prefer Australian English in new code, comments, and documentation
  (e.g. "colour", "authorisation"). When an identifier comes from an
  external library that uses American spelling, use the library's name
  verbatim.
- Do not auto-correct existing American spelling to Australian or
  British English. Only switch spelling when you are rewriting the
  surrounding code anyway.

## Compound Words

These apply to code comments as well as documentation.

- Agents SHOULD prefer plain phrasing over compound modifiers. When a
  modifier before a noun would need a hyphen, first try moving the idea
  after the noun:
  - "compatibility with existing contracts", not
    "existing-contract compatibility".
  - "a trial of resuming from cold", not "a cold-resumption trial".
  - "cells holding whole paragraphs", not "paragraph-length cells".
- Agents SHOULD hyphenate only established words, such as "trade-off",
  "read-only", "built-in", and "follow-up" as a noun, or where an open
  compound would be misread. Write the verb open: "follow up".
- Agents SHOULD NOT coin a compound for a single use, or coin compounds
  of three or more words. Established ones such as "end-to-end" are fine.
- Treat endings such as "-specific", "-level", "-facing", "-based",
  "-owned", "-driven", and "-first" as a prompt to rephrase. Keep them
  only when the alternative reads worse.
- Agents SHOULD write a project's established terms open and
  consistently, including before a noun. If a project uses the term
  "work item", write "work item record", not "work-item record".
- Do not just delete a hyphen and leave a stack of nouns. If the phrase
  still reads badly, rewrite the sentence.
- Agents SHOULD NOT rewrite existing text solely to remove hyphens. Apply
  these rules when writing new text or revising the surrounding passage.
