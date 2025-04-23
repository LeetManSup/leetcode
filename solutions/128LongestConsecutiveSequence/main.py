'''
Given an unsorted array of integers nums, return the
length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
'''

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums_set = set(nums)
        max_sequence = 0
        for num in nums:
            if num-1 in nums_set:
                continue
            sequence, tmp = 1, num+1
            while tmp in nums_set:
                sequence += 1
                tmp += 1
            max_sequence = max(max_sequence, sequence)
            if sequence > len(nums) / 2:
                break
        return max_sequence