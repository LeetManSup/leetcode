'''
Given an integer array nums and an integer k, return the k most
frequent elements. You may return the answer in any order.
'''

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counts = {} # num: count
        for num in nums:
            if num not in counts:
                counts[num] = 0
            counts[num] += 1
        result = sorted(counts.keys(), key=lambda x: counts[x], reverse=True)[:k]
        return result