# Go Design

This document guides agents in designing Go solutions to problems and
developing the architecture as a whole. It applies to new code and, where
required by the task, to changes to existing code. Apply the guidance in
context; it is not a mandate to redesign unrelated code.

## Packages as Cohesive, Testable Components

### Purpose and Boundaries

- Design each package around a clear, cohesive purpose: a capability it
  provides within the system. Treat its public API and documented
  behaviour as the contract for that capability.
- Every type, function, and dependency in a package should contribute to
  its stated purpose and contract. Supporting implementation details
  belong when they serve that purpose.
- When new behaviour introduces a distinct responsibility or stretches
  the package's purpose, consider placing it in a separate package.
  Split along meaningful capability boundaries, not merely to reduce
  file size or isolate individual functions.
- Keep packages composable. Expose what callers need to use the
  capability without requiring knowledge of its internal implementation.
  A clear contract does not require a Go `interface` type in every package.

### Dependencies and Composition

- Make external dependencies explicit and supply them through
  constructors or other wiring functions. A package should ask for the
  capabilities it needs rather than construct unrelated dependencies
  internally.
- Define clear expectations for each dependency, including the behaviour
  and failure conditions the package relies on. Prefer narrow contracts
  expressed in terms of the consuming package's needs.
- Keep construction of unrelated infrastructure and assembly of
  components at explicit composition boundaries. Do not make these
  incidental side effects of constructing or using a package's capability.
- Construct internal implementation details within the package when
  appropriate. Dependency injection is for meaningful external
  boundaries, not a requirement to expose or inject every helper.
- Make resource ownership and cleanup responsibilities explicit. Injecting
  a resource does not by itself transfer ownership; do not close an
  injected resource unless ownership transfer is part of its contract.
- For resources a component owns, define how they are released during
  normal shutdown and on failure, including partially completed
  construction.

### The Package as a Unit of Testing

- Treat the package's cohesive capability as a primary unit under test.
  Verify its contract and observable behaviour, rather than relying only
  on isolated tests of individual functions.
- Test through the public API where practical, so tests remain useful
  when internal implementation details change. Add focused internal tests
  when they help verify complex logic or diagnose failures.
- Make the package's contract testable with controlled dependencies.
  Basic contract tests should not require assembling unrelated parts of
  the system.
- If testing a capability requires substantial infrastructure unrelated to
  the behaviour being verified, examine its responsibilities, dependency
  boundaries, and wiring. Do not compensate for hidden coupling with
  increasingly elaborate test scaffolding.
- Infrastructure-heavy setup is appropriate when the real interaction is
  itself under test. Treat unnecessary infrastructure as a design signal,
  not automatic permission to redesign unrelated code.
- Keep integration tests for verifying real dependency interactions and
  component wiring. They complement, rather than replace, independently
  testable package contracts.

## Error Contracts

For fixed, project-owned failure conditions, follow the
[constant sentinel error pattern](patterns/constant-errors.md). It covers
immutable declarations, ownership, wrapping, matching, and value identity.
Use structured error types when callers need runtime details rather than
only a fixed failure category.

## Parameter and Return Types

- Accept interfaces when a function needs a capability rather than a
  particular implementation, especially for dependencies. Prefer narrow
  interfaces describing only the behaviour the function requires.
- Define dependency interfaces in the consuming package where practical.
  Do not require callers to satisfy a broader interface merely because
  an implementation already provides one.
- Use concrete parameter types when the function needs a specific value
  or representation. Do not introduce interfaces for simple data
  arguments or solely to make a signature more abstract.
- Prefer concrete return types so callers receive an explicit result and
  can choose the interfaces through which they use it. In particular,
  constructors should normally return the concrete type they construct.
- Return an interface when abstraction is part of the operation's
  contract, such as selecting between implementations or preserving an
  intentional implementation boundary. Use conventional interface
  results, such as `error`, where appropriate.
- Apply this guidance alongside package-boundary rules. Do not expose
  internal representations merely to return a concrete type.

## Representation Contracts and Ownership

- Obtain representations through a type's public contract. Do not infer
  representation guarantees from its underlying type or current
  implementation.
- When a type does not provide what another package requires, first decide
  which package owns the missing capability.
- Extend the original type when the capability belongs to its general
  meaning and responsibilities. Document and test the new guarantees,
  preserving existing representation contracts.
- When the requirement belongs to the consumer, prefer a consumer-owned
  representation and explicit translation through supported contracts.
  Introduce a new type when it has distinct meaning, invariants, or
  behaviour. An existing type or a conversion function may be sufficient;
  do not introduce wrappers mechanically.
- Translation must not depend on undocumented source representations.
  Creating a consumer-owned type does not legitimise an unsupported
  conversion from the original type.
- Treat pressure to inspect or expose implementation details as a reason
  to examine the package boundary. Do not automatically expand the original
  type's API or redesign unrelated code to satisfy a consumer.

## Interface Arguments and Representation Boundaries

- When an operation takes several related inputs or is expected to gain
  options, prefer a named argument struct over a growing list of positional
  parameters. Avoid argument structs for simple, stable signatures.
- Define argument types alongside the owning interface when they belong
  specifically to that operation. Export them when they form part of an
  exported interface's contract.
- Reuse an established domain or contract type when its meaning and
  guarantees match the operation's needs. Do not duplicate argument types
  solely to colocate them with each consumer interface. Keep separate
  types when their contracts differ.
- Keep `context.Context` as the first parameter. Keep execution concerns,
  such as authorisation credentials, separate from operation options.
- Design interface argument types around caller intent, not an
  implementation's storage, transport, or serialisation format.
- Do not reuse internal HTTP payloads, database records, or other
  implementation-specific types as interface arguments merely because
  their fields currently match.
- Translate interface arguments into internal representations at the
  implementation boundary. Explicit mapping is preferable to coupling
  otherwise independent contracts.
- Keep serialisation tags and representation-specific fields on internal
  types unless serialisation is explicitly part of the interface contract.
- When adding argument fields, define their zero-value behaviour,
  defaults, and validation. A stable method signature does not by itself
  guarantee behavioural compatibility.

### Example: Reusing an Established Contract

At an HTTP-client boundary, a consumer can define a narrow interface using
`net/http`'s established request and response types:

```go
type httpDoer interface {
    Do(*http.Request) (*http.Response, error)
}
```

The consumer owns the interface; `net/http` owns the request and response
contracts. `*http.Client` implements this interface without consumer-local
copies of those types. Reuse is appropriate because the capability being
modelled is HTTP request execution under `http.Client.Do`'s contract.

This does not justify exposing HTTP types through unrelated business
interfaces. When a consumer has different requirements or guarantees,
define an appropriate consumer-owned contract rather than forcing reuse.
