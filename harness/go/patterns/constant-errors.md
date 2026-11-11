# Constant Sentinel Errors

Use this pattern when defining or changing project-owned package-level
sentinel errors. A sentinel identifies a fixed failure condition that
callers can recognise without inspecting error text.

For new fixed sentinel errors, use named constants backed by a defined
string type implementing `error`, rather than mutable package variables
initialised with `errors.New`.

## Purpose and Scope

A constant sentinel cannot be reassigned by its package, callers, or tests.
It is a comparable value that works with standard Go error wrapping and
matching. The implementation needs only a defined string type and an
`Error()` method with a value receiver.

This pattern does not replace structured error types. When callers need
runtime details such as a field, position, or resource identifier, expose
those details through a structured error and inspect it with `errors.As`
or `errors.AsType`. Do not encode those details into dynamically created
constant-error values for callers to parse.

Keep using standard-library and dependency sentinels where their contracts
apply. Do not recreate them as local constants.

## Type Placement

Reuse an existing project-wide `ConstError` type when one is available.
Otherwise, defining the small string-backed error type in each package
that needs it is acceptable. Do not introduce a shared errors package
solely to eliminate duplication of this type.

Choose the support type's name and visibility according to the project's
needs and conventions. Export named sentinels when callers need them; the
support type does not have to be exported. Declare each sentinel in the
package that owns the failure contract.

## Definition and Use

This self-contained example uses a package-private support type and an
exported sentinel. Imports and declarations are shown together so the
example can be compiled as a package.

```go
package records

import "fmt"

// constError represents a fixed failure condition.
type constError string

const (
    // ErrNotFound indicates that the requested record does not exist.
    ErrNotFound constError = "record not found"
)

// Error returns the failure description.
func (e constError) Error() string {
    return string(e)
}

// Lookup returns the record associated with key.
//
// The following errors may be expected:
//   - [ErrNotFound]: key has no associated record.
func Lookup(records map[string]string, key string) (string, error) {
    value, ok := records[key]
    if !ok {
        return "", fmt.Errorf("looking up record %q: %w", key, ErrNotFound)
    }
    return value, nil
}
```

Return the named constant directly when no additional context is needed.
When adding operation-specific context, wrap it with `%w` to retain its
identity. Keep the sentinel fixed; do not reconstruct it with a formatted
string or turn its message into a new `errors.New` value.

Callers and tests match the particular named sentinel with `errors.Is`.
Matching only the support type with `errors.As` does not distinguish its
individual failure conditions.

Follow the [error inspection guidance](../style.md#error-inspection) and
[error assertion guidance](../testing.md#error-assertions). Do not compare
`Error()` strings or reconstruct a sentinel from its underlying literal
at call sites.

## Identity and Compatibility

Constant errors have value identity, not allocation identity. Two values
of the same defined string type with identical contents compare equal and
match through `errors.Is`. No custom `Is` method is needed.

- Give distinct failure conditions distinct values when they use the same
  defined type. Different constant names do not create different identity.
- A shared support type creates a shared identity namespace across the
  packages using it. Avoid accidental collisions between unrelated errors.
- Separately defined package-local types have distinct identities even
  when their string contents are identical. A type alias does not create
  that separation.
- Reference the owning package's named constant. Do not copy its literal
  or reconstruct its value to recognise the failure.
- The underlying string participates in identity. Review changes to that
  value or its defined type for compatibility, not merely presentation.
- An empty constant-error value is still a non-nil error when returned as
  `error`. Return `nil` to indicate success.

Do not mechanically migrate existing public error variables. Changing a
variable to a constant or changing its concrete type can affect callers.
Establish the compatibility requirements and keep migrations within the
requested task's scope.

## Testing the Contract

Test that the component returns the expected named sentinel, including
through contextual wrapping, and that a different failure does not match.
For the example above, this companion test exercises the public API:

```go
package records_test

import (
    "errors"
    "io"
    "testing"

    "example.com/project/records"
)

// TestLookupMissing checks that a missing record retains its sentinel.
func TestLookupMissing(t *testing.T) {
    recordsByKey := map[string]string{"present": "value"}
    _, err := records.Lookup(recordsByKey, "missing")
    if !errors.Is(err, records.ErrNotFound) {
        t.Fatalf("expected ErrNotFound, got %v", err)
    }
    if errors.Is(err, io.EOF) {
        t.Errorf("unexpected EOF match: %v", err)
    }
}
```

The import path is illustrative; use the consuming project's actual module
path and established testing framework.

When introducing the support type, also test its own contract: equal
values match, different values do not, direct and wrapped sentinels match,
and `Error()` returns the underlying string. Checking that string is
appropriate here because it is the support type's explicit representation
contract, not a way to classify errors from other operations.

## References

These implementations illustrate the pattern, not required dependencies
or package layouts:

- [Juju's shared support type](https://github.com/juju/juju/blob/0506e43aad9083a8be5b9a0e061bedffe346881e/internal/errors/const.go)
- [Juju's value and wrapping tests](https://github.com/juju/juju/blob/0506e43aad9083a8be5b9a0e061bedffe346881e/internal/errors/const_test.go)
- [Workshop's package-local support type and sentinels](https://github.com/canonical/workshop/blob/5060f5372e2f20153efb8c653b84f8318ee883e0/internal/secrets/errors.go)
- [Workshop's private support type with exported sentinels](https://github.com/canonical/workshop/blob/5060f5372e2f20153efb8c653b84f8318ee883e0/internal/secrets/provider/secretservice/errors.go)
