/*
Given two strings s and t, return true if t is an anagram
of s, and false otherwise.
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int, len(s))
	for _, char := range []rune(s) {
		counts[char]++
	}
	for _, char := range []rune(t) {
		if counts[char] == 0 {
			return false
		}
		counts[char]--
	}
	return true
}
