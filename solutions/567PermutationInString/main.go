/*
Given two strings s1 and s2, return true if s2
contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's
permutations is the substring of s2.
*/

func checkInclusion(s1 string, s2 string) bool {
	freqs := [26]int{}
	for i := range s1 {
		freqs[s1[i]-'a']++
	}
	left := 0
	for right := range s2 {
		freqs[s2[right]-'a']--
		if freqs == [26]int{} {
			return true
		}
		if right+1 >= len(s1) {
			freqs[s2[left]-'a']++
			left++
		}
	}
	return false
}