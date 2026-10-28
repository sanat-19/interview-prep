package valid_palindrome

// IsPalindrome reports whether s reads the same forwards and backwards,
// looking only at letters and digits and ignoring upper/lower case.
//
// Example: "A man, a plan, a canal: Panama" -> true
//
//	"race a car"                     -> false
//
// Constraints: 0 <= len(s) <= 2*10^5, s has ASCII characters.
// Target: O(n) time, O(1) extra space.
func IsPalindrome(s string) bool {
	// TODO
	left, right := 0, len(s)-1
	for left < right {
		if !isAlphaNum(s[left]) {
			left++
			continue
		}
		if !isAlphaNum(s[right]) {
			right--
			continue
		}

		if lower(s[left]) != lower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}
