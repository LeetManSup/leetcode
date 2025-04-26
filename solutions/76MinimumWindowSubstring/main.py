'''
Given two strings s and t of lengths m and n respectively,
return the minimum window substring of s such that every
character in t (including duplicates) is included in the
window. If there is no such substring, return the empty
string "".

The testcases will be generated such that the answer is
unique.
'''

from collections import Counter

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        counts = Counter(t)
        required = len(t)
        best_left = -1
        min_length = len(s) + 1
        l = 0
        for r, ch in enumerate(s):
            counts[ch] -= 1
            if counts[ch] >= 0:
                required -= 1
            while required == 0:
                if r - l + 1 < min_length:
                    best_left = l
                    min_length = r - l + 1
                counts[s[l]] += 1
                if counts[s[l]] > 0:
                    required += 1
                l += 1
        return "" if best_left == -1 else s[best_left : best_left + min_length]