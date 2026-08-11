# write-code-every-day

Daily Go practice for data structures and algorithms, organized around a phased curriculum and spaced repetition (katacoding).

## Prerequisites

- Go 1.22+
- `make` (optional, but recommended)

Clone the repo, then run tests from the repository root:

```bash
make test
```

## How to use this repo

### 1. Find your next problem

1. Open [plans/progress.md](plans/progress.md) and find the next `not started` problem in the current phase.
2. Read the phase notes in [plans/datastructures_and_algorithms.md](plans/datastructures_and_algorithms.md) for context and practice rhythm.

**Start here:** Phase 0 at `go/problems/phase0/lc0225_stack_using_queues/`.

### 2. Solve a problem (daily workflow)

Each problem lives in its own folder:

```
go/problems/phase0/lc0225_stack_using_queues/
├── solution.go       # your implementation
└── solution_test.go  # table-driven tests
```

**Steps:**

1. Open `solution.go` and implement the solution.
2. Run tests for that problem:

```bash
cd go && go test ./problems/phase0/lc0225_stack_using_queues/
```

3. When tests pass, remove `t.Skip(...)` from `solution_test.go` if it is still there.
4. Update the problem status in [plans/progress.md](plans/progress.md) (e.g. `not started` → `solved`).

**For a brand-new problem** (not yet scaffolded), create a folder using this naming convention:

```
go/problems/phaseN/lc<NNNN>_<slug>/
```

Example: `go/problems/phase1/lc0206_reverse_linked_list/`

Add `solution.go`, `solution_test.go`, and optionally `notes.md` for pattern notes or complexity analysis.

### 3. Run tests

From the repository root:

```bash
make test              # all tests
make test-archive      # pre-curriculum solutions only
make test-phase0       # Phase 0 problems only
make test-phase1       # Phase 1 problems only
make fmt               # format all Go files
```

Or run `go test` directly from the `go/` directory:

```bash
cd go
go test ./problems/phase0/...
go test ./problems/archive/...
```

### 4. Katacoding (weekend re-solve from memory)

Katacoding means backing up your solution and re-implementing it from scratch on a later day, without notes, until the pattern feels automatic.

```bash
make kata PROBLEM=problems/phase0/lc0225_stack_using_queues
```

This will:

1. Move `solution.go` to `solution_YYYYMMDD.go.bak`
2. Leave you with an empty slot to create a fresh `solution.go`

After re-solving, run the problem's tests again and update [plans/progress.md](plans/progress.md) (`solved` → `kata-1` → `kata-2` → `automatic`).

### 5. Use shared types (tree / linked-list problems)

For new problems that need `TreeNode` or `ListNode`, import the shared package instead of redefining types:

```go
import "github.com/qu1queee/write-code-every-day/go/internal/ds"
```

Archived solutions under `go/problems/archive/` keep their original local type definitions.

## Repository layout

```
go/
├── go.mod
├── internal/ds/           # shared TreeNode, ListNode
└── problems/
    ├── phase0/            # current curriculum — start here
    ├── phase1/ … phase4/  # upcoming phases
    └── archive/           # pre-curriculum solutions (reference + katacoding)
plans/
├── datastructures_and_algorithms.md   # full mastery plan
└── progress.md                        # checklist tracker
scripts/
└── kata.sh                # used by `make kata`
```

## Typical week

| Day | Activity |
|---|---|
| Mon / Wed / Fri | One new problem — timed, no IDE if possible, narrate your approach first |
| Tue / Thu | Go depth (pointers, receivers, memory layout) or a testing drill |
| Weekend | Katacoding — re-solve one structure from memory with `make kata` |

If the week is thin, protect the weekend katacoding slot over new material.

## CI

GitHub Actions runs `make test` on every push and pull request. Keep tests green before merging.

## Curriculum

- [plans/datastructures_and_algorithms.md](plans/datastructures_and_algorithms.md) — phased plan (32 problems, Phases 0–5)
- [plans/progress.md](plans/progress.md) — current status per problem
