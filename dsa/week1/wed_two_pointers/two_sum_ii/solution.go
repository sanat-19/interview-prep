package two_sum_ii

// TwoSum takes numbers sorted in non-decreasing order and returns the
// 1-based indexes [i, j] (i < j) of the two numbers that add up to target.
// Exactly one answer exists. You may not use the same element twice.
//
// Example: numbers = [2, 7, 11, 15], target = 9 -> [1, 2]
//
// Constraints: 2 <= len(numbers) <= 10^5, -1000 <= numbers[i] <= 1000 (tests also use bigger values).
// Target: O(n) time, O(1) extra space.
func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
