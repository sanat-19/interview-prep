package group_anagrams

import (
	"sort"
)

// GroupAnagrams groups the strings that are anagrams of each other.
// Order of groups and order inside a group don't matter.
//
// Constraints: 1 <= len(strs) <= 10^4, 0 <= len(strs[i]) <= 100, lowercase letters.

func GroupAnagrams(strs []string) [][]string {

	group := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		bytes := []byte(strs[i])
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		key := string(bytes)
		group[key] = append(group[key], strs[i])
	}

	collection := [][]string{}

	for _, v := range group {
		collection = append(collection, v)
	}

	return collection
}

// func GroupAnagrams(strs []string) [][]string {
// 	// TODO
// 	group := [][]string{}

// 	for _, word := range strs {

// 		found := false
// 		for i := 0; i < len(group); i++ {
// 			if isAnagram(word, group[i][0]) {
// 				group[i] = append(group[i], word)
// 				found = true
// 				break
// 			}
// 		}

// 		if !found {
// 			group = append(group, []string{word})
// 		}
// 	}

// 	return group
// }

// func isAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	count := [26]int{}

// 	for i := 0; i < len(s); i++ {
// 		count[s[i]-'a']++
// 		count[t[i]-'a']--
// 	}

// 	for i := 0; i < len(count); i++ {
// 		if count[i] != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
