package trapping_rain_water

// Trap takes the heights of bars (each bar is 1 wide). After rain, water is
// stuck between the bars. Return how many units of water are trapped.
//
// Example: height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1] -> 6
//
// Constraints: 0 <= len(height) <= 2*10^4 (tests also use 10^5), 0 <= height[i] <= 10^5.
// Target: O(n) time. O(1) extra space is possible with two pointers.
func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

/*
i, 10
left, right = 0, 1
count = 0
loop 0
left > right -> count = count + (left-right) = 2, right++

loop 1
left > right -> count = 2 + (4-0) = 6, right++

loop 2
left < right -> left == right, right++

loop 3
left > right -> count = 6 + (5-2) = 9, right++

loop 4
left == right -> left == right, right++

loop 5
left > right -> count = 9 + (5-2) = 12, right++

loop 6
left > right ->

//


*/
