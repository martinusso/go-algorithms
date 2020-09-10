/*
Package palindrome checks for Palindrome strings or numbers

A palindrome is a word, phrase, number, or other sequence of characters which reads the same backward or forward.

Examples:
- redivider
- deified
- civic
- A man, a plan, a canal, Panama
- Able was I ere I saw Elba
- Madam, I'm Adam
- Never odd or even
*/
package palindrome

import (
	"regexp"
	"strings"
)

// IsPalindromic returns if a string is palindromic
func IsPalindromic(value string) bool {
	value = sanitize(value)
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}
	return true
}

func sanitize(value string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(value, "")
	return strings.ToLower(strings.Trim(safe, ""))
}
