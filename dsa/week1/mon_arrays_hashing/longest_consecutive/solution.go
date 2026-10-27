package longest_consecutive

// LongestConsecutive returns the length of the longest run of consecutive
// integers (like 1, 2, 3, 4) that can be formed from the numbers in nums.
// The numbers do not have to be next to each other in the array.
//
// Example: nums = [100, 4, 200, 1, 3, 2] -> 4 (the run is 1, 2, 3, 4)
//
// Constraints: 0 <= len(nums) <= 10^5, -10^9 <= nums[i] <= 10^9.
// Target: O(n) time.
func LongestConsecutive(nums []int) int {
	// TODO
	if len(nums) == 0 {
		return 0
	}
	seen := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = true
	}

	max := 1

	for x := range seen {
		if seen[x-1] {
			continue
		}
		count := 1
		for seen[x+count] {
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}
