# Go Style

These rules apply to writing Go code.

## General


- Prefer non-nested error handling; use guard clauses or `else if` chains to
  avoid nested `if` blocks for error checks.
- Combine consecutive, mutually exclusive checks of the same error into an
  `if` / `else if` chain rather than separate `if` statements. This makes
  the alternatives explicit and prevents later checks from running after
  an earlier branch has handled the result. For example:
    ```go
    err := x()
    if err == nil {
        handleSuccess()
    } else if !errors.Is(err, SomeError) {
        return err
    }
    ```
  Keep independent checks separate when both must run.
- Never use inline error assignments in `if` statements
  (e.g. `if err := fn(); err != nil`). Assign the error to a variable on
  its own line first, then check it on the next line.
- Do not use `panic` as a failure-reporting mechanism. Return errors from
  functions whose contracts support them. At boundaries whose required
  signatures cannot return errors, use the established failure-reporting
  mechanism rather than changing the signature or silently discarding
  the failure.
- Avoid `New*` functions that return errors; use names like `Make*` or
  `Construct*` when construction can fail.
- Prefer package-private helpers unless API exposure is necessary.
- Prefer passing and returning small, simple structs by value. Do not use
  pointers by default or solely to avoid copying.
- Use a pointer when the function must modify the caller's struct, shared
  identity is required, `nil` represents a meaningful absent value, or the
  type must not be copied (for example, a struct containing a used mutex).
  Pointers are also appropriate when required by an existing API or when
  copying a large struct has a demonstrated performance cost.
- Passing a struct by value copies its fields, not the data referenced by
  slice, map, or pointer fields. Do not treat a value parameter as a
  guarantee that the caller's underlying data cannot be modified.
- Do not assume pointers improve performance or values remain on the
  stack. Let escape analysis and measurements guide allocation-related
  decisions.
- Prefer value receivers for methods on small, simple structs when the
  method does not modify the receiver or its referenced data.
- Use pointer receivers when a method must modify the original struct,
  preserve shared identity, or avoid copying a type that must not be
  copied (for example, a struct containing a used mutex). They are also
  appropriate when copying a large receiver has a demonstrated
  performance cost.
- Keep receiver types consistent across a type's methods. If some methods
  require pointer receivers, prefer pointer receivers for the remaining
  methods rather than mixing receiver types.
- Do not use pointer receivers merely because the receiver is a struct.
  If none of the methods needs pointer semantics and the struct is small
  and safe to copy, use value receivers throughout.
- Omit the receiver name when a method does not use the receiver in its
  body. Do not introduce an unused receiver name or replace it with `_`.
  Retain the appropriate value or pointer receiver type.
