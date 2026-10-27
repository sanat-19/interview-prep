package product_except_self

import (
	"testing"
	"time"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"happy: basic", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"happy: with ones", []int{1, 1, 2, 3}, []int{6, 6, 3, 2}},
		{"happy: two elements", []int{3, 4}, []int{4, 3}},
		{"edge: one zero", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"edge: two zeros", []int{0, 0, 5}, []int{0, 0, 0}},
		{"edge: all zeros", []int{0, 0}, []int{0, 0}},
		{"edge: negatives", []int{-1, -2, -3}, []int{6, 3, 2}},
		{"edge: zero at start", []int{0, 2, 3}, []int{6, 0, 0}},
		{"edge: zero at end", []int{2, 3, 0}, []int{0, 0, 6}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.nums...)
			got := ProductExceptSelf(in)
			if !equal(got, tc.want) {
				t.Fatalf("ProductExceptSelf(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 numbers. O(n^2) takes many seconds and fails.
func TestProductExceptSelfTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1
	}
	for i := 0; i < 30; i++ {
		nums[i*3000] = 2
	}
	total := 1 << 30

	done := make(chan []int, 1)
	start := time.Now()
	go func() { done <- ProductExceptSelf(nums) }()

	select {
	case got := <-done:
		if len(got) != n {
			t.Fatalf("wrong length %d", len(got))
		}
		for i := range got {
			if got[i] != total/nums[i] {
				t.Fatalf("wrong answer at index %d: got %d, want %d", i, got[i], total/nums[i])
			}
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
