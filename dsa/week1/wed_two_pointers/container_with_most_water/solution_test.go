package container_with_most_water

import (
	"testing"
	"time"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"happy: classic", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"happy: two lines", []int{1, 1}, 1},
		{"happy: tall at both ends", []int{4, 3, 2, 1, 4}, 16},
		{"happy: small", []int{1, 2, 1}, 2},
		{"edge: two zeros", []int{0, 0}, 0},
		{"edge: one wall is zero", []int{0, 2}, 0},
		{"edge: tall lines near the end", []int{2, 3, 4, 5, 18, 17, 6}, 17},
		{"edge: increasing", []int{1, 2, 3, 4, 5}, 6},
		{"edge: decreasing", []int{5, 4, 3, 2, 1}, 6},
		{"edge: tall ends, empty middle", []int{3, 0, 0, 0, 3}, 12},
		{"edge: short ends, empty middle", []int{1, 0, 0, 0, 0, 0, 1}, 6},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.height...)
			if got := MaxArea(in); got != tc.want {
				t.Fatalf("MaxArea(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 lines of height 1, so the answer is n-1.
// Trying every pair of lines is O(n^2) and fails.
func TestMaxAreaTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	height := make([]int, n)
	for i := range height {
		height[i] = 1
	}

	done := make(chan int, 1)
	start := time.Now()
	go func() { done <- MaxArea(height) }()

	select {
	case got := <-done:
		if got != n-1 {
			t.Fatalf("wrong answer: got %d, want %d", got, n-1)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
