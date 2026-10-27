package first_unique_char

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"happy: first char unique", "leetcode", 0},
		{"happy: unique in the middle", "loveleetcode", 2},
		{"edge: no unique", "aabb", -1},
		{"edge: single char", "z", 0},
		{"happy: unique at the end", "aabbc", 4},
		{"edge: all same", "aaaa", -1},
		{"edge: two chars, both unique", "ab", 0},
		{"happy: repeats interleaved", "abcabcd", 6},
		{"edge: empty", "", -1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstUniqChar(tc.s); got != tc.want {
				t.Fatalf("FirstUniqChar(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
