package remove_duplicates

import (
	"reflect"
	"testing"
	"time"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"happy: basic", []int{1, 1, 2}, []int{1, 2}},
		{"happy: longer", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{"happy: no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"edge: empty", []int{}, []int{}},
		{"edge: single element", []int{1}, []int{1}},
		{"edge: all same", []int{1, 1, 1, 1}, []int{1}},
		{"edge: two same", []int{1, 1}, []int{1}},
		{"edge: negatives", []int{-3, -3, 0, 0, 5}, []int{-3, 0, 5}},
		{"edge: duplicates only at the end", []int{1, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"edge: duplicates only at the start", []int{1, 1, 1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := make([]int, len(tc.nums))
			copy(in, tc.nums)
			k := RemoveDuplicates(in)
			if k != len(tc.want) {
				t.Fatalf("RemoveDuplicates(%v) = %d, want %d", tc.nums, k, len(tc.want))
			}
			if !reflect.DeepEqual(in[:k], tc.want) {
				t.Fatalf("after RemoveDuplicates(%v) the first %d elements are %v, want %v", tc.nums, k, in[:k], tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 numbers where every value appears twice.
// Shifting the whole array on every duplicate is O(n^2) and fails.
func TestRemoveDuplicatesTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i / 2
	}

	done := make(chan int, 1)
	start := time.Now()
	go func() { done <- RemoveDuplicates(nums) }()

	select {
	case got := <-done:
		if got != n/2 {
			t.Fatalf("wrong answer: got %d, want %d", got, n/2)
		}
		for i := 0; i < got; i++ {
			if nums[i] != i {
				t.Fatalf("nums[%d] = %d, want %d", i, nums[i], i)
			}
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
