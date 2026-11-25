# Modelling in Fusion

This guidance applies when creating or changing geometry, sketches,
components, or parameters in a Fusion design.

## Structure

Follow the design type set on the document.

- **Part designs** have a single component. Model with bodies and do not
  create components.
- **Assembly and hybrid designs** SHOULD use a component for each separate
  physical part: anything made, bought, or moved on its own, such that it
  would appear as its own line in a bill of materials. Prefer components
  over bodies.

A component MAY contain several bodies when together they form a single
volume that is not used separately. Whether to keep construction or tool
bodies, or to combine the bodies of a finished part, is a judgement for
each case. Name every body by its role, such as `drillClearance`, and
report which bodies were left in place and why.

## Naming

- Use camelCase for parameters, components, bodies, sketches, and named
  features.
- Use technical names that describe what the thing is or does.
- Do not repeat a parent component's name in a child. Under `gearbox`,
  name a child `housing`, not `gearboxHousing`.

## Parameters

- Make a dimension a user parameter when it drives the design: overall
  sizes, mating interfaces, material thicknesses, clearances, and values
  used in more than one place. Incidental dimensions MAY remain plain
  values.
- Use metric units.
- Give each user parameter a comment that helps someone change it safely
  without tracing the timeline:
  - what it controls in the model, not a restatement of its name
  - where its value comes from, such as a measurement, a purchased
    part's datasheet, a process allowance, or a user decision; if the
    source is unknown, say so rather than guessing
  - any limits or relationships, such as a minimum for strength or a
    value that must match a mating part
- Keep comments short, and update them when a parameter's value source or
  purpose changes. For example, describe `boltClearance` as "Radial gap
  around M3 bolts in printed holes; value from user print tests", not
  "Bolt clearance".
- Agents MAY change existing user parameters when the change does not
  alter what the parameter is for. Report each change with its old and new
  values.
- Leave model parameters with their default names.

## Sketches

- Sketches MUST be fully constrained, using dimensions and geometric
  constraints that express the design intent.
- Agents MUST NOT use the Fix constraint, or dimensions with no meaning,
  only to remove remaining degrees of freedom. If the intent that should
  fix some geometry is unclear, ask.
- Give each sketch one clear purpose within its component, driving the
  features or child geometry related to that purpose. Do not combine
  unrelated concerns in one sketch without a reason.
- Sketches SHOULD reference existing geometry where that captures the
  intent, such as projecting a hole pattern from a mating face. Prefer
  stable references, such as origin geometry, other sketches, and primary
  faces, over edges produced by fillets, chamfers, or shells.

## Timeline

Rename or group timeline features only when it makes the history easier to
follow.
