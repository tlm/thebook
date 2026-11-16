---
name: grilling
description: >-
  Grilling the user about a plan, design, or decision until you reach a
  shared understanding, when the user asks to be grilled or to stress-test
  their thinking, or a skill or workflow step directs it.
---

# Grilling

## Operating Context

Before executing this skill, read and follow
[Skill operating context](../CONTEXT.md).

## Role

Interview the user relentlessly until you both reach a shared
understanding of the subject. Probe every decision and its consequences.
Nothing is settled by assumption.

## Caller Settings

The user, or a skill or workflow step invoking this skill, may supply the
settings below. Supplied settings take precedence over the defaults. When
the user invokes the skill directly, use the defaults unless they say
otherwise.

- **Subject:** what is being decided, and its boundaries. Default: the
  plan, design, or decision the user names. If it is unclear, ask first.
- **Fact-finding:** how facts are obtained. Default: look them up
  yourself, directly or through a subagent, within your existing
  permissions. A caller may route some investigations elsewhere, such as
  a tracked research process.
- **Round focus:** which open questions each round asks. Default: the
  whole frontier. A caller may limit rounds to a current topic.
- **Recommendation policy:** Default: give a recommended answer for each
  question. A caller may require options instead, with any recommendation
  labelled as yours.
- **Output and destination:** what to produce and where. Default: keep
  everything in the conversation and write nothing. A caller may name an
  output and where to record it.
- **Completion:** who confirms shared understanding, and what follows.
  Default: the user confirms, then control returns to whoever invoked the
  skill.

## Instructions

### Design Tree

Map the subject as a design tree: each decision branches into the
decisions that depend on it. The frontier is every open decision whose
prerequisites are settled, meaning the questions you can ask now without
guessing at answers you have not heard.

### Rounds

- Work in rounds. Ask the frontier, within the round focus, in a single
  round using the round format below. Then wait for the user's answers.
- Hold a question for a later round when it depends on another question
  still open in the current one.
- After each round, record what was settled, recompute the frontier, and
  ask the next round. Settled decisions unblock the questions that
  depended on them.
- When an answer changes an earlier decision, reopen the decisions that
  depended on it rather than carrying them forward.

### Facts and Decisions

- Look up facts you can find yourself, following the fact-finding
  setting. Ask the user only for decisions and for knowledge only they
  hold, such as their goals, preferences, constraints, and history.
- While a lookup is running, keep asking frontier questions that do not
  depend on it. Only the questions downstream of it wait.
- Decisions belong to the user. Put each one to them and wait for an
  answer. A recommendation is not a decision, and an unanswered question
  stays open.

### Finishing

- The interview is complete when the frontier is empty: every branch
  visited and nothing silently assumed.
- Confirm with the user that you have reached a shared understanding
  before acting on the outcome. Then follow the completion setting.
- If the user pauses or stops early, keep the open frontier according to
  the output setting and stop.

## Round Format

Number questions across the interview, `Q1`, `Q2`, and so on, so answers
can refer to them. Word each question so that "yes" accepts the
recommended answer. When the recommendation policy requires options, list
them, label any recommendation as yours, and ask the user to choose.

```text
❓ **Q1 - <question title>**: <the question, with context and any options>

➡️ <recommended answer, or options with your recommendation labelled>

---

❓ **Q2 - <question title>**: <the question>

➡️ <recommended answer>
```

Treat a reply as settling only the questions it clearly addresses. Carry
any question it leaves unanswered into the next round. This format is for
conversation only; records and documents follow the harness documentation
guidance.

## Output

Follow the output setting. When recording the outcome, include settled
decisions with any reasoning the user gave, open questions and what they
wait on, assumptions the user accepted, and facts found with their
sources.

## Attribution

Adapted from the `grilling` skill in
[Matt Pocock's skills repository](https://github.com/mattpocock/skills),
used under the MIT licence. See [LICENSE](LICENSE) for the notice.
