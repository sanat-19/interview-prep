package two_sum

// TwoSum returns the indices of the two numbers that add up to target.
// Exactly one solution exists; the same element can't be used twice.
// Any order of the two indices is accepted.
//
// Constraints: 2 <= len(nums) <= 10^4, -10^9 <= nums[i], target <= 10^9
func TwoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		if idx, ok := seen[x]; ok {
			return []int{i, idx}
		}
		seen[nums[i]] = i
	}
	return nil
}
