/*
You are given an integer array height of length n.
There are n vertical lines drawn such that the two
endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a
container, such that the container contains the most
water.

Return the maximum amount of water a container can
store.

Notice that you may not slant the container.
*/

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	maxWater := 0
	for l < r && l < n-1 && r > 0 {
		water := (r - l) * min(height[l], height[r])
		maxWater = max(maxWater, water)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return maxWater
}