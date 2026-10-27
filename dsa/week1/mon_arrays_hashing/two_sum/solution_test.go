package two_sum

import (
	"testing"
	"time"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{"happy: basic", []int{2, 7, 11, 15}, 9},
		{"happy: pair in middle", []int{3, 2, 4}, 6},
		{"edge: duplicates", []int{3, 3}, 6},
		{"edge: negatives", []int{-1, -2, -3, -4, -5}, -8},
		{"edge: zero target", []int{0, 4, 3, 0}, 0},
		{"happy: mixed signs", []int{-3, 4, 3, 90}, 0},
		{"happy: answer at the end", []int{1, 5, 9, 13, 20, 7}, 27},
		{"edge: large values", []int{1000000000, -1000000000, 5}, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if len(got) != 2 {
				t.Fatalf("want 2 indices, got %v", got)
			}
			i, j := got[0], got[1]
			if i == j || i < 0 || j < 0 || i >= len(tc.nums) || j >= len(tc.nums) {
				t.Fatalf("invalid indices %v for %v", got, tc.nums)
			}
			if tc.nums[i]+tc.nums[j] != tc.target {
				t.Fatalf("nums[%d]+nums[%d] = %d, want %d", i, j, tc.nums[i]+tc.nums[j], tc.target)
			}
		})
	}
}

// TestTwoSumTimeLimit is the "Time Limit Exceeded" check, like LeetCode.
// Worst case for brute force: the answer sits at the very end of a big array.
// O(n^2) takes many seconds here and fails; O(n) finishes in milliseconds.
// (n is deliberately larger than the problem's 10^4 so slow code gets caught.)
func TestTwoSumTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second

	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	target := (n - 1) + (n - 2) // only the last two elements add up to this

	done := make(chan []int, 1)
	start := time.Now()
	go func() { done <- TwoSum(nums, target) }()

	select {
	case got := <-done:
		elapsed := time.Since(start)
		if len(got) != 2 || nums[got[0]]+nums[got[1]] != target || got[0] == got[1] {
			t.Fatalf("wrong answer on big input: %v", got)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, elapsed, limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
