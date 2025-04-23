/*
Given an unsorted array of integers nums, return the
length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}
	maxSequence := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}
		sequence, tmp := 1, num+1
		for set[tmp] {
			sequence++
			tmp++
		}
		maxSequence = max(maxSequence, sequence)
		if sequence > len(nums)/2 {
			break
		}
	}
	return maxSequence
}
