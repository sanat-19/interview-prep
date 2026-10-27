package contiguous_array

import (
	"testing"
	"time"
)

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"happy: basic", []int{0, 1}, 2},
		{"happy: one extra zero", []int{0, 1, 0}, 2},
		{"happy: longer", []int{0, 0, 1, 0, 0, 0, 1, 1}, 6},
		{"happy: starts with ones", []int{1, 1, 0, 0, 1}, 4},
		{"edge: single zero", []int{0}, 0},
		{"edge: single one", []int{1}, 0},
		{"edge: all zeros", []int{0, 0, 0}, 0},
		{"edge: all ones", []int{1, 1, 1}, 0},
		{"edge: whole array is balanced", []int{0, 0, 0, 1, 1, 1}, 6},
		{"edge: answer starts at index 0", []int{0, 0, 1, 1, 0}, 4},
		{"edge: answer is in the middle", []int{1, 1, 1, 0, 1, 1, 1}, 2},
		{"edge: answer at the end", []int{0, 1, 1, 1, 1, 1, 0, 0, 0}, 6},
		{"edge: reversed pair", []int{1, 0}, 2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.nums...)
			if got := FindMaxLength(in); got != tc.want {
				t.Fatalf("FindMaxLength(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 numbers, whole array balanced.
// Checking every subarray is O(n^2) and fails.
func TestFindMaxLengthTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i % 2
	}

	done := make(chan int, 1)
	start := time.Now()
	go func() { done <- FindMaxLength(nums) }()

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
