---
name: leetcode-practice
description: >-
  Interactive step-driven LeetCode practice guide for this repo's phased curriculum.
  Pick a phase problem, coach through restate → approach → scaffold → implement → tests → verify
  with next/back/run/drop controls. Never spoil tests or solutions early.
  Use when the user says start practice guide, practice phase N, leetcode guide,
  interactive practice, or wants guided solving of a curriculum LC problem.
---

# LeetCode Practice Guide

Interactive coach for one problem at a time from this repo's curriculum. Generic steps; the user chooses the problem.

## Session start

1. Read `plans/progress.md` and `plans/datastructures_and_algorithms.md`.
2. If no phase given, use the earliest phase that still has `not started` (or ask).
3. List that phase's problems with LC #, title, status, folder (or `—` if missing).
4. Ask the user to pick a problem. Prefer the **Ask questions** tool (clickable options: one option per LC). If that tool is unavailable, list options in chat and wait for a typed pick.
5. Enter **step 1**. Track: `phase`, `lc`, `folder`, `step` (1–6).

Folder convention: `go/problems/phaseN/lc<NNNN>_<slug>/`  
Shared types: `github.com/qu1queee/write-code-every-day/go/internal/ds`  
Kata: `make kata PROBLEM=problems/phaseN/lcNNNN_slug`

## Controls

After every step reply (except when the guide has ended), collect the next action via **Ask questions** — never as plain-text-only when the tool exists.

**Required:** Call the Ask questions / `AskQuestion` tool with a single question, e.g. prompt `Guide control — what next?`, `allowMultiple: false`, and exactly these options:

| id | label |
|----|-------|
| `next` | next — advance to the next step |
| `back` | back — return to the previous step |
| `run` | run — execute this step's action |
| `drop` | drop — end the guide |

Then handle the selected id:

| Command | Behavior |
|---------|----------|
| `next` | Advance one step if current is done, or user explicitly skips |
| `back` | Previous step; keep files; restate that step's guidance |
| `run` | Execute **only** the current step's action (table below) |
| `drop` | End guide immediately; no further coaching; do not update progress unless step 6 already finished |

**Fallback** (only if Ask questions is not available in this session): end the reply with `[next] [back] [run] [drop]` and accept typed `next`/`back`/`run`/`drop` or `n`/`b`/`r`/`d`.

Do not ask the control question as free-form chat prose when the tool is available.

## Reply template

1. Post the step body in chat:

```
### Step N/6 — <Title>
Phase X · LC <n> · `<folder or pending>`

<2–5 sentences of guidance for this step only>

**Your move:** <one concrete prompt for the user>
```

2. Immediately call Ask questions for the four controls (or print the text fallback bar).

Keep replies short. Do not dump all steps at once.

## Steps

### 1 — Restate
Show problem title, LeetCode URL, and a brief prompt (from folder header comment or public problem statement). User restates constraints + I/O in their own words.

**run:** Fetch/show the problem prompt only (no solution, no tests).

### 2 — Approach
User narrates algorithm and time/space complexity. **No code.**

**run:** Critique the narration; give hints only. Full algorithm dump only if user says they are stuck and asks for it.

### 3 — Scaffold
Ensure `go/problems/phaseN/lcNNNN_slug/` exists with stub `solution.go` (package, types/signatures, `panic("not implemented")` bodies). Match Phase 0 stub style.

**run:** Create/update stubs only. Do **not** create or open `solution_test.go`. For list/tree problems, import `internal/ds` instead of redefining nodes.

If `solution.go` already has a real implementation (not just panics), offer kata mode: `make kata PROBLEM=problems/phaseN/lcNNNN_slug`, then continue from step 4.

### 4 — Implement
User writes the solution in `solution.go`.

**run:** Review their code for bugs/edge cases. Do **not** rewrite a full working solution unless they say `reveal solution`.

### 5 — Tests
User adds or enables table-driven `solution_test.go` (may start with `t.Skip(...)`).

**run:** Help write tests or remove `t.Skip`. Opening tests is allowed from this step onward.

### 6 — Verify
**run:** `cd go && go test ./problems/phaseN/lcNNNN_slug/`  
On green: update that row in `plans/progress.md` (`not started` → `solved`, or advance kata status). Confirm done and end the guide — do not show the four-way control picker.

## Anti-spoil rules (while guide active)

Until step 5:

- Do not open, cite, or paraphrase `solution_test.go`, `*.bak`, or archive solutions for the same problem.
- Do not paste a complete working solution in step 4 unless the user says `reveal solution`.

If the user asks to peek early, refuse briefly and point them to `next`/`run` for the current step, or `drop` to leave the guide.

## Done checklist

- [ ] Tests pass for the problem package
- [ ] `plans/progress.md` status updated
- [ ] Guide ended (no further control picker)
