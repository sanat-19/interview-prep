package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"happy: classic sentence", "A man, a plan, a canal: Panama", true},
		{"happy: not a palindrome", "race a car", false},
		{"happy: plain palindrome", "abba", true},
		{"happy: plain non palindrome", "abc", false},
		{"edge: empty string", "", true},
		{"edge: only a space", " ", true},
		{"edge: single letter", "a", true},
		{"edge: only punctuation", ".,", true},
		{"edge: digit and letter differ", "0P", false},
		{"edge: underscore is ignored", "ab_a", true},
		{"edge: upper and lower case", "Aa", true},
		{"edge: digits palindrome", "12321", true},
		{"edge: digits and letters", "1a2", false},
		{"edge: odd length with symbols", "a.b,a", true},
		{"edge: mismatch near the middle", "abcdba", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.s); got != tc.want {
				t.Fatalf("IsPalindrome(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
