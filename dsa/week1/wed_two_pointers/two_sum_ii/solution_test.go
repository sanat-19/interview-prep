package two_sum_ii

import (
	"reflect"
	"testing"
	"time"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"happy: basic", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"happy: ends of array", []int{2, 3, 4}, 6, []int{1, 3}},
		{"happy: two numbers only", []int{-1, 0}, -1, []int{1, 2}},
		{"edge: duplicates", []int{3, 3}, 6, []int{1, 2}},
		{"edge: negatives and positives", []int{-5, -3, 0, 3, 8}, 0, []int{2, 4}},
		{"edge: same value twice in the middle", []int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
		{"edge: zeros", []int{0, 0, 3, 4}, 0, []int{1, 2}},
		{"edge: cannot use the same element twice", []int{1, 2, 3}, 4, []int{1, 3}},
		{"edge: all negative", []int{-10, -8, -3, -1}, -11, []int{1, 4}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.numbers...)
			if got := TwoSum(in, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("TwoSum(%v, %d) = %v, want %v", tc.numbers, tc.target, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 sorted numbers, the answer is the last two.
// Trying every pair is O(n^2) and fails.
func TestTwoSumTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 2 * i
	}
	target := nums[n-2] + nums[n-1]

	done := make(chan []int, 1)
	start := time.Now()
	go func() { done <- TwoSum(nums, target) }()

	select {
	case got := <-done:
		if !reflect.DeepEqual(got, []int{n - 1, n}) {
			t.Fatalf("wrong answer: got %v, want [%d %d]", got, n-1, n)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
