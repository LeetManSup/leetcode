'''
You are given an array of integers nums, there is
a sliding window of size k which is moving from the
very left of the array to the very right. You can
only see the k numbers in the window. Each time the
sliding window moves right by one position.

Return the max sliding window.
'''

from collections import deque

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        result = []
        deq = deque()
        for i, num in enumerate(nums):
            while len(deq) > 0 and nums[deq[-1]] <= num:
                deq.pop()
            deq.append(i)
            if i >= k-1:
                result.append(nums[deq[0]])
            if deq[0] == i - k + 1:
                deq.popleft()
        return result