# Go

This entry applies to Go code, tests, APIs, design and documentation.
Use the harness routing convention for the guidance below.

## Project-Specific Instructions

Treat this Go guidance as shared defaults. Explicit, applicable
instructions from the consuming project or component take precedence
where they specify different requirements.

Apply overrides narrowly. Continue following the remaining Go guidance
where it is compatible with those project-specific instructions.

Distinguish explicit project requirements from incidental patterns in
existing code. Where this guidance expressly defers to established
project conventions, follow that direction.

If the applicable requirement or its scope is unclear, explain the
specific conflict briefly and ask for guidance.

## Guidance

| Guidance | Applies to | Read policy |
| --- | --- | --- |
| [Design](design.md) | Go design, APIs or architecture | Conditional |
| [Style](style.md) | Go code style | Conditional |
| [Testing](testing.md) | Go tests or behaviour changes | Conditional |
| [Documentation](documentation.md) | Go docs, comments or APIs | Conditional |
| [Constant sentinel errors](patterns/constant-errors.md) | Defining or changing package-level sentinel errors and error contracts | Conditional |

## General

- Do not edit generated output by hand. Use the repository's documented
  regeneration process and review the resulting diff. Do not assume that
  `go generate` is the correct command.
- If the correct regeneration process cannot be established or run,
  report the blocker rather than patching generated output manually.


