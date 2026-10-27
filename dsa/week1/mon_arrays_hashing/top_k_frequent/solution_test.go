package top_k_frequent

import (
	"sort"
	"testing"
	"time"
)

func sameSet(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	a := append([]int(nil), got...)
	b := append([]int(nil), want...)
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"happy: basic", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"happy: k is 1", []int{4, 4, 4, 5, 5, 6}, 1, []int{4}},
		{"happy: spread out", []int{3, 0, 1, 0, 3, 3, 0, 0, 2}, 2, []int{0, 3}},
		{"edge: single element", []int{1}, 1, []int{1}},
		{"edge: all same", []int{7, 7, 7, 7}, 1, []int{7}},
		{"edge: k equals distinct count", []int{1, 2, 3, 3, 2, 3}, 3, []int{1, 2, 3}},
		{"edge: negatives", []int{-1, -1, -2, -2, -2, 3}, 2, []int{-2, -1}},
		{"edge: negative and zero", []int{0, 0, 0, -5, -5, 9}, 2, []int{0, -5}},
		{"edge: most frequent comes last", []int{1, 2, 2, 3, 3, 3}, 1, []int{3}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopKFrequent(tc.nums, tc.k)
			if !sameSet(got, tc.want) {
				t.Fatalf("TopKFrequent(%v, %d) = %v, want %v (any order)", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

// Time-limit check: 10^5 numbers. O(n^2) counting takes many seconds and fails.
func TestTopKFrequentTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i%20000 - 10000
	}
	for i := 99200; i < 99700; i++ {
		nums[i] = 777
	}
	for i := 99700; i < n; i++ {
		nums[i] = 888
	}

	done := make(chan []int, 1)
	start := time.Now()
	go func() { done <- TopKFrequent(nums, 2) }()

	select {
	case got := <-done:
		if !sameSet(got, []int{777, 888}) {
			t.Fatalf("wrong answer on big input: %v", got)
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^2)?", n, limit)
	}
}
