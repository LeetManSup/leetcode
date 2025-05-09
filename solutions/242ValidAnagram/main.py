'''
Given two strings s and t, return true if t is an anagram
of s, and false otherwise.
'''

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        for letter in set(s):
            if s.count(letter) != t.count(letter):
                return False
        return True
