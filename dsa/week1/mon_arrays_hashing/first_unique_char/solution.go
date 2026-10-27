package first_unique_char

// FirstUniqChar returns the index of the first non-repeating character
// in s, or -1 if there is none.
//
// Constraints: 1 <= len(s) <= 10^5, lowercase English letters.
func FirstUniqChar(s string) int {
	// TODO

	// for i := 0; i < len(s); i++ {
	// 	count := 0
	// 	for j := 0; j < len(s); j++ {
	// 		if s[i] == s[j] {
	// 			count++
	// 		}
	// 	}
	// 	if count == 1 {
	// 		return i
	// 	}
	// }

	// return -1

	count := [26]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
