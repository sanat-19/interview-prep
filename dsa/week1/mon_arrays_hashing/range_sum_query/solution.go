package range_sum_query

// NumArray answers many "sum of nums[left..right]" questions on the same array.
// Constructor is called once. SumRange is called many times.
//
// Example: nums = [-2, 0, 3, -5, 2, -1]
//
//	SumRange(0, 2) = -2 + 0 + 3 = 1
//	SumRange(2, 5) = 3 + -5 + 2 + -1 = -1
//	SumRange(0, 5) = -3
//
// Constraints: 1 <= len(nums) <= 10^5, 0 <= left <= right < len(nums),
// up to 10^5 calls to SumRange. Target: each SumRange call in O(1).
type NumArray struct {
	// TODO: add the fields you need
	pre []int
}

func Constructor(nums []int) NumArray {
	// TODO
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] + nums[i]
	}
	return NumArray{pre: pre}
}

func (this *NumArray) SumRange(left int, right int) int {
	// TODO
	if left == 0 {
		return this.pre[right]
	}
	return this.pre[right] - this.pre[left-1]
}
