// Package modulus11 is an implementation of the Modulus 11 check digit calculation
package modulus11

import "strconv"

// Mod11 calculate the MOD11 check digit for a slice ([]int)
func Mod11(value []int) int {
	return computeCheckDigit(value)
}

// Mod11s calculate the MOD11 check digit for a string
func Mod11s(value string) (int, error) {
	var data []int
	for _, s := range value {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return -1, err
		}
		data = append(data, n)
	}
	return Mod11(data), nil
}

// Mod11i calculate the MOD11 check digit for a number
func Mod11i(value int) int {
	s := strconv.Itoa(value)
	m, _ := Mod11s(s)
	return m
}

func computeCheckDigit(data []int) int {
	var sum int
	for i, f := len(data)-1, 2; i >= 0; i, f = i-1, f+1 {
		sum += data[i] * f
		if f == 9 {
			f = 2
		}
	}
	remainder := 11 - int(sum%11)
	if remainder > 9 {
		return 0
	}
	return remainder
}
