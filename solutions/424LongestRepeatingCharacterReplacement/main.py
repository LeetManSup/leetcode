'''
You are given a string s and an integer k. You
can choose any character of the string and change
it to any other uppercase English character. You
can perform this operation at most k times.

Return the length of the longest substring containing
the same letter you can get after performing the
above operations.
'''

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        freqs = {}
        result, start, max_freq = 0, 0, 0
        for end in range(len(s)):
            if s[end] not in freqs:
                freqs[s[end]] = 0
            freqs[s[end]] += 1
            max_freq = max(max_freq, freqs[s[end]])
            while (end-start+1)-max_freq > k:
                freqs[s[start]] -= 1
                start += 1
            result = max(result, end-start+1)
        return result