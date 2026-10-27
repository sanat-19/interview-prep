package product_except_self

// ProductExceptSelf returns an array answer where answer[i] is the product of
// all elements of nums except nums[i].
//
// Do it WITHOUT division. Aim for O(n) time (brute force O(n^2) first is fine).
// The product of any prefix or suffix fits in a 32-bit integer.
//
// Constraints: 2 <= len(nums) <= 10^5, -30 <= nums[i] <= 30.
func ProductExceptSelf(nums []int) []int {
	// TODO

	// brute force
	// result := []int{}
	// for i := 0; i < len(nums); i++ {
	// 	product := 1
	// 	for j := 0; j < len(nums); j++ {
	// 		if i != j {
	// 			product = nums[j] * product
	// 		}
	// 	}
	// 	result = append(result, product)
	// }

	// return result

	left := make([]int, len(nums))
	left[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		left[i+1] = left[i] * nums[i]
	}

	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		right[i-1] = right[i] * nums[i]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = right[i] * left[i]
	}

	return result
}

/*
loop 0
result = [24, ]
left = [1, ]

loop 1
result = [24, 12, ]
left = [1, 1, ]

loop 2
result = [24, 12, 8, ]
left = [1, 1, 6, ]
*/
