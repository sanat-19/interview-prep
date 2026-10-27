package group_anagrams

import (
	"sort"
	"strings"
	"testing"
)

// normalize makes the result order-independent so any valid answer passes.
func normalize(groups [][]string) []string {
	out := make([]string, 0, len(groups))
	for _, g := range groups {
		c := append([]string(nil), g...)
		sort.Strings(c)
		out = append(out, strings.Join(c, ","))
	}
	sort.Strings(out)
	return out
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want [][]string
	}{
		{"happy: basic", []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"edge: single empty string", []string{""}, [][]string{{""}}},
		{"edge: single word", []string{"a"}, [][]string{{"a"}}},
		{"happy: no anagrams", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{"happy: all anagrams", []string{"abc", "bca", "cab"}, [][]string{{"abc", "bca", "cab"}}},
		{"edge: duplicates", []string{"a", "a", "b"}, [][]string{{"a", "a"}, {"b"}}},
		{"edge: multiple empty strings", []string{"", ""}, [][]string{{"", ""}}},
		{"edge: same letters, different counts", []string{"aab", "abb", "baa"}, [][]string{{"aab", "baa"}, {"abb"}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize(GroupAnagrams(tc.in))
			want := normalize(tc.want)
			if strings.Join(got, "|") != strings.Join(want, "|") {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
