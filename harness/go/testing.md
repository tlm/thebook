# Go Testing

These rules apply when writing or modifying Go tests.

## General

- For a new Go project with no established testing library or pattern,
  use `github.com/juju/tc`.
- In an existing project, follow its established testing libraries and
  patterns, including standard-library testing where that is the
  convention. Do not introduce `tc` solely to replace them.
- For test files that do not define direct tests or suites, follow the same
  section and alphabetical ordering rules as non-test Go files.
- Use table tests when input/output combinations are the main concern.
- When using table tests, execute each case as a subtest (`t.Run(...)`).
- Avoid test helpers with heavy branching logic. Duplicate test code
  rather than over-applying DRY in tests.
- Verify interface contracts through named, documented tests rather than
  standalone compile-time assertions. Exercise the intended value or
  pointer form through the interface; a direct method call may not verify
  that contract.
- When conformance itself needs a named test failure, use a checked runtime
  type assertion, then verify behaviour through the resulting interface.
- Prefer complete, realistic test fixtures over partially populated structs.
  When a field is required for a valid domain value, set it explicitly in
  tests even when the specific test behaviour does not depend on that field.
- Prefer constructing domain values through their public constructors in
  tests. Do not cast arbitrary literals to domain types when a `Make*` or
  `Construct*` function exists, unless the test is specifically exercising
  invalid or legacy values.
- Prefer deterministic tests:
    - avoid sleeps used to assume another goroutine has progressed; prefer
      explicit signals and bounded waits. Use deliberate real-time delays
      only when necessary to verify timing behaviour itself.
    - use fakes/stubs over heavyweight mocks
