package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"happy: basic true", "anagram", "nagaram", true},
		{"happy: basic false", "rat", "car", false},
		{"edge: different length", "a", "ab", false},
		{"edge: single char same", "a", "a", true},
		{"edge: single char diff", "a", "b", false},
		{"happy: same string", "abc", "abc", true},
		{"edge: repeated letters, wrong counts", "aacc", "ccac", false},
		{"edge: all same letter", "aaaa", "aaaa", true},
		{"edge: same set, different counts", "aab", "abb", false},
		{"edge: empty strings", "", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsAnagram(tc.s, tc.t); got != tc.want {
				t.Fatalf("IsAnagram(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
