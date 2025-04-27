'''
Given a string s containing just the characters '(', ')',
'{', '}', '[' and ']', determine if the input string is
valid.

An input string is valid if:

Open brackets must be closed by the same type ofbrackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of
the same type.
'''

class Solution:
    def isValid(self, s: str) -> bool:
        match = {')': '(', ']': '[', '}': '{'}
        stack = []
        for bracket in s:
            if bracket in match:
                if stack and stack[-1] == match[bracket]:
                    stack.pop()
                else:
                    return False
            else:
                stack.append(bracket)
        return not stack