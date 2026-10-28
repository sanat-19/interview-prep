package three_sum

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

// normalize sorts every triplet and then the list of triplets, so the
// order of your answer does not matter.
func normalize(res [][]int) [][]int {
	out := make([][]int, 0, len(res))
	for _, tr := range res {
		c := append([]int(nil), tr...)
		sort.Ints(c)
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < 3 && k < len(out[i]) && k < len(out[j]); k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return false
	})
	return out
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"happy: classic", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"happy: no triplet", []int{0, 1, 1}, [][]int{}},
		{"happy: all zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"happy: mixed", []int{3, -1, -2, 0, 1, 2, -3}, [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
		{"edge: empty", []int{}, [][]int{}},
		{"edge: one number", []int{0}, [][]int{}},
		{"edge: two numbers", []int{0, 0}, [][]int{}},
		{"edge: many zeros, one triplet", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"edge: duplicates at the end", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"edge: duplicates at the start", []int{-1, -1, -1, 2, 2}, [][]int{{-1, -1, 2}}},
		{"edge: four numbers, no triplet", []int{1, 2, -2, -1}, [][]int{}},
		{"edge: many duplicates, many triplets", []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{"edge: zero used once, duplicate zero", []int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := append([]int(nil), tc.nums...)
			got := normalize(ThreeSum(in))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ThreeSum(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

// Time-limit check: the numbers -1500..1499 (n=3000) have 1124250 distinct
// triplets that add to 0. Trying every triple is O(n^3) and fails.
func TestThreeSumTimeLimit(t *testing.T) {
	const n = 3000
	const want = 1124250
	const limit = 3 * time.Second
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i - 1500
	}

	done := make(chan [][]int, 1)
	start := time.Now()
	go func() { done <- ThreeSum(nums) }()

	select {
	case got := <-done:
		if len(got) != want {
			t.Fatalf("wrong answer: got %d triplets, want %d", len(got), want)
		}
		for _, tr := range got {
			if len(tr) != 3 || tr[0]+tr[1]+tr[2] != 0 {
				t.Fatalf("triplet %v does not add up to 0", tr)
			}
		}
		t.Logf("n=%d solved in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Is your solution O(n^3)?", n, limit)
	}
}