- Prefer `typeVal.String()` when its contract provides the representation
  the operation requires. Convert a defined string-based type directly to
  `string` only when its public contract explicitly guarantees the required
  meaning of that conversion. Conversions within the type's own
  implementation of its representation methods are allowed.
  Implementing `fmt.Stringer` alone does not guarantee a stable or
  machine-readable representation. If no existing contract supplies the
  required representation, follow the
  [representation guidance](design.md#representation-contracts-and-ownership).
- When building a slice in a loop, prefer pre-allocating capacity with
  `make([]T, 0, n)` and growing with `append`, rather than pre-sizing with
  `make([]T, n)` and assigning by index. The append form makes the build
  pattern obvious, avoids stray zero values if the loop exits early, and
  removes the need for a manual index variable.

## Error Messages

- When adding operation context to an error, begin with a lowercase
  present participle ("-ing" verb) describing the operation. For example:
  `fmt.Errorf("reading credentials file: %w", err)`.
- Prefer "reading", "creating", and "checking" over "read", "create",
  and "check". Avoid redundant prefixes such as "failed to" or
  "error while".
- Preserve the underlying error with `%w` when wrapping.
- Errors describing invalid state or validation constraints may state
  the problem directly, e.g. "credentials file is empty".

## File and Declaration Layout

- For non-test Go files, order sections as: imports, types, consts, vars,
  funcs. Keep entries in each section alphabetical.
- For structs and interfaces, keep member declarations in alphabetical order.
  When members have their own comments, separate each documented member with a
  blank line. Undocumented members may remain grouped together when that
  improves readability.
- Preserve semantic ordering when alphabetical ordering would affect
  behaviour, initialisation, or a required data layout, including meaningful
  ordering within `iota` groups.
- Avoid underscores used only as word separators in Go filenames.
  Preserve underscores required by Go tooling, including platform-selection
  and `_test.go` suffixes.
  For example, use `parsereadrequest_test.go`, not
  `parse_read_request_test.go`.

## Formatting

- Prefer shallow scope and fewer indentation levels. Long functions are
  fine when the logic is simple and flat (negative-space style).
- Do not collapse function bodies onto a single line; always use the full
  brace-and-newline form for readability.
- Group related logic together without blank lines between related
  statements. Use blank lines only to separate distinct blocks of logic.
- Do not use section divider comments (e.g. `// -----`).
- Keep Go source lines within 80 characters. Allow exceptions for content
  that cannot reasonably be split, such as struct tags, tool directives,
  import paths, URLs, or indivisible literals. Do not change semantics
  solely to satisfy the line-length limit. The line-length and multi-line
  argument rules below apply to both function definitions and call sites.
- When a function call or signature does not fit on one line, place each
  argument or parameter on its own indented line. Start the closing
  parenthesis on a new line.
- If the return types still exceed 80 characters when on one line, break
  them onto individual lines between their own parentheses.
- When a function has more than three return values, return a struct
  (alongside the error) rather than a long return tuple.

### Formatting Examples

The examples progress from compact forms to nested calls and function-type
conversions. They are fragments; imports and surrounding declarations are
omitted.

#### Compact Functions and Calls

Keep short signatures and calls compact; wrapping is not required merely
because a function has parameters.

```go
func normaliseName(name string) string {
    return strings.TrimSpace(name)
}
```

```go
name := normaliseName(input)
```

#### Wrapped Calls

When a call exceeds 80 characters, place each argument on its own line.

```go
qualifiedName := qualifiedResourceName(
    resourceNamespace,
    resourceKind,
    resourceName,
)
```

#### Wrapped Function Definitions

When an ordinary function definition's signature exceeds 80 characters,
place each parameter on its own line. Keep the return types on the closing
parenthesis line when they fit.

```go
func qualifiedResourceName(
    resourceNamespace string,
    resourceKind string,
    resourceName string,
) string {
    return resourceNamespace + "/" + resourceKind + "/" + resourceName
}
```

#### Wrapped Return Values

If the return types still exceed 80 characters after wrapping the
parameters, wrap the return list too. Here, `ResourceConfiguration` and
`ResourceValidationReport` represent application types.

```go
func buildResourceIndexes(
    resources []ResourceConfiguration,
) (
    map[string]ResourceConfiguration,
    map[string]ResourceValidationReport,
    error,
) {
    // Implementation omitted.
}
```

#### Nested Calls

Consider each call independently. The outer call needs wrapping, but the
inner call remains compact because it fits.

```go
return fmt.Errorf(
    "validating resource %q in namespace %q: %w",
    strings.TrimSpace(resourceName),
    namespace,
    err,
)
```

#### Function Literals as Arguments

`http.HandleFunc` registers the supplied function as a request handler.
The function literal is one argument even though its body spans lines.
Its short signature remains compact; the trailing comma follows its
closing brace, and the call's closing parenthesis is on a separate line.

```go
http.HandleFunc(
    "/health",
    func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
    },
)
```

#### Wrapped Function-Literal Signatures

Apply the parameter-layout rule to anonymous functions too. Here,
`processRequest` represents existing application logic; `handleRequest`
is a function value that delegates to it. The same signature layout
applies when passing the function literal as an argument.

```go
handleRequest := func(
    requestContext context.Context,
    resourceName string,
    resourceVersion string,
) error {
    return processRequest(requestContext, resourceName, resourceVersion)
}
```

#### Named Function-Type Conversions

Unlike `http.HandleFunc`, `http.HandlerFunc` is a named function type.
This expression converts a function value into that type, which implements
`http.Handler`. It does not register or execute the handler. The conversion
uses the same multiline layout as a call with a function-literal argument.

```go
handler := http.HandlerFunc(
    func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
    },
)
```

