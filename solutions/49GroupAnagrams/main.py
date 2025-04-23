'''
Given an array of strings strs, group the anagrams together.
You can return the answer in any order.
'''

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for str in strs:
            key = "".join(sorted(str))
            if key not in groups:
                groups[key] = []
            groups[key].append(str)
        result = []
        for group in groups.values():
            result.append(group)
        return result