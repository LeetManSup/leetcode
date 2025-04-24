/*
Given a string s, find the length of the longest
substring without duplicate characters.
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLength := 0
	lastIdx := map[byte]int{}
	start := 0
	for end := range n {
		currentChar := s[end]
		if lastIdx[currentChar] > start {
			start = lastIdx[currentChar]
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		lastIdx[currentChar] = end + 1
	}
	return maxLength
}