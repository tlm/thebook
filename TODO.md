# TODO

- [ ] Add a section covering IPv6, dual-stack networking, and IETF
  documentation address ranges.
- [ ] Add guidance on prompt injection and ignoring hidden elements in
  Markdown files.
- [ ] Write rules for when to add and how to use Go's `context.Context`.
- [ ] Avoid verbatim error-string assertions unless the text is part of a UX
  contract.
- [ ] Improve test documentation to explain the contract under test.
- [ ] Require new test cases for new functionality rather than adding new
  requirements to existing test cases.
- [ ] Require complete, valid input data in test fixtures and doubles, even
  when knowledge of the implementation suggests some fields are unused.
- [ ] Do not use table-based tests inside Go check suites.
- [ ] Restrict table-based tests to frameworks with native support for Go
  subtests.
- [ ] Keep test setup and input composition visible and readable. Avoid
  hiding fixture construction behind helpers when doing so obscures what
  the test uses.
- [ ] Define guidance for opinionated test helpers whose names explain the
  fixture or behaviour they provide. Avoid signatures full of tuning knobs;
  add a separately named and documented helper when different assumptions
  are required. Assertion helpers can be useful, but every helper needs a
  justification beyond hiding lines of code.
- [ ] Document Go error design patterns.
- [ ] Add an 80-character line-length rule for Markdown files.
- [ ] Describe grouping and ordering of package-level variables and
  constants, including naming conventions modelled on Go's `net/http`
  package.
- [ ] Use an appropriate portable shebang for shell scripts.
- [ ] Use the features of the chosen shell interpreter; for example, use
  Bash features in Bash scripts.
- [ ] Prefer Bash for shell scripts.
- [ ] Add guidance on Go example tests for packages.
- [ ] Document Conventional Commits.
- [ ] Expand the new-project initialisation workflow incrementally with
  version-control setup and other project working conventions.
- [ ] Add a skill for setting up the Canonical knowledge-base convention in
  existing repositories, preserving existing guidance and documentation
  and agreeing any migration scope with the user.
- [ ] Define source-control action authorisation: distinguish staging,
  committing, pushing, pull requests, merging, and history rewriting;
  keep permission scoped to the requested work rather than treating it
  as standing authorisation.
- [ ] Define working-tree safety for discarding, overwriting, cleaning,
  resetting, and stashing changes, preserving unrelated work and requiring
  permission for destructive operations.
- [ ] Define signing and identity rules: respect configured identity and
  signing, and do not bypass failures by disabling signing or substituting
  an identity.
- [ ] Define hook and validation handling: address relevant failures or
  report blockers rather than bypassing hooks to make a commit succeed.
- [ ] Define remote-operation safety: verify the intended remote and branch,
  confirm publication scope, and require explicit permission to force-push.
- [ ] Define branch and worktree safety when creating, switching, or
  removing them, preserving existing work and following project conventions.
- [ ] Define conflict-resolution guidance: understand and preserve intended
  changes, avoid mechanical resolutions, and validate the result.
- [ ] Define checks for accidental credentials and private data before
  staging or publishing changes.
