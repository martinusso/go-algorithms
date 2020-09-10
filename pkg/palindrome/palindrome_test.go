package palindrome

import (
	"strconv"
	"testing"
)

func TestPalindromicNumber(t *testing.T) {
	n := 1234321
	if !IsPalindromic(strconv.Itoa(n)) {
		t.Errorf("'%d' should be palindrome", n)
	}
}

func TestPalindromicString(t *testing.T) {
	palindromes := []string{
		"Able was I ere I saw Elba",
		"Madam, I'm Adam",
		"Never odd or even",
		"No 'x' in Nixon",
		"A man, a plan, a canal, Panama"}
	for _, s := range palindromes {
		if !IsPalindromic(s) {
			t.Errorf("'%s' should be palindrome", s)
		}
	}
}

func TestNotPalindromic(t *testing.T) {
	s := "May the Force be with you"
	if IsPalindromic(s) {
		t.Errorf("'%s' should'nt be palindrome", s)
	}
}
