# Requirement Levels

The key words MUST, MUST NOT, SHOULD, SHOULD NOT, and MAY in this harness
are to be interpreted as described in BCP 14, RFC 2119 and RFC 8174, when,
and only when, they appear in all capitals.

- [RFC 2119](https://www.rfc-editor.org/rfc/rfc2119)
- [RFC 8174](https://www.rfc-editor.org/rfc/rfc8174)

## Meanings

- **MUST:** an absolute requirement.
- **MUST NOT:** an absolute prohibition.
- **SHOULD:** follow the requirement unless a valid reason justifies an
  exception. Understand and weigh the consequences before departing.
- **SHOULD NOT:** avoid the behaviour unless a valid reason justifies an
  exception. Understand and weigh the consequences before proceeding.
- **MAY:** the behaviour is permitted but optional.

Ordinary instructions remain applicable without these keywords. Lowercase
wording does not make an instruction optional.

## Applying Requirements

These terms express obligations, not enforcement mechanisms. They do not
change instruction precedence, grant permissions, or replace sandboxing,
approval controls, or validation.

If a MUST requirement cannot be met, the agent MUST report the blocker
and pause the affected action. It MUST NOT silently skip the requirement
or claim compliance. Independent work MAY continue where it does not
violate the requirement. Higher-priority instructions still take precedence.

When departing from a SHOULD or SHOULD NOT requirement, the agent MUST
record the reason and material consequences in the relevant work record
or response. Convenience alone is not sufficient justification.

## Writing Requirements

Use these keywords selectively. Reserve MUST and MUST NOT for constraints
that admit no discretionary exception within their applicable scope.
Use SHOULD and SHOULD NOT for defaults with legitimate exceptions, and
MAY for explicit options.

State the scope and required behaviour clearly. Do not rely on uppercase
emphasis to resolve ambiguous wording or make a requirement enforceable.
