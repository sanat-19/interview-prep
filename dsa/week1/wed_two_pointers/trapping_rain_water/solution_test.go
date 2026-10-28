package trapping_rain_water

import (
	"testing"
	"time"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"happy: classic", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"happy: valley", []int{4, 2, 0, 3, 2, 5}, 9},
		{"happy: one pit", []int{2, 0, 2}, 2},
		{"happy: two pits", []int{1, 0, 2, 0, 1}, 2},
		{"edge: empty", []int{}, 0},
		{"edge: single bar", []int{1}, 0},
		{"edge: two bars", []int{1, 2}, 0},
		{"edge: flat", []int{3, 3, 3}, 0},
		{"edge: all zeros", []int{0, 0, 0}, 0},
		{"edge: decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"edge: increasing", []int{1, 2, 3, 4, 5}, 0},
		{"edge: wide pit", []int{3, 0, 0, 0, 3}, 9},
		{"edge: two pits between three walls", []int{5, 0, 5, 0, 5}, 10},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.height...)
			if got := Trap(in); got != tc.want {
				t.Fatalf("Trap(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 bars 0,1,0,1,... Every inner 0 holds 1 unit of water.
// Scanning left and right for the tallest bar at every position is O(n^2) and fails.
func TestTrapTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	height := make([]int, n)
	for i := range height {
		height[i] = i % 2
	}

	done := make(chan int, 1)
	start := time.Now()
	go func() { done <- Trap(height) }()

	select {
	case got := <-done:
		if got != 49999 {
			t.Fatalf("wrong answer: got %d, want %d", got, 49999)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
