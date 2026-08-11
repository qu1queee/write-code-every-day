package lc0155

import "testing"

func TestMinStack(t *testing.T) {
	t.Skip("Phase 0 — implement MinStack before enabling this test")

	tests := []struct {
		name   string
		ops    []string
		args   []int
		expect []any
	}{
		{
			name:   "tracks minimum across pushes and pops",
			ops:    []string{"push", "push", "push", "getMin", "pop", "top", "getMin"},
			args:   []int{-2, 0, -3, 0, 0, 0, 0},
			expect: []any{nil, nil, nil, -3, nil, 0, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := Constructor()
			var got []any
			argIdx := 0

			for _, op := range tt.ops {
				switch op {
				case "push":
					stack.Push(tt.args[argIdx])
					argIdx++
					got = append(got, nil)
				case "pop":
					stack.Pop()
					got = append(got, nil)
				case "top":
					got = append(got, stack.Top())
				case "getMin":
					got = append(got, stack.GetMin())
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
