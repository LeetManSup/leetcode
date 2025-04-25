/*
You are given a string s and an integer k. You
can choose any character of the string and change
it to any other uppercase English character. You
can perform this operation at most k times.

Return the length of the longest substring containing
the same letter you can get after performing the
above operations.
*/

func characterReplacement(s string, k int) int {
	freqs := map[byte]int{}
	result, start, maxFreq := 0, 0, 0
	for end := range len(s) {
		freqs[s[end]]++
		maxFreq = max(maxFreq, freqs[s[end]])
		for (end-start+1)-maxFreq > k {
			freqs[s[start]]--
			start++
		}
		result = max(result, end-start+1)
	}
	return result
}