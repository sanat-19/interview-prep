package encode_decode_strings

import (
	"strconv"
	"strings"
)

// Encode turns a list of strings into ONE string.
// Decode must turn that string back into exactly the same list.
//
// Strings can contain any characters: spaces, digits, #, commas, newlines,
// unicode. They can also be empty. The list itself can be empty too.
//
// Example: ["hello", "world"] -> Encode -> "some string" -> Decode -> ["hello", "world"]
//
// Constraints: 0 <= len(strs) <= 10^5, each string can have any characters.
// Target: O(total length) time for both.

func Encode(strs []string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(strconv.Itoa(len(s)))
		b.WriteByte('#')
		b.WriteString(s)
	}
	return b.String()
}

func Decode(s string) []string {
	res := []string{}
	i := 0
	for i < len(s) {
		j := i + strings.IndexByte(s[i:], '#')
		n, _ := strconv.Atoi(s[i:j])
		res = append(res, s[j+1:j+1+n])
		i = j + 1 + n
	}
	return res
}
