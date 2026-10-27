package contiguous_array

// FindMaxLength returns the length of the longest contiguous subarray
// that has an equal number of 0s and 1s.
//
// Example: nums = [0, 1, 0] -> 2 (the subarray [0, 1] or [1, 0])
//
// Constraints: 1 <= len(nums) <= 10^5, nums[i] is 0 or 1.
// Target: O(n) time.
func FindMaxLength(nums []int) int {

	// max := 0
	// for i := 0; i < len(nums); i++ {
	// 	zeros, ones := 0, 0
	// 	for j := i; j < len(nums); j++ {
	// 		if nums[j] == 1 {
	// 			ones++
	// 		} else {
	// 			zeros++
	// 		}
	// 		if zeros == ones {
	// 			if max < (zeros + ones) {
	// 				max = zeros + ones
	// 			}
	// 		}

	// 	}

	// }

	// return max

	seen := map[int]int{0: -1}
	sum, max := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}

		if idx, ok := seen[sum]; ok {
			if i-idx > max {
				max = i - idx
			}
		} else {
			seen[sum] = i
		}
	}

	return max
}
