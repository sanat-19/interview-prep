package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"happy: basic", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"happy: three colors", []int{2, 0, 1}, []int{0, 1, 2}},
		{"happy: mixed", []int{1, 2, 0, 1, 0, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"edge: empty", []int{}, []int{}},
		{"edge: single zero", []int{0}, []int{0}},
		{"edge: single one", []int{1}, []int{1}},
		{"edge: all twos", []int{2, 2, 2}, []int{2, 2, 2}},
		{"edge: all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"edge: already sorted", []int{0, 1, 2}, []int{0, 1, 2}},
		{"edge: reverse sorted", []int{2, 1, 0}, []int{0, 1, 2}},
		{"edge: two colors only, 0 and 2", []int{2, 0, 2, 0}, []int{0, 0, 2, 2}},
		{"edge: two colors only, 1 and 0", []int{1, 0, 1, 0}, []int{0, 0, 1, 1}},
		{"edge: two elements", []int{2, 1}, []int{1, 2}},
		{"edge: two twos then a zero", []int{2, 2, 0}, []int{0, 2, 2}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := make([]int, len(tc.nums))
			copy(in, tc.nums)
			SortColors(in)
			if !reflect.DeepEqual(in, tc.want) {
				t.Fatalf("SortColors(%v) -> %v, want %v", tc.nums, in, tc.want)
			}
		})
	}
}
