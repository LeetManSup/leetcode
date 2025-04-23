'''
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and
you may not use the same element twice.

You can return the answer in any order.
'''

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numEnters = {}
        for i, num in enumerate(nums):
            j = target - num
            if j in numEnters:
                return [i, numEnters[j]]
            numEnters[num] = i