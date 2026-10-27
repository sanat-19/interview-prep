package subarray_sum_equals_k

import (
	"testing"
	"time"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"happy: basic", []int{1, 1, 1}, 2, 2},
		{"happy: positives", []int{1, 2, 3}, 3, 2},
		{"happy: classic example", []int{3, 4, 7, 2, -3, 1, 4, 2}, 7, 4},
		{"happy: repeated pattern", []int{1, 2, 1, 2, 1}, 3, 4},
		{"edge: single element equals k", []int{5}, 5, 1},
		{"edge: single element not equal", []int{5}, 3, 0},
		{"edge: whole array is the answer", []int{1, 2, 3}, 6, 1},
		{"edge: negatives and zero", []int{1, -1, 0}, 0, 3},
		{"edge: all zeros, k is zero", []int{0, 0, 0}, 0, 6},
		{"edge: negatives, k is zero", []int{-1, -1, 1}, 0, 1},
		{"edge: opposite numbers", []int{-5, 5}, 0, 1},
		{"edge: k is zero, no zero subarray", []int{1}, 0, 0},
		{"edge: negative k, no match", []int{2, 2, 2}, -2, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.nums...)
			if got := SubarraySum(in, tc.k); got != tc.want {
				t.Fatalf("SubarraySum(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 ones, k = 50000.
// Trying every subarray is O(n^2) and fails.
func TestSubarraySumTimeLimit(t *testing.T) {
	const n = 100000
	const k = 50000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}

	done := make(chan int, 1)
	start := time.Now()
	go func() { done <- SubarraySum(nums, k) }()

	select {
	case got := <-done:
		if want := n - k + 1; got != want {
			t.Fatalf("wrong answer: got %d, want %d", got, want)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
