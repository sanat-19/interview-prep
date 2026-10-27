package top_k_frequent

import (
	"sort"
)

// TopKFrequent returns the k most frequent elements of nums, in any order.
// The answer is guaranteed to be unique (no tie at the k-th place).
//
// Constraints: 1 <= len(nums) <= 10^5, -10^4 <= nums[i] <= 10^4,
// 1 <= k <= number of distinct elements in nums.
func TopKFrequent(nums []int, k int) []int {
	// TODO

	//brute force
	// count := map[int]int{}

	// for i := 0; i < len(nums); i++ {
	// 	count[nums[i]]++
	// }

	// result := []int{}

	// for round := 0; round < k; round++ {
	// 	max, maxKey := 0, 0
	// 	for x, v := range count {
	// 		if v > max {
	// 			max, maxKey = v, x
	// 		}
	// 	}
	// 	result = append(result, maxKey)
	// 	delete(count, maxKey)
	// }

	// return result

	count := map[int]int{}

	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	type pair struct {
		num int
		cnt int
	}

	pairs := make([]pair, 0, len(count))
	for key, val := range count {
		pairs = append(pairs, pair{key, val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].cnt > pairs[j].cnt
	})

	result := []int{}

	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}

	return result
}
