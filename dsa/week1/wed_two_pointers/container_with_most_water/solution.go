package container_with_most_water

// MaxArea takes the heights of vertical lines placed at x = 0, 1, 2, ...
// Pick two lines so that, with the x-axis, they form a container holding the
// most water. Return that largest amount of water. The container cannot be tilted.
//
// Example: height = [1, 8, 6, 2, 5, 4, 8, 3, 7] -> 49
// (lines at index 1 and 8: width 7, height min(8, 7) = 7, area 49)
//
// Constraints: 2 <= len(height) <= 10^5, 0 <= height[i] <= 10^4.
// Target: O(n) time, O(1) extra space.
func MaxArea(height []int) int {

	left, right := 0, len(height)-1
	max := 0
	for left < right {
		count := 0
		if height[left] < height[right] {
			count = height[left] * (right - left)
			left++
			if count > max {
				max = count
			}
		} else if height[left] == height[right] {
			count = height[left] * (right - left)
			left++
			right--
			if count > max {
				max = count
			}
		} else if height[left] > height[right] {
			count = height[right] * (right - left)
			right--
			if count > max {
				max = count
			}
		}
	}

	return max
}

/*
[7, 8, 6, 7, 7], 14
    i        j

loop 0
h[i] > h[j] {h[j] * (j-i)}

loop 1

*/
