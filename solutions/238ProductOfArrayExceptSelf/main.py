'''
Given an integer array nums, return an array answer such
that answer[i] is equal to the product of all the elements
of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed
to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and
without using the division operation.
'''

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [0] * n
        prefix_prod = [0] * n
        postfix_prod = [0] * n
        prefix_prod[0], postfix_prod[-1] = 1, 1

        for i in range(1, n):
            prefix_prod[i] = prefix_prod[i-1] * nums[i-1]
        for i in range(n-2, -1, -1):
            postfix_prod[i] = postfix_prod[i+1] * nums[i+1]
        for i in range(n):
            result[i] = prefix_prod[i] * postfix_prod[i]
        
        return result