# Verifying Fusion Work

This guidance applies after changing a Fusion design, before reporting the
work complete.

## Checks

Run each check that applies to the change. If a check cannot be run with
the available tools, report it as unverified rather than skipping it
silently.

- Confirm that every sketch the work created or changed is fully
  constrained.
- Confirm that the timeline rebuilds without errors or warnings.
- Measure key dimensions and compare them with the parameters that should
  drive them.
- Check for interference between components.
- Flex key parameters to confirm the model rebuilds with the intended
  behaviour, then restore their original values and confirm the
  restoration.
- Capture screenshots or section views that show the result.

## Report

Report the following at the end of the work:

- what changed, and the versions saved
- user parameters added or changed, with old and new values
- bodies left in place, and why
- the checks run and their results
- anything that could not be verified
