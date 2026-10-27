package range_sum_query

import (
	"testing"
	"time"
)

func TestSumRange(t *testing.T) {
	type q struct{ l, r, want int }
	tests := []struct {
		name string
		nums []int
		qs   []q
	}{
		{"happy: basic", []int{-2, 0, 3, -5, 2, -1}, []q{{0, 2, 1}, {2, 5, -1}, {0, 5, -3}}},
		{"happy: positives", []int{1, 2, 3, 4, 5}, []q{{0, 4, 15}, {1, 3, 9}, {2, 2, 3}}},
		{"edge: single element", []int{7}, []q{{0, 0, 7}}},
		{"edge: left equals right", []int{4, -4, 9}, []q{{0, 0, 4}, {1, 1, -4}, {2, 2, 9}}},
		{"edge: starts at index 0", []int{5, 5, 5, 5}, []q{{0, 1, 10}, {0, 3, 20}}},
		{"edge: ends at last index", []int{5, 5, 5, 5}, []q{{2, 3, 10}, {3, 3, 5}}},
		{"edge: all negatives", []int{-1, -2, -3}, []q{{0, 2, -6}, {1, 2, -5}}},
		{"edge: zeros", []int{0, 0, 0}, []q{{0, 2, 0}, {1, 1, 0}}},
		{"edge: same query twice", []int{3, 1, 2}, []q{{0, 1, 4}, {0, 1, 4}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.nums...)
			na := Constructor(in)
			for _, c := range tc.qs {
				if got := na.SumRange(c.l, c.r); got != c.want {
					t.Fatalf("nums=%v SumRange(%d,%d) = %d, want %d", tc.nums, c.l, c.r, got, c.want)
				}
			}
		})
	}
}

// Time-limit check: 10^5 numbers and 10^5 queries over big ranges.
// Adding up the range on every call is O(n) per query, O(n*q) total, and fails.
func TestSumRangeTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i%7 - 3
	}
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}

	done := make(chan string, 1)
	start := time.Now()
	go func() {
		na := Constructor(nums)
		for k := 0; k < n; k++ {
			l := k % 10
			r := n - 1 - (k % 10)
			want := pre[r+1] - pre[l]
			if got := na.SumRange(l, r); got != want {
				done <- "wrong answer"
				return
			}
		}
		done <- ""
	}()

	select {
	case msg := <-done:
		if msg != "" {
			t.Fatalf("%s", msg)
		}
		t.Logf("n=%d, %d queries in %v (limit %v)", n, n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: took more than %v. Does every SumRange loop over the range?", limit)
	}
}
