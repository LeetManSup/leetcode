/*
Given an integer array nums and an integer k, return the k most
frequent elements. You may return the answer in any order.
*/

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	top := [][]int{}
	for num, count := range counts {
		top = append(top, []int{num, count})
	}
	sort.Slice(top, func(i, j int) bool { return top[i][1] > top[j][1] })
	result := []int{}
	for i := range k {
		result = append(result, top[i][0])
	}
	return result
}