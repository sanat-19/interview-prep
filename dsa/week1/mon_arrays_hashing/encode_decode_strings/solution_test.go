package encode_decode_strings

import (
	"strings"
	"testing"
	"time"
)

func equal(a, b []string) bool {
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

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{"happy: basic", []string{"hello", "world"}},
		{"happy: three words", []string{"neet", "code", "love", "you"}},
		{"happy: one string", []string{"solo"}},
		{"edge: empty list", []string{}},
		{"edge: one empty string", []string{""}},
		{"edge: two empty strings", []string{"", ""}},
		{"edge: empty string in the middle", []string{"a", "", "b"}},
		{"edge: string with your delimiter", []string{"a#b", "#", "##", ",", ":;"}},
		{"edge: looks like your own encoding", []string{"4#abc", "3#", "0#", "12"}},
		{"edge: digits only", []string{"123", "0", "9999"}},
		{"edge: spaces and newlines", []string{" ", "a b", "line1\nline2", "\t"}},
		{"edge: unicode", []string{"héllo", "日本語", "🙂"}},
		{"edge: one very long string", []string{strings.Repeat("x#1,", 5000)}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encoded := Encode(tc.strs)
			got := Decode(encoded)
			if !equal(got, tc.strs) {
				t.Fatalf("Decode(Encode(%q)) = %q (encoded as %q)", tc.strs, got, encoded)
			}
		})
	}
}

// Time-limit check: 10^5 strings. Building the answer with s += x in a loop
// copies the whole string every time (O(n^2)) and fails.
func TestEncodeDecodeTimeLimit(t *testing.T) {
	const n = 100000
	const limit = time.Second
	strs := make([]string, n)
	for i := range strs {
		strs[i] = strings.Repeat("ab#", i%4) + "z"
	}

	done := make(chan []string, 1)
	start := time.Now()
	go func() { done <- Decode(Encode(strs)) }()

	select {
	case got := <-done:
		if !equal(got, strs) {
			t.Fatalf("round trip gave a different list (len %d, want %d)", len(got), len(strs))
		}
		t.Logf("n=%d round trip in %v (limit %v)", n, time.Since(start), limit)
	case <-time.After(limit):
		t.Fatalf("Time Limit Exceeded: n=%d took more than %v. Are you using s += x in a loop?", n, limit)
	}
}
