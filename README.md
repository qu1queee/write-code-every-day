# write-code-every-day

Daily coding practice repo — Go data structures & algorithms with spaced repetition (katacoding).

## Quick start

```bash
make test              # run all tests
make test-archive      # run archived solutions only
make test-phase0       # run Phase 0 stubs
```

## Layout

```
go/
├── go.mod
├── internal/ds/           # shared TreeNode, ListNode
└── problems/
    ├── phase0/            # current curriculum (start here)
    ├── phase1/ … phase4/
    └── archive/           # pre-curriculum solutions
plans/
├── datastructures_and_algorithms.md   # mastery plan
└── progress.md                        # checklist tracker
```

## Workflow

1. Pick the next problem from `plans/progress.md`.
2. Create a folder: `go/problems/phaseN/lc<NNNN>_<slug>/`.
3. Implement `solution.go` with table-driven tests in `solution_test.go`.
4. Remove `t.Skip(...)` from the test once the solution works.
5. Weekend katacoding — back up and re-solve from memory:

```bash
make kata PROBLEM=problems/phase0/lc0225_stack_using_queues
```

## Curriculum

See [plans/datastructures_and_algorithms.md](plans/datastructures_and_algorithms.md) for the full phased plan and [plans/progress.md](plans/progress.md) for current status.
