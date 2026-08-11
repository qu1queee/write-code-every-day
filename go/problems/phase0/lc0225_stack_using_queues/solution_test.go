package lc0225

import "testing"

func TestMyStack(t *testing.T) {
	t.Skip("Phase 0 — implement MyStack before enabling this test")

	tests := []struct {
		name   string
		ops    []string
		args   []int
		expect []any
	}{
		{
			name:   "basic push pop top",
			ops:    []string{"push", "push", "top", "pop", "empty"},
			args:   []int{1, 2, 0, 0, 0},
			expect: []any{nil, nil, 2, 2, false},
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
					got = append(got, stack.Pop())
				case "top":
					got = append(got, stack.Top())
				case "empty":
					got = append(got, stack.Empty())
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
