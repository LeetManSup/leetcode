/*
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and
you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	numEnters := make(map[int]int)
	for i, num := range nums {
		if j, ok := numEnters[target-num]; ok {
			return []int{i, j}
		}
		numEnters[num] = i
	}
	return nil
}

