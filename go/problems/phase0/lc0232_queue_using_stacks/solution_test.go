package lc0232

import "testing"

func TestMyQueue(t *testing.T) {
	t.Skip("Phase 0 — implement MyQueue before enabling this test")

	tests := []struct {
		name   string
		ops    []string
		args   []int
		expect []any
	}{
		{
			name:   "basic enqueue dequeue peek",
			ops:    []string{"push", "push", "peek", "pop", "empty"},
			args:   []int{1, 2, 0, 0, 0},
			expect: []any{nil, nil, 1, 1, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := Constructor()
			var got []any
			argIdx := 0

			for _, op := range tt.ops {
				switch op {
				case "push":
					queue.Push(tt.args[argIdx])
					argIdx++
					got = append(got, nil)
				case "pop":
					got = append(got, queue.Pop())
				case "peek":
					got = append(got, queue.Peek())
				case "empty":
					got = append(got, queue.Empty())
				}
			}

			if len(got) != len(tt.expect) {
				t.Fatalf("got %d results, want %d", len(got), len(tt.expect))
			}
			for i := range tt.expect {
				if got[i] != tt.expect[i] {
					t.Errorf("step %d: got %v, want %v", i, got[i], tt.expect[i])
				}
			}
		})
	}
}
