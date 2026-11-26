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
- **Pacing:** how many questions each message asks. Default: one. A
  caller may ask for a round's questions together.
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

- Work in rounds. A round covers the frontier, within the round focus.
- Following the pacing setting, ask one question, wait for the user's
  answer, then ask the next. Use the round format below.
- Ask the most consequential question first. State each question's
  place in the round, such as "2 of 7", so the user can see how much
  remains without reading it all.
- Give the remaining questions together when the user asks for them.
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

### Showing the Question

- Give the user what they need to answer in the message that asks. Do
  not rely on them recalling earlier discussion or sharing your framing.
- Where the question concerns something the user can see, show it before
  asking: a worked example in the form its consumer meets, such as a
  command with output, a configuration fragment, or an exchange. Show
  each option in the same form so the difference is visible.
- Keep explanation around the example to a sentence or two. Ask plainly
  when a question has no such form, such as one about goals or audience.
  Do not invent an example to fill the format.

### Deferred Questions

- The user may defer a question or say they cannot answer it yet.
  Accept that as an answer. Do not press, and do not ask it again in
  the same round.
- Note what a deferred question waits on, if the user says. Hold the
  questions that depend on it and carry on with the rest.
- Do not settle a deferred question by assumption or by your
  recommendation.
- Raise deferred questions again when what they wait on arrives, when
  the user asks, and before finishing.

### Finishing

- The interview is complete when the frontier is empty and each deferred
  question has been raised again: every branch visited and nothing
  silently assumed. A question the user leaves deferred is recorded as
  open, not settled.
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
❓ **Q3 - <question title>** (2 of 7)

<worked example of the situation, where the question has one>

<each option, shown in the same form>

<the question, in a sentence>

➡️ <recommended answer, or options with your recommendation labelled>
```

When asking several questions together, separate them with `---`.

Treat a reply as settling only the questions it clearly addresses. Carry
any question it leaves unanswered into the next round. This format is for
conversation only; records and documents follow the harness documentation
guidance.

## Output

Follow the output setting. When recording the outcome, include settled
decisions with any reasoning the user gave, open and deferred questions
and what they wait on, assumptions the user accepted, and facts found
with their sources.

## Attribution

Adapted from the `grilling` skill in
[Matt Pocock's skills repository](https://github.com/mattpocock/skills),
used under the MIT licence. See [LICENSE](LICENSE) for the notice.
