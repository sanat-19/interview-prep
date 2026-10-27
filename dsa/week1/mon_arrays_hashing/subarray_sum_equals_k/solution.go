package subarray_sum_equals_k

// SubarraySum returns how many contiguous subarrays of nums add up to k.
//
// Example: nums = [1, 1, 1], k = 2 -> 2 (the subarrays [1,1] at index 0 and at index 1)
//
// Constraints: 1 <= len(nums) <= 10^5, -1000 <= nums[i] <= 1000, -10^7 <= k <= 10^7.
// Numbers can be negative, zero, or positive. Target: O(n) time.
func SubarraySum(nums []int, k int) int {
	/* O(n^2)
	x := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				x++
			}
		}
	}
	return x*/

	if len(nums) == 0 {
		return 0
	}
	pre := make([]int, len(nums)) // koi extra cell nahi
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] + nums[i]
	}

	seen := map[int]int{} // khaali map, koi {0:1} nahi
	count := 0
	for i := 0; i < len(pre); i++ {
		if pre[i] == k { // hissa index 0 se shuru hota hai
			count++
		}
		count += seen[pre[i]-k]
		seen[pre[i]]++
	}
	return count
}
