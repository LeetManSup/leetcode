/*
Given an integer array nums, return true if any value
appears at least twice in the array, and return false
if every element is distinct.
*/

package containsduplicate

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	inNums := make(map[int]bool)
	for _, num := range nums {
		if !inNums[num] {
			inNums[num] = true
		} else {
			return true
		}
	}
	return false
}
