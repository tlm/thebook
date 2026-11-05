# Go Documentation

These rules apply when writing or updating Go documentation and comments.

- Always document exported constants, types, methods, variables, and
  functions.
- Write a `doc.go` for each new package containing a short summary of the
  package's purpose. Update this summary when the package's purpose
  changes.
- When referring to Go identifiers in comments, use Go doc link notation:
  `[Type]`, `[funcName]`, `[pkg.Type]`. For symbols in another package, use
  its imported name or alias when it is imported in the documenting file.
  Otherwise, use the full import path inside the brackets, such as
  `[net/url.Error]`. Do not add imports solely for documentation links.

## Error Contracts

- When documenting exported or package-visible functions and methods,
  describe expected error conditions that callers can reasonably act on.
  Useful handling includes recovery, alternative behaviour, meaningful
  diagnostics, or translation into the caller's own error contract.
- Include supported error contracts propagated from dependencies hidden
  behind a receiver or implementation. When these failures form part of
  the operation's public contract, callers must not need to inspect its
  internals to discover them.
- Do not routinely repeat error contracts already apparent from the
  signature or established public contracts of supplied dependencies.
  Include them when the operation changes their meaning or handling, or
  when a specific propagation guarantee provides useful information.
- Describe the conditions under which documented errors occur and the
  supported ways to recognise them, including `errors.Is` for sentinel
  errors and `errors.As` for typed errors where appropriate.
- Document supported guarantees, not every error the current implementation
  could produce. Neither a dependency's presence nor an incidental
  propagation path is sufficient to establish a public error contract.
- When meaningful error contracts are documented, include a final section
  in the following form:
    - "The following errors may be expected:"
    - An alphabetical bullet list with each entry in the form
      `[ErrorType] <condition that causes the error>`, using a Go doc link
      to the error type or sentinel, followed by why the operation returns
      it. Apply the package qualification rules above to each link.
- Describe meaningful failure conditions without a named error in prose
  rather than inventing an error identifier for a link.
- Omit the error section when it adds no meaningful caller-facing
  information.

## Purpose and Behavioural Contracts

- Begin documentation by explaining the purpose of the type, function,
  or method: what it represents or accomplishes and why a caller would
  use it. Do not merely restate its name or signature.
- For types, describe their role and any invariants that callers must
  preserve or can rely on.
- For functions and methods, document meaningful preconditions: what
  callers must provide, arrange, or ensure before calling. Include
  required state, valid input constraints, and ordering or lifecycle
  requirements where relevant.
- Do not repeat expectations already clear from the signature or
  ordinary Go conventions. Focus on requirements a reasonable Go
  developer could not infer.
- Document meaningful postconditions: what the operation guarantees on
  success, what state it changes, and any externally observable effects.
  Make clear what callers can rely on after the operation returns.
- When an operation can fail, document whether failure leaves state
  unchanged, permits partial progress, or has other effects that matter
  to callers. Do not imply guarantees the implementation does not make.
- Describe observable behaviour and guarantees, not internal steps.
  Documentation should remain valid when internal composition changes
  without changing the contract.
- Keep documentation proportionate to the contract. Use prose for simple
  cases and separate precondition or postcondition sections when they
  improve clarity; do not add empty or formulaic sections.
