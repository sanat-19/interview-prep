package sort_colors

// SortColors sorts nums IN PLACE so that all 0s come first, then 1s, then 2s.
// Do not use the sort package.
//
// Example: nums = [2, 0, 2, 1, 1, 0] -> [0, 0, 1, 1, 2, 2]
//
// Constraints: 0 <= len(nums) <= 300, nums[i] is 0, 1 or 2.
// Target: O(n) time, O(1) extra space. A one-pass solution exists.
func SortColors(nums []int) {
	// TODO
	// c0, c1, c2 := 0, 0, 0
	// for _, x := range nums {
	// 	switch x {
	// 	case 0:
	// 		c0++
	// 	case 1:
	// 		c1++
	// 	default:
	// 		c2++
	// 	}
	// }

	// i := 0
	// for k := 0; k < c0; k++ {
	// 	nums[i] = 0
	// 	i++
	// }
	// for k := 0; k < c1; k++ {
	// 	nums[i] = 1
	// 	i++
	// }
	// for k := 0; k < c2; k++ {
	// 	nums[i] = 2
	// 	i++
	// }

	lo, mid, hi := 0, 0, len(nums)-1

	for mid <= hi {
		switch nums[mid] {
		case 0:
			nums[lo], nums[mid] = nums[mid], nums[lo]
			lo++
			mid++
		case 1:
			mid++
		default:
			nums[mid], nums[hi] = nums[hi], nums[mid]
			hi--
		}
	}
}

/*
3-pointers, not 2
its like
if mid value is 2 then swap with hi and hi--
if mid value is 1 then just mid++
if mid value is 0 then swap with lo and lo++
*/
