package three_sum

import "slices"

// ThreeSum returns all UNIQUE triplets [a, b, c] from nums (different indexes,
// a + b + c == 0). The same triplet must not appear twice. Order of the
// triplets, and of numbers inside a triplet, does not matter.
//
// Example: nums = [-1, 0, 1, 2, -1, -4] -> [[-1, -1, 2], [-1, 0, 1]]
//
// Constraints: 0 <= len(nums) <= 3000, -10^5 <= nums[i] <= 10^5.
// Target: O(n^2) time.
func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if (nums[i] + nums[left] + nums[right]) < 0 {
				left++
			} else if (nums[i] + nums[left] + nums[right]) > 0 {
				right--
			} else if (nums[i] + nums[left] + nums[right]) == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}
