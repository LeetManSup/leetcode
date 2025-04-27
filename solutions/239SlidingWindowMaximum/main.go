/*
You are given an array of integers nums, there is
a sliding window of size k which is moving from the
very left of the array to the very right. You can
only see the k numbers in the window. Each time the
sliding window moves right by one position.

Return the max sliding window.
*/

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	deque := []int{}
	for i, num := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
		if deque[0] == i-k+1 {
			deque = deque[1:]
		}
	}
	return result
}