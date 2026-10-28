package remove_duplicates

// RemoveDuplicates removes duplicates from the sorted slice nums IN PLACE,
// so each value appears once, and returns k, the number of unique values.
// The first k elements of nums must hold the unique values in order.
// What is stored after index k does not matter.
//
// Example: nums = [0, 0, 1, 1, 1, 2] -> returns 3, nums[:3] = [0, 1, 2]
//
// Constraints: 0 <= len(nums) <= 3*10^4 (tests also use 10^5), nums sorted ascending.
// Target: O(n) time, O(1) extra space.
func RemoveDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