- For self-contained concurrent or time-dependent tests, consider
  [testing/synctest](https://pkg.go.dev/testing/synctest) when the project's
  minimum supported Go version is 1.25 or later. Use `synctest.Test` for an
  isolated bubble with fake time, and `synctest.Wait` to wait for the other
  goroutines in that bubble to become durably blocked.
  Sleeps inside a bubble use fake time rather than wall-clock delays.
  The package does not control real network I/O or external processes.
  Do not raise the project's Go version solely to use it.
- Use real servers or substantial infrastructure when their actual
  behaviour or integration is part of what the test must verify. Otherwise,
  prefer controlled test doubles at meaningful dependency boundaries.
  Keep real test dependencies controlled; avoid relying on uncontrolled
  external services.
- Infrastructure-backed tests need a controlled execution environment.
  Establish dependency readiness, state isolation, bounded waits, and
  reliable cleanup rather than relying on timing assumptions or unrelated
  shared state.
- Do not accept known, avoidable nondeterminism as an inherent cost of
  testing real integrations. If the ordinary test harness cannot provide
  the required environmental guarantees, use an appropriate integration
  or system-test harness that can. Preserve the required real-interaction
  coverage.
- Investigate intermittent failures before classifying them as test
  flakiness; they may reveal defects in the code under test. Moving tests
  or retrying them until they pass is not a substitute for addressing the
  cause.
- Run assertions that terminate a test, such as `t.Fatal` or assertions
  that call `FailNow`, only from the goroutine executing that test or
  subtest. Worker goroutines should report results or errors to that
  goroutine for assertion.
- Arrange cleanup to stop and join test-owned goroutines, including when
  assertions fail, so workers do not outlive the test or its resources.
- Place shared test doubles in `mocks_test.go` for the package.
- Document each test with a short comment describing:
    - what interface/functionality is under test
    - what behaviour is expected
- Document custom test doubles with a short comment describing their role.

## Contract-Driven Testing

- Use documented purpose, preconditions, postconditions, and failure
  guarantees to guide test cases and assertions.
- Contract tests verify observable guarantees callers rely on, including
  results, state changes, and failure behaviour. They do not prescribe
  internal composition. Do not reproduce implementation logic in tests
  merely to confirm that the same steps were followed.
- Derive expected results independently of the behaviour under test.
  Do not use the same parsing, normalisation, or calculation path to
  produce both actual and expected results.
- Public constructors remain appropriate for ordinary fixtures when their
  behaviour is not what the test verifies.
- For bug fixes, establish the intended contract and focused regression
  coverage before changing the implementation. Use existing tests when
  they already expose the defect; otherwise add dedicated test cases.
- Run the regression tests against the unfixed implementation and confirm
  they fail because of the original defect, not invalid setup or unrelated
  failures. Then make the fix and rerun the same tests to verify the
  required behaviour.
- If reproducing the original failure before the change is impractical,
  explain the limitation and report what was validated and what remains
  unverified.
- Contract tests should not require changes solely because internal
  composition changes while the observable contract remains unchanged.
  Their purpose is to detect behavioural regressions across refactors,
  not prevent reorganisation of private code.
- Avoid assertions about internal call order, helper usage, or private
  representation unless those details are necessary to verify an
  explicit contract. Prefer assertions about results, state changes,
  and externally observable effects.
- Use test doubles to control dependencies and verify meaningful
  interactions. Do not prescribe an exact sequence of dependency calls
  unless that sequence is part of the required behaviour.
- Focused internal tests are appropriate for complex implementation
  logic, but they complement contract tests rather than replace them.
  Keep implementation-specific assertions local to those tests. These
  tests may change with the implementation details they verify.
- If a change breaks a test, determine whether the intended contract
  changed, the implementation violated it, or the test was unnecessarily
  coupled to implementation details. Treat unexpected contract-test
  failures as potential regressions requiring investigation, not routine
  test maintenance. When the intended contract is unchanged, establish
  the cause before changing expectations. Do not weaken valid assertions
  or relabel contract tests as implementation-specific merely to make
  changed code pass.
- Existing projects may have tests coupled to private implementation
  choices. Identify the behaviour they protect before changing them.
  Refactor tests to reduce unnecessary coupling when the improvement is
  clear, the change is small and within the current task's scope, and it
  preserves checks of required behaviour.
- If the intended contract or the value or scope of a test refactor is
  unclear, ask the human for guidance. Explain the concrete issue and
  proposed approach in plain language, then ask a focused question. Keep
  the explanation concise unless more detail is requested.

## Error Assertions

- Follow the [error inspection guidance](style.md#error-inspection) in
  tests. Use `errors.Is`, `errors.As`, and `errors.AsType` for the matching
  operations they define.
- Tests MUST NOT hard-code error text as an expected result unless the
  wording is an explicit contract, a user-facing element being tested,
  or part of the narrowly permitted dependency integration described in
  that guidance. Testing user-facing wording does not permit using it as
  a substitute for error classification.
- For typed errors, assert successful extraction before inspecting the
  result. Check the structured details that establish the expected
  failure, rather than stopping at the error type when the test requires
  a more specific distinction.
- When removing error-text assertions, preserve the required failure
  distinctions through error identity, types, and relevant structured
  details. Do not weaken the test to accept any error merely because
  message wording changed.
- If project-owned errors cannot express a required failure distinction
  without parsing text, address the error contract rather than adding
  message-dependent assertions.

## Testing with `github.com/juju/tc`

The rules in this subsection apply only when the `github.com/juju/tc`
package is being used for testing in the project. If the project does not
import `github.com/juju/tc`, skip this subsection entirely.

- Use suite tests with `github.com/juju/tc` (`tc`) for behavior-focused
  command or workflow testing.
- For test files that define `tc` suites, order sections as:
    - types in alphabetical order
    - consts in alphabetical order
    - vars in alphabetical order
    - non-test funcs (helpers, fakes, builders) in alphabetical order
    - standard Go `TestXxx(t *testing.T)` suite registration funcs
    - suite test case methods in alphabetical order
- When defining new `tc` suites, document the suite struct with a short
  comment describing the type, functionality, or grouping under test.
- For suites, use:
    - `tc.Run(t, &suiteType{})`
    - suite methods with `func (suiteType) TestXxx(c *tc.C)`
    - when a context is needed inside suite test cases, use `c.Context()`
- In suite tests, when entering a closure that takes `*testing.T`, do not use
  the outer `*tc.C` inside that closure.
- Name the closure `*testing.T` parameter `c` to reduce accidental outer
  `*tc.C` usage.
- In closures where `*tc.C` is unavailable, use `tc.Assert(c, ...)` and
  `tc.Check(c, ...)` with the closure `c *testing.T`.
- In suite tests, use `c.Assert` and `c.Check` instead of `t.Fatal*`.
- Prefer `c.Check` when a failure can reasonably allow the rest of the test
  to continue and provide additional feedback.
- Use `c.Assert` when a failure means the test cannot continue safely or
  meaningfully.
- If an assertion is the last check in a test (no meaningful steps follow),
  prefer `c.Check` over `c.Assert`.
- Keep test case methods ordered alphabetically in suite files.
