'''
Given two strings s1 and s2, return true if s2
contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's
permutations is the substring of s2.
'''

from collections import Counter

class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        n, m = len(s1), len(s2)
        if m < n:
            return False
        freq1, freq2 = Counter(s1), Counter(s2[0:n])
        if freq1 == freq2:
            return True
        l, r = 1, n
        while r < m:
            freq2[s2[l-1]] -= 1
            freq2[s2[r]] += 1
            if freq1 == freq2:
                return True
            r += 1
            l += 1
        return False