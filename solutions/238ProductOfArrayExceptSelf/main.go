/*
Given an integer array nums, return an array answer such
that answer[i] is equal to the product of all the elements
of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed
to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and
without using the division operation.
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefixProd := make([]int, n)
	postfixProd := make([]int, n)
	prefixProd[0] = 1
	postfixProd[n-1] = 1

	for i := 1; i < n; i++ {
		prefixProd[i] = prefixProd[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		postfixProd[i] = postfixProd[i+1] * nums[i+1]
	}
	for i := range n {
		result[i] = prefixProd[i] * postfixProd[i]
	}

	return result
}