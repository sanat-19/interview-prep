package valid_anagram

// IsAnagram reports whether t is an anagram of s
// (same letters, same counts, any order).
//
// Constraints: 1 <= len(s), len(t) <= 5*10^4, lowercase English letters.
// Follow-up: what if the input has Unicode characters?
func IsAnagram(s, t string) bool {
	// TODO
	if len(s) != len(t) {
		return false
	}

	// there := make(map[byte]int)

	// for i := 0; i < len(s); i++ {
	// 	there[s[i]]++
	// }

	// for i := 0; i < len(t); i++ {
	// 	if there[t[i]] > 0 {
	// 		there[t[i]]--
	// 	}
	// }

	// for _, v := range there {
	// 	if v > 0 {
	// 		return false
	// 	}
	// }

	// return true

	count := [26]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}

	return true
}
