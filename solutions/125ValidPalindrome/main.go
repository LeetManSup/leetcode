/*
A phrase is a palindrome if, after converting all uppercase
letters into lowercase letters and removing all non-alphanumeric
characters, it reads the same forward and backward. Alphanumeric
characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false
otherwise.
*/

func isPalindrome(s string) bool {
	strToCheck := []rune{}
	for _, ch := range []rune(s) {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			strToCheck = append(strToCheck, unicode.ToLower(ch))
		}
	}
	for i, j := 0, len(strToCheck)-1; i < j; i, j = i+1, j-1 {
		if strToCheck[i] != strToCheck[j] {
			return false
		}
	}
	return true
}
