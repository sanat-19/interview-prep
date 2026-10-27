package longest_consecutive

import (
	"testing"
	"time"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"happy: basic", []int{100, 4, 200, 5, 3, 2}, 4},
		{"happy: longer run with duplicate zero", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"happy: already sorted", []int{1, 2, 3, 4, 5}, 5},
		{"edge: empty", []int{}, 0},
		{"edge: single element", []int{7}, 1},
		{"edge: duplicates", []int{1, 2, 0, 1}, 3},
		{"edge: all same", []int{5, 5, 5, 5}, 1},
		{"edge: no two are consecutive", []int{10, 30, 20, 50}, 1},
		{"edge: negatives", []int{-1, -2, -3, 5}, 3},
		{"edge: run crosses zero", []int{1, -1, 0}, 3},
		{"edge: two separate runs, longer one is later", []int{1, 2, 3, 10, 11, 12, 13}, 4},
		{"edge: extreme values", []int{1000000000, 999999999, -1000000000}, 2},
		{"edge: answer is not at the start", []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.nums...)
			got := LongestConsecutive(in)
			if got != tc.want {
				t.Fatalf("LongestConsecutive(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 numbers in shuffled order, one big run.
// Scanning for x+1 again and again from every number is O(n^2) and fails.
func TestLongestConsecutiveTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = (i * 7919) % n
	}

	done := make(chan int, 1)
	start := time.Now()
	go func() { done <- LongestConsecutive(nums) }()

	select {
	case got := <-done:
		if got != n {
			t.Fatalf("wrong answer: got %d, want %d", got, n)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
