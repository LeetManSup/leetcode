'''
Given a string s, find the length of the longest
substring without duplicate characters.
'''

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        n = len(s)
        last_idx = {}
        max_length = 0
        start = 0
        for end in range(n):
            current_char = s[end]
            if current_char in last_idx and last_idx[current_char] > start:
                start = last_idx[current_char]
            max_length = max(max_length, end-start+1)
            last_idx[current_char] = end + 1
        return max_length