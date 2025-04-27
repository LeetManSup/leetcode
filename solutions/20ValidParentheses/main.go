/*
Given a string s containing just the characters '(', ')',
'{', '}', '[' and ']', determine if the input string is
valid.

An input string is valid if:

Open brackets must be closed by the same type ofbrackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of
the same type.
*/

func isValid(s string) bool {
	match := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}
	for _, bracket := range s {
		if open, ok := match[bracket]; ok {
			if len(stack) != 0 && stack[len(stack)-1] == open {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, bracket)
		}
	}
	return len(stack) == 0
}