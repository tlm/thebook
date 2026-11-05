# Go

These rules apply when writing or modifying Go code.

For Go design, follow [design.md](design.md).
For Go code style, follow [style.md](style.md).
For Go testing, follow [testing.md](testing.md).
For Go documentation, follow [documentation.md](documentation.md).

## General


- Do not edit generated output by hand. Use the repository's documented
  regeneration process and review the resulting diff. Do not assume that
  `go generate` is the correct command.
- If the correct regeneration process cannot be established or run,
  report the blocker rather than patching generated output manually.


