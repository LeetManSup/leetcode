/*
Given two strings s and t of lengths m and n respectively,
return the minimum window substring of s such that every
character in t (including duplicates) is included in the
window. If there is no such substring, return the empty
string "".

The testcases will be generated such that the answer is
unique.
*/

func minWindow(s string, t string) string {
	count := map[byte]int{}
	for _, ch := range []byte(t) {
		count[ch]++
	}
	required := len(t)
	bestLeft := -1
	minLength := len(s) + 1
	l := 0
	for r, ch := range []byte(s) {
		count[ch]--
		if count[ch] >= 0 {
			required--
		}
		for required == 0 {
			if r-l+1 < minLength {
				bestLeft = l
				minLength = r - l + 1
			}
			count[s[l]]++
			if count[s[l]] > 0 {
				required++
			}
			l++
		}
	}
	if bestLeft == -1 {
		return ""
	} else {
		return s[bestLeft : bestLeft+minLength]
	}
}