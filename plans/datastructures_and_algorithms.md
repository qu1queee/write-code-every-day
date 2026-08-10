# Go Data Structures & Algorithms — Mastery Plan

> Duration: ~10–11 weeks at 3 new problems/week. Weekend katacoding review.
> Self-paced — phases are the fixed sequence, weeks are elastic. No deadline.
> 32 new problems total
> Katacoding is repeatedly re-solving the same problem from memory over spaced-out days until the underlying
  pattern becomes automatic, the way a martial-arts kata is drilled until the moves need no conscious thought.
---

## Phase 0 — Go Implementation & Testing Fluency
**~1 week. Do this first — it's a multiplier for every phase after it.**

Memorize the table-driven test skeleton cold:

```go
func TestFoo(t *testing.T) {
    tests := []struct {
        name string
        in   int
        want int
    }{
        {"basic", 3, 9},
        {"zero", 0, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Foo(tt.in)
            if got != tt.want {
                t.Errorf("Foo(%d) = %d, want %d", tt.in, got, tt.want)
            }
        })
    }
}
```

Drill: given an interface spec, scaffold struct + stub methods + one table-driven test in under 15 minutes, timed.

| # | Problem | Link |
|---|---|---|
| LC 225 | Implement Stack using Queues | leetcode.com/problems/implement-stack-using-queues |
| LC 232 | Implement Queue using Stacks | leetcode.com/problems/implement-queue-using-stacks |
| LC 155 | Min Stack | leetcode.com/problems/min-stack |


---

## Phase 1 — Linear Structures
**~2–3 weeks (7 problems).**

| # | Problem | Link |
|---|---|---|
| LC 206 | Reverse Linked List | leetcode.com/problems/reverse-linked-list |
| LC 141 | Linked List Cycle | leetcode.com/problems/linked-list-cycle |
| LC 21 | Merge Two Sorted Lists | leetcode.com/problems/merge-two-sorted-lists |
| LC 19 | Remove Nth Node From End of List | leetcode.com/problems/remove-nth-node-from-end-of-list |
| LC 20 | Valid Parentheses | leetcode.com/problems/valid-parentheses |
| LC 3 | Longest Substring Without Repeating Characters | leetcode.com/problems/longest-substring-without-repeating-characters |
| LC 643 | Maximum Average Subarray I | leetcode.com/problems/maximum-average-subarray-i |

**Weekly rhythm:** Mon/Wed/Fri — one new problem, timed, no-IDE, narrate first. Tue/Thu — Go depth (pointer receivers, linked-structure memory layout) or a testing drill. Weekend — katacoding: re-implement one structure from memory, no notes.

---

## Phase 2 — Hashing, Trees, Heaps
**~2–3 weeks (7 problems). Hashing itself is already strong — no new hashing problems needed.**

| # | Problem | Link |
|---|---|---|
| LC 94 | Binary Tree Inorder Traversal | leetcode.com/problems/binary-tree-inorder-traversal |
| LC 102 | Binary Tree Level Order Traversal | leetcode.com/problems/binary-tree-level-order-traversal |
| LC 98 | Validate Binary Search Tree | leetcode.com/problems/validate-binary-search-tree |
| LC 701 | Insert into a Binary Search Tree | leetcode.com/problems/insert-into-a-binary-search-tree |
| LC 215 | Kth Largest Element in an Array | leetcode.com/problems/kth-largest-element-in-an-array |
| LC 347 | Top K Frequent Elements | leetcode.com/problems/top-k-frequent-elements |
| LC 1046 | Last Stone Weight | leetcode.com/problems/last-stone-weight |

215, 347, and 1046 are your heap-from-scratch practice — build the min-heap by hand first, then re-solve using Go's `container/heap` to learn both.

---

## Phase 3 — Graphs, Tries, Design Problems
**~2–3 weeks (7 problems). This phase is built directly around today's failure category.**

| # | Problem | Link |
|---|---|---|
| LC 200 | Number of Islands | leetcode.com/problems/number-of-islands |
| LC 207 | Course Schedule | leetcode.com/problems/course-schedule |
| LC 133 | Clone Graph | leetcode.com/problems/clone-graph |
| LC 208 | Implement Trie (Prefix Tree) | leetcode.com/problems/implement-trie-prefix-tree |
| LC 146 | LRU Cache | leetcode.com/problems/lru-cache |
| LC 380 | Insert Delete GetRandom O(1) | leetcode.com/problems/insert-delete-getrandom-o1 |
| LC 706 | Design HashMap | leetcode.com/problems/design-hashmap |

**LC 208 (Trie) is the single highest-value problem in this entire plan.** `Insert`, `Search`, `StartsWith` — small interface, custom structure underneath. It's today's failure, reproduced on purpose, with room to fail safely and re-solve.

**How to practice this category specifically:** before typing, write down in words what data structure backs each method and why. That's the step that had no time to happen under pressure today — rehearsing it deliberately is the fix.

---

## Phase 4 — DP & Backtracking
**~3 weeks (8 problems). Lower priority — Nebius's own scoping says DP is explicitly out of scope for backend rounds. Included for real mastery, not likelihood.**

| # | Problem | Link |
|---|---|---|
| LC 70 | Climbing Stairs | leetcode.com/problems/climbing-stairs |
| LC 198 | House Robber | leetcode.com/problems/house-robber |
| LC 322 | Coin Change | leetcode.com/problems/coin-change |
| LC 1143 | Longest Common Subsequence | leetcode.com/problems/longest-common-subsequence |
| LC 139 | Word Break | leetcode.com/problems/word-break |
| LC 78 | Subsets | leetcode.com/problems/subsets |
| LC 46 | Permutations | leetcode.com/problems/permutations |
| LC 39 | Combination Sum | leetcode.com/problems/combination-sum |

---

## Phase 5 — Consolidation
**~1 week. No new problems.**

Full mock rounds combining a Phase 3 design problem with a table-driven test, fully timed, no notes. Then a katacoding sweep: re-implement one structure from each phase, from memory, back to back.

---

## Duration summary

| Phase | New problems | Rough duration |
|---|---|---|
| 0 — Testing fluency | 3 | ~1 week |
| 1 — Linear structures | 7 | ~2–3 weeks |
| 2 — Trees & heaps | 7 | ~2–3 weeks |
| 3 — Graphs, tries, design | 7 | ~2–3 weeks |
| 4 — DP & backtracking | 8 | ~3 weeks |
| 5 — Consolidation | 0 | ~1 week |
| **Total** | **32** | **~10–11 weeks** |

Protect the weekend katacoding slot over new material if a week gets thin — that's still the mechanism that makes any of this stick, and there's no deadline forcing a trade-off the other way.
