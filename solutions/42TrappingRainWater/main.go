/*
Given n non-negative integers representing an elevation
map where the width of each bar is 1, compute how much
water it can trap after raining.
*/

func trap(height []int) int {
	n := len(height)
	maxLeftList := make([]int, n)
	maxRightList := make([]int, n)
	maxCurrent := 0
	for i := range n {
		maxLeftList[i] = maxCurrent
		maxCurrent = max(maxCurrent, height[i])
	}
	maxCurrent = 0
	for i := n - 1; i >= 0; i-- {
		maxRightList[i] = maxCurrent
		maxCurrent = max(maxCurrent, height[i])
	}
	storedWater := 0
	for i, h := range height {
		water := min(maxLeftList[i], maxRightList[i]) - h
		if water > 0 {
			storedWater += water
		}
	}
	return storedWater
}